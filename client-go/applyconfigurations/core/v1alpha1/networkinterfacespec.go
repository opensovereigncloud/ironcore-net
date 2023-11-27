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
	net "github.com/ironcore-dev/ironcore-net/apimachinery/api/net"
	v1 "k8s.io/api/core/v1"
)

// NetworkInterfaceSpecApplyConfiguration represents an declarative configuration of the NetworkInterfaceSpec type for use
// with apply.
type NetworkInterfaceSpecApplyConfiguration struct {
	NodeRef    *v1.LocalObjectReference                     `json:"nodeRef,omitempty"`
	NetworkRef *v1.LocalObjectReference                     `json:"networkRef,omitempty"`
	IPs        []net.IP                                     `json:"ips,omitempty"`
	Prefixes   []net.IPPrefix                               `json:"prefixes,omitempty"`
	NATs       []NetworkInterfaceNATApplyConfiguration      `json:"natGateways,omitempty"`
	PublicIPs  []NetworkInterfacePublicIPApplyConfiguration `json:"publicIPs,omitempty"`
}

// NetworkInterfaceSpecApplyConfiguration constructs an declarative configuration of the NetworkInterfaceSpec type for use with
// apply.
func NetworkInterfaceSpec() *NetworkInterfaceSpecApplyConfiguration {
	return &NetworkInterfaceSpecApplyConfiguration{}
}

// WithNodeRef sets the NodeRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeRef field is set to the value of the last call.
func (b *NetworkInterfaceSpecApplyConfiguration) WithNodeRef(value v1.LocalObjectReference) *NetworkInterfaceSpecApplyConfiguration {
	b.NodeRef = &value
	return b
}

// WithNetworkRef sets the NetworkRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkRef field is set to the value of the last call.
func (b *NetworkInterfaceSpecApplyConfiguration) WithNetworkRef(value v1.LocalObjectReference) *NetworkInterfaceSpecApplyConfiguration {
	b.NetworkRef = &value
	return b
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *NetworkInterfaceSpecApplyConfiguration) WithIPs(values ...net.IP) *NetworkInterfaceSpecApplyConfiguration {
	for i := range values {
		b.IPs = append(b.IPs, values[i])
	}
	return b
}

// WithPrefixes adds the given value to the Prefixes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Prefixes field.
func (b *NetworkInterfaceSpecApplyConfiguration) WithPrefixes(values ...net.IPPrefix) *NetworkInterfaceSpecApplyConfiguration {
	for i := range values {
		b.Prefixes = append(b.Prefixes, values[i])
	}
	return b
}

// WithNATs adds the given value to the NATs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NATs field.
func (b *NetworkInterfaceSpecApplyConfiguration) WithNATs(values ...*NetworkInterfaceNATApplyConfiguration) *NetworkInterfaceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNATs")
		}
		b.NATs = append(b.NATs, *values[i])
	}
	return b
}

// WithPublicIPs adds the given value to the PublicIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PublicIPs field.
func (b *NetworkInterfaceSpecApplyConfiguration) WithPublicIPs(values ...*NetworkInterfacePublicIPApplyConfiguration) *NetworkInterfaceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPublicIPs")
		}
		b.PublicIPs = append(b.PublicIPs, *values[i])
	}
	return b
}
