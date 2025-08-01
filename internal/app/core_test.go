// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package app_test

import (
	"github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	"github.com/ironcore-dev/ironcore-net/apimachinery/api/net"
	. "github.com/ironcore-dev/ironcore-net/utils/testing"
	. "github.com/ironcore-dev/ironcore/utils/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"
)

var _ = Describe("Core", func() {
	ns := SetupNamespace(&k8sClient)

	Context("Network", func() {
		It("should maintain network ID allocations for networks", func(ctx SpecContext) {
			By("creating a network")
			network := &v1alpha1.Network{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns.Name,
					GenerateName: "network-",
				},
			}
			Expect(k8sClient.Create(ctx, network)).To(Succeed())

			By("inspecting the network for its ID")
			Expect(network.Spec.ID).NotTo(BeEmpty())

			By("retrieving the corresponding network ID")
			networkID := &v1alpha1.NetworkID{}
			networkIDKey := client.ObjectKey{Name: network.Spec.ID}
			Expect(k8sClient.Get(ctx, networkIDKey, networkID)).To(Succeed())

			By("inspecting the network ID")
			Expect(networkID.Spec.ClaimRef).To(Equal(v1alpha1.NetworkIDClaimRef{
				Group:     v1alpha1.GroupName,
				Resource:  "networks",
				Namespace: network.Namespace,
				Name:      network.Name,
				UID:       network.UID,
			}))

			By("deleting the network")
			Expect(k8sClient.Delete(ctx, network)).To(Succeed())

			By("asserting the corresponding network ID is gone")
			Expect(k8sClient.Get(ctx, networkIDKey, networkID)).To(Satisfy(apierrors.IsNotFound))
		})
	})

	Context("IP", func() {
		It("should handle multiple public prefixes of the same IP family", func(ctx SpecContext) {
			By("creating IPs to exhaust the first prefix (10.0.0.0/29)")
			var firstPrefixIPs []*v1alpha1.IP
			for i := 0; i < 8; i++ { // /29 has 8 IPs
				ip := &v1alpha1.IP{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:    ns.Name,
						GenerateName: "ip-",
					},
					Spec: v1alpha1.IPSpec{
						Type:     v1alpha1.IPTypePublic,
						IPFamily: corev1.IPv4Protocol,
					},
				}
				Expect(k8sClient.Create(ctx, ip)).To(Succeed())
				Expect(&ip.Spec.IP).To(Satisfy((*net.IP).IsValid))

				By("verifying the new IP is from a different prefix")
				ipAddr := ip.Spec.IP.Addr
				Expect(ipAddr.Is4()).To(BeTrue())
				Expect(ipAddr.String()).To(HavePrefix("10.0.0."))
				firstPrefixIPs = append(firstPrefixIPs, ip)
			}

			By("creating another IP which should come from the second prefix (10.0.1.0/29)")
			newIP := &v1alpha1.IP{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns.Name,
					GenerateName: "ip-",
				},
				Spec: v1alpha1.IPSpec{
					Type:     v1alpha1.IPTypePublic,
					IPFamily: corev1.IPv4Protocol,
				},
			}
			Expect(k8sClient.Create(ctx, newIP)).To(Succeed())
			Expect(&newIP.Spec.IP).To(Satisfy((*net.IP).IsValid))

			By("verifying the new IP is from a different prefix")
			newIPAddr := newIP.Spec.IP.Addr
			Expect(newIPAddr.Is4()).To(BeTrue())
			Expect(newIPAddr.String()).To(HavePrefix("10.0.1."))

			By("cleaning up all IPs")
			for _, ip := range firstPrefixIPs {
				Expect(k8sClient.Delete(ctx, ip)).To(Succeed())
			}
			Expect(k8sClient.Delete(ctx, newIP)).To(Succeed())
		})
		It("should maintain IP address allocations for IPs", func(ctx SpecContext) {
			By("creating an IP")
			ip := &v1alpha1.IP{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns.Name,
					GenerateName: "ip-",
				},
				Spec: v1alpha1.IPSpec{
					Type:     v1alpha1.IPTypePublic,
					IPFamily: corev1.IPv4Protocol,
				},
			}
			Expect(k8sClient.Create(ctx, ip)).To(Succeed())

			By("inspecting the IP for its address")
			Expect(&ip.Spec.IP).To(Satisfy((*net.IP).IsValid))

			By("retrieving the corresponding IP address")
			ipAddress := &v1alpha1.IPAddress{}
			ipAddressKey := client.ObjectKey{Name: ip.Spec.IP.String()}
			Expect(k8sClient.Get(ctx, ipAddressKey, ipAddress)).To(Succeed())

			By("inspecting the IP address")
			Expect(ipAddress.Spec.ClaimRef).To(Equal(v1alpha1.IPAddressClaimRef{
				Group:     v1alpha1.GroupName,
				Resource:  "ips",
				Namespace: ip.Namespace,
				Name:      ip.Name,
				UID:       ip.UID,
			}))

			By("deleting the IP")
			Expect(k8sClient.Delete(ctx, ip)).To(Succeed())

			By("asserting the corresponding IP address is gone")
			Expect(k8sClient.Get(ctx, ipAddressKey, ipAddress)).To(Satisfy(apierrors.IsNotFound))
		})
	})

	Context("LoadBalancer", func() {
		It("should maintain IP allocations for load balancers", func(ctx SpecContext) {
			By("creating a load balancer")
			loadBalancer := &v1alpha1.LoadBalancer{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns.Name,
					GenerateName: "lb-",
				},
				Spec: v1alpha1.LoadBalancerSpec{
					Type:       v1alpha1.LoadBalancerTypePublic,
					NetworkRef: corev1.LocalObjectReference{Name: "my-network"},
					IPs: []v1alpha1.LoadBalancerIP{
						{
							Name:     "ip-1",
							IPFamily: corev1.IPv4Protocol,
						},
					},
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"foo": "bar"},
					},
					Template: v1alpha1.InstanceTemplate{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"foo": "bar"},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, loadBalancer)).To(Succeed())

			By("inspecting the load balancer for its IPs")
			var ip net.IP
			Expect(loadBalancer.Spec.IPs).To(ConsistOf(
				HaveField("IP", Capture(&ip, AsRef(Satisfy((*net.IP).IsValid)))),
			))

			By("retrieving the corresponding IP")
			Eventually(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)).Should(HaveField("Items", ConsistOf(
				HaveField("Spec.IP", ip),
			)))

			By("updating the load balancer IPs")
			Eventually(Update(loadBalancer, func() {
				loadBalancer.Spec.IPs = []v1alpha1.LoadBalancerIP{
					{
						Name:     "new-ip-1",
						IPFamily: corev1.IPv4Protocol,
					},
				}
			})).Should(Succeed())

			By("inspecting the load balancer IPs")
			var newIP net.IP
			Expect(loadBalancer.Spec.IPs).To(ConsistOf(
				HaveField("IP", Capture(&newIP, SatisfyAll(
					Not(Equal(ip)),
					AsRef(Satisfy((*net.IP).IsValid)),
				))),
			))

			By("retrieving the corresponding IP")
			Eventually(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)).Should(HaveField("Items", ConsistOf(
				HaveField("Spec.IP", newIP),
			)))

			By("deleting the load balancer")
			Expect(k8sClient.Delete(ctx, loadBalancer)).To(Succeed())

			By("asserting there are no IPs")
			Expect(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)()).To(HaveField("Items", BeEmpty()))
		})
	})

	Context("NATGateway", func() {
		It("should maintain IP allocations for NAT gateways", func(ctx SpecContext) {
			By("creating a NAT gateway")
			natGateway := &v1alpha1.NATGateway{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns.Name,
					GenerateName: "lb-",
				},
				Spec: v1alpha1.NATGatewaySpec{
					NetworkRef: corev1.LocalObjectReference{Name: "foo"},
					IPFamily:   corev1.IPv4Protocol,
					IPs: []v1alpha1.NATGatewayIP{
						{
							Name: "ip-1",
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, natGateway)).To(Succeed())

			By("inspecting the NAT gateway for its IPs")
			var ip net.IP
			Expect(natGateway.Spec.IPs).To(ConsistOf(
				HaveField("IP", Capture(&ip, AsRef(Satisfy((*net.IP).IsValid)))),
			))

			By("retrieving the corresponding IP")
			Eventually(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)).Should(HaveField("Items", ConsistOf(
				HaveField("Spec.IP", ip),
			)))

			By("updating the NAT gateway IPs")
			Eventually(Update(natGateway, func() {
				natGateway.Spec.IPs = []v1alpha1.NATGatewayIP{
					{
						Name: "new-ip-1",
					},
				}
			})).Should(Succeed())

			By("inspecting the NAT gateway IPs")
			var newIP net.IP
			Expect(natGateway.Spec.IPs).To(ConsistOf(
				HaveField("IP", Capture(&newIP, SatisfyAll(
					Not(Equal(ip)),
					AsRef(Satisfy((*net.IP).IsValid)),
				))),
			))

			By("retrieving the corresponding IP")
			Eventually(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)).Should(HaveField("Items", ConsistOf(
				HaveField("Spec.IP", newIP),
			)))

			By("deleting the NAT gateway")
			Expect(k8sClient.Delete(ctx, natGateway)).To(Succeed())

			By("asserting there are no IPs")
			Expect(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)()).To(HaveField("Items", BeEmpty()))
		})
	})

	Context("NetworkInterface", func() {
		It("should maintain IP allocations for network interfaces", func(ctx SpecContext) {
			By("creating a network interface")
			networkInterface := &v1alpha1.NetworkInterface{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns.Name,
					GenerateName: "lb-",
				},
				Spec: v1alpha1.NetworkInterfaceSpec{
					PublicIPs: []v1alpha1.NetworkInterfacePublicIP{
						{
							Name:     "ip-1",
							IPFamily: corev1.IPv4Protocol,
						},
					},
					NodeRef: corev1.LocalObjectReference{Name: "my-node"},
				},
			}
			Expect(k8sClient.Create(ctx, networkInterface)).To(Succeed())

			By("inspecting the network interface for its IPs")
			var ip net.IP
			Expect(networkInterface.Spec.PublicIPs).To(ConsistOf(
				HaveField("IP", Capture(&ip, AsRef(Satisfy((*net.IP).IsValid)))),
			))

			By("retrieving the corresponding IP")
			Eventually(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)).Should(HaveField("Items", ConsistOf(
				HaveField("Spec.IP", ip),
			)))

			By("updating the network interface IPs")
			Eventually(Update(networkInterface, func() {
				networkInterface.Spec.PublicIPs = []v1alpha1.NetworkInterfacePublicIP{
					{
						Name:     "new-ip-1",
						IPFamily: corev1.IPv4Protocol,
					},
				}
			})).Should(Succeed())

			By("inspecting the network interface IPs")
			var newIP net.IP
			Expect(networkInterface.Spec.PublicIPs).To(ConsistOf(
				HaveField("IP", Capture(&newIP, SatisfyAll(
					Not(Equal(ip)),
					AsRef(Satisfy((*net.IP).IsValid)),
				))),
			))

			By("retrieving the corresponding IP")
			Eventually(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)).Should(HaveField("Items", ConsistOf(
				HaveField("Spec.IP", newIP),
			)))

			By("deleting the network interface")
			Expect(k8sClient.Delete(ctx, networkInterface)).To(Succeed())

			By("asserting there are no IPs")
			Expect(ObjectList(&v1alpha1.IPList{},
				client.InNamespace(ns.Name),
			)()).To(HaveField("Items", BeEmpty()))
		})
	})
})
