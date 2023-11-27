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
	v1 "k8s.io/api/core/v1"
)

// NATGatewaySpecApplyConfiguration represents an declarative configuration of the NATGatewaySpec type for use
// with apply.
type NATGatewaySpecApplyConfiguration struct {
	IPFamily                 *v1.IPFamily                     `json:"ipFamily,omitempty"`
	NetworkRef               *v1.LocalObjectReference         `json:"networkRef,omitempty"`
	IPs                      []NATGatewayIPApplyConfiguration `json:"ips,omitempty"`
	PortsPerNetworkInterface *int32                           `json:"portsPerNetworkInterface,omitempty"`
}

// NATGatewaySpecApplyConfiguration constructs an declarative configuration of the NATGatewaySpec type for use with
// apply.
func NATGatewaySpec() *NATGatewaySpecApplyConfiguration {
	return &NATGatewaySpecApplyConfiguration{}
}

// WithIPFamily sets the IPFamily field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPFamily field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithIPFamily(value v1.IPFamily) *NATGatewaySpecApplyConfiguration {
	b.IPFamily = &value
	return b
}

// WithNetworkRef sets the NetworkRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkRef field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithNetworkRef(value v1.LocalObjectReference) *NATGatewaySpecApplyConfiguration {
	b.NetworkRef = &value
	return b
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *NATGatewaySpecApplyConfiguration) WithIPs(values ...*NATGatewayIPApplyConfiguration) *NATGatewaySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIPs")
		}
		b.IPs = append(b.IPs, *values[i])
	}
	return b
}

// WithPortsPerNetworkInterface sets the PortsPerNetworkInterface field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortsPerNetworkInterface field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithPortsPerNetworkInterface(value int32) *NATGatewaySpecApplyConfiguration {
	b.PortsPerNetworkInterface = &value
	return b
}
