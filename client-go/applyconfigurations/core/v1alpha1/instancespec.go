/*
 * Copyright (c) 2022 by the IronCore authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	net "github.com/ironcore-dev/ironcore-net/apimachinery/api/net"
	v1 "k8s.io/api/core/v1"
)

// InstanceSpecApplyConfiguration represents an declarative configuration of the InstanceSpec type for use
// with apply.
type InstanceSpecApplyConfiguration struct {
	Type                      *v1alpha1.InstanceType                       `json:"type,omitempty"`
	LoadBalancerType          *v1alpha1.LoadBalancerType                   `json:"loadBalancerType,omitempty"`
	NetworkRef                *v1.LocalObjectReference                     `json:"networkRef,omitempty"`
	IPs                       []net.IP                                     `json:"ips,omitempty"`
	LoadBalancerPorts         []LoadBalancerPortApplyConfiguration         `json:"loadBalancerPorts,omitempty"`
	Affinity                  *AffinityApplyConfiguration                  `json:"affinity,omitempty"`
	TopologySpreadConstraints []TopologySpreadConstraintApplyConfiguration `json:"topologySpreadConstraints,omitempty"`
	NodeRef                   *v1.LocalObjectReference                     `json:"nodeRef,omitempty"`
}

// InstanceSpecApplyConfiguration constructs an declarative configuration of the InstanceSpec type for use with
// apply.
func InstanceSpec() *InstanceSpecApplyConfiguration {
	return &InstanceSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithType(value v1alpha1.InstanceType) *InstanceSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithLoadBalancerType sets the LoadBalancerType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LoadBalancerType field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithLoadBalancerType(value v1alpha1.LoadBalancerType) *InstanceSpecApplyConfiguration {
	b.LoadBalancerType = &value
	return b
}

// WithNetworkRef sets the NetworkRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkRef field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithNetworkRef(value v1.LocalObjectReference) *InstanceSpecApplyConfiguration {
	b.NetworkRef = &value
	return b
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *InstanceSpecApplyConfiguration) WithIPs(values ...net.IP) *InstanceSpecApplyConfiguration {
	for i := range values {
		b.IPs = append(b.IPs, values[i])
	}
	return b
}

// WithLoadBalancerPorts adds the given value to the LoadBalancerPorts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the LoadBalancerPorts field.
func (b *InstanceSpecApplyConfiguration) WithLoadBalancerPorts(values ...*LoadBalancerPortApplyConfiguration) *InstanceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithLoadBalancerPorts")
		}
		b.LoadBalancerPorts = append(b.LoadBalancerPorts, *values[i])
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithAffinity(value *AffinityApplyConfiguration) *InstanceSpecApplyConfiguration {
	b.Affinity = value
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *InstanceSpecApplyConfiguration) WithTopologySpreadConstraints(values ...*TopologySpreadConstraintApplyConfiguration) *InstanceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTopologySpreadConstraints")
		}
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, *values[i])
	}
	return b
}

// WithNodeRef sets the NodeRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeRef field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithNodeRef(value v1.LocalObjectReference) *InstanceSpecApplyConfiguration {
	b.NodeRef = &value
	return b
}
