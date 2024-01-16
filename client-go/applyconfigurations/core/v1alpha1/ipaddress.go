// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	internal "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/internal"
	v1 "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// IPAddressApplyConfiguration represents an declarative configuration of the IPAddress type for use
// with apply.
type IPAddressApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *IPAddressSpecApplyConfiguration `json:"spec,omitempty"`
}

// IPAddress constructs an declarative configuration of the IPAddress type for use with
// apply.
func IPAddress(name string) *IPAddressApplyConfiguration {
	b := &IPAddressApplyConfiguration{}
	b.WithName(name)
	b.WithKind("IPAddress")
	b.WithAPIVersion("core.apinet.ironcore.dev/v1alpha1")
	return b
}

// ExtractIPAddress extracts the applied configuration owned by fieldManager from
// iPAddress. If no managedFields are found in iPAddress for fieldManager, a
// IPAddressApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// iPAddress must be a unmodified IPAddress API object that was retrieved from the Kubernetes API.
// ExtractIPAddress provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractIPAddress(iPAddress *corev1alpha1.IPAddress, fieldManager string) (*IPAddressApplyConfiguration, error) {
	return extractIPAddress(iPAddress, fieldManager, "")
}

// ExtractIPAddressStatus is the same as ExtractIPAddress except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractIPAddressStatus(iPAddress *corev1alpha1.IPAddress, fieldManager string) (*IPAddressApplyConfiguration, error) {
	return extractIPAddress(iPAddress, fieldManager, "status")
}

func extractIPAddress(iPAddress *corev1alpha1.IPAddress, fieldManager string, subresource string) (*IPAddressApplyConfiguration, error) {
	b := &IPAddressApplyConfiguration{}
	err := managedfields.ExtractInto(iPAddress, internal.Parser().Type("com.github.ironcore-dev.ironcore-net.api.core.v1alpha1.IPAddress"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(iPAddress.Name)

	b.WithKind("IPAddress")
	b.WithAPIVersion("core.apinet.ironcore.dev/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithKind(value string) *IPAddressApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithAPIVersion(value string) *IPAddressApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithName(value string) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithGenerateName(value string) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithNamespace(value string) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithUID(value types.UID) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithResourceVersion(value string) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithGeneration(value int64) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithCreationTimestamp(value metav1.Time) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *IPAddressApplyConfiguration) WithLabels(entries map[string]string) *IPAddressApplyConfiguration {
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
func (b *IPAddressApplyConfiguration) WithAnnotations(entries map[string]string) *IPAddressApplyConfiguration {
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
func (b *IPAddressApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *IPAddressApplyConfiguration {
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
func (b *IPAddressApplyConfiguration) WithFinalizers(values ...string) *IPAddressApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *IPAddressApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *IPAddressApplyConfiguration) WithSpec(value *IPAddressSpecApplyConfiguration) *IPAddressApplyConfiguration {
	b.Spec = value
	return b
}
