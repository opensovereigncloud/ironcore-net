/*
 * Copyright (c) 2022 by the OnMetal authors.
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
	v1alpha1 "github.com/onmetal/onmetal-api-net/api/core/v1alpha1"
	net "github.com/onmetal/onmetal-api-net/apimachinery/api/net"
)

// NetworkInterfaceStatusApplyConfiguration represents an declarative configuration of the NetworkInterfaceStatus type for use
// with apply.
type NetworkInterfaceStatusApplyConfiguration struct {
	State      *v1alpha1.NetworkInterfaceState `json:"state,omitempty"`
	PCIAddress *PCIAddressApplyConfiguration   `json:"pciAddress,omitempty"`
	Prefixes   []net.IPPrefix                  `json:"prefixes,omitempty"`
	PublicIPs  []net.IP                        `json:"publicIPs,omitempty"`
	NATIPs     []net.IP                        `json:"natIPs,omitempty"`
}

// NetworkInterfaceStatusApplyConfiguration constructs an declarative configuration of the NetworkInterfaceStatus type for use with
// apply.
func NetworkInterfaceStatus() *NetworkInterfaceStatusApplyConfiguration {
	return &NetworkInterfaceStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *NetworkInterfaceStatusApplyConfiguration) WithState(value v1alpha1.NetworkInterfaceState) *NetworkInterfaceStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithPCIAddress sets the PCIAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PCIAddress field is set to the value of the last call.
func (b *NetworkInterfaceStatusApplyConfiguration) WithPCIAddress(value *PCIAddressApplyConfiguration) *NetworkInterfaceStatusApplyConfiguration {
	b.PCIAddress = value
	return b
}

// WithPrefixes adds the given value to the Prefixes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Prefixes field.
func (b *NetworkInterfaceStatusApplyConfiguration) WithPrefixes(values ...net.IPPrefix) *NetworkInterfaceStatusApplyConfiguration {
	for i := range values {
		b.Prefixes = append(b.Prefixes, values[i])
	}
	return b
}

// WithPublicIPs adds the given value to the PublicIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PublicIPs field.
func (b *NetworkInterfaceStatusApplyConfiguration) WithPublicIPs(values ...net.IP) *NetworkInterfaceStatusApplyConfiguration {
	for i := range values {
		b.PublicIPs = append(b.PublicIPs, values[i])
	}
	return b
}

// WithNATIPs adds the given value to the NATIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NATIPs field.
func (b *NetworkInterfaceStatusApplyConfiguration) WithNATIPs(values ...net.IP) *NetworkInterfaceStatusApplyConfiguration {
	for i := range values {
		b.NATIPs = append(b.NATIPs, values[i])
	}
	return b
}
