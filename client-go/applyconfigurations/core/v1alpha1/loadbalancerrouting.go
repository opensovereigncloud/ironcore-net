// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	internal "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// LoadBalancerRoutingApplyConfiguration represents an declarative configuration of the LoadBalancerRouting type for use
// with apply.
type LoadBalancerRoutingApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Destinations                     []LoadBalancerDestinationApplyConfiguration `json:"destinations,omitempty"`
}

// LoadBalancerRouting constructs an declarative configuration of the LoadBalancerRouting type for use with
// apply.
func LoadBalancerRouting(name, namespace string) *LoadBalancerRoutingApplyConfiguration {
	b := &LoadBalancerRoutingApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("LoadBalancerRouting")
	b.WithAPIVersion("core.apinet.ironcore.dev/v1alpha1")
	return b
}

// ExtractLoadBalancerRouting extracts the applied configuration owned by fieldManager from
// loadBalancerRouting. If no managedFields are found in loadBalancerRouting for fieldManager, a
// LoadBalancerRoutingApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// loadBalancerRouting must be a unmodified LoadBalancerRouting API object that was retrieved from the Kubernetes API.
// ExtractLoadBalancerRouting provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractLoadBalancerRouting(loadBalancerRouting *corev1alpha1.LoadBalancerRouting, fieldManager string) (*LoadBalancerRoutingApplyConfiguration, error) {
	return extractLoadBalancerRouting(loadBalancerRouting, fieldManager, "")
}

// ExtractLoadBalancerRoutingStatus is the same as ExtractLoadBalancerRouting except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractLoadBalancerRoutingStatus(loadBalancerRouting *corev1alpha1.LoadBalancerRouting, fieldManager string) (*LoadBalancerRoutingApplyConfiguration, error) {
	return extractLoadBalancerRouting(loadBalancerRouting, fieldManager, "status")
}

func extractLoadBalancerRouting(loadBalancerRouting *corev1alpha1.LoadBalancerRouting, fieldManager string, subresource string) (*LoadBalancerRoutingApplyConfiguration, error) {
	b := &LoadBalancerRoutingApplyConfiguration{}
	err := managedfields.ExtractInto(loadBalancerRouting, internal.Parser().Type("com.github.ironcore-dev.ironcore-net.api.core.v1alpha1.LoadBalancerRouting"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(loadBalancerRouting.Name)
	b.WithNamespace(loadBalancerRouting.Namespace)

	b.WithKind("LoadBalancerRouting")
	b.WithAPIVersion("core.apinet.ironcore.dev/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithKind(value string) *LoadBalancerRoutingApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithAPIVersion(value string) *LoadBalancerRoutingApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithName(value string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithGenerateName(value string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithNamespace(value string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithUID(value types.UID) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithResourceVersion(value string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithGeneration(value int64) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithCreationTimestamp(value metav1.Time) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *LoadBalancerRoutingApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *LoadBalancerRoutingApplyConfiguration) WithLabels(entries map[string]string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *LoadBalancerRoutingApplyConfiguration) WithAnnotations(entries map[string]string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *LoadBalancerRoutingApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *LoadBalancerRoutingApplyConfiguration) WithFinalizers(values ...string) *LoadBalancerRoutingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *LoadBalancerRoutingApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithDestinations adds the given value to the Destinations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Destinations field.
func (b *LoadBalancerRoutingApplyConfiguration) WithDestinations(values ...*LoadBalancerDestinationApplyConfiguration) *LoadBalancerRoutingApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDestinations")
		}
		b.Destinations = append(b.Destinations, *values[i])
	}
	return b
}
