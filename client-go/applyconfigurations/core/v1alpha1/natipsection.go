// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	net "github.com/ironcore-dev/ironcore-net/apimachinery/api/net"
)

// NATIPSectionApplyConfiguration represents a declarative configuration of the NATIPSection type for use
// with apply.
type NATIPSectionApplyConfiguration struct {
	IP        *net.IP                                `json:"ip,omitempty"`
	Port      *int32                                 `json:"port,omitempty"`
	EndPort   *int32                                 `json:"endPort,omitempty"`
	TargetRef *NATTableIPTargetRefApplyConfiguration `json:"targetRef,omitempty"`
}

// NATIPSectionApplyConfiguration constructs a declarative configuration of the NATIPSection type for use with
// apply.
func NATIPSection() *NATIPSectionApplyConfiguration {
	return &NATIPSectionApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *NATIPSectionApplyConfiguration) WithIP(value net.IP) *NATIPSectionApplyConfiguration {
	b.IP = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *NATIPSectionApplyConfiguration) WithPort(value int32) *NATIPSectionApplyConfiguration {
	b.Port = &value
	return b
}

// WithEndPort sets the EndPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndPort field is set to the value of the last call.
func (b *NATIPSectionApplyConfiguration) WithEndPort(value int32) *NATIPSectionApplyConfiguration {
	b.EndPort = &value
	return b
}

// WithTargetRef sets the TargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetRef field is set to the value of the last call.
func (b *NATIPSectionApplyConfiguration) WithTargetRef(value *NATTableIPTargetRefApplyConfiguration) *NATIPSectionApplyConfiguration {
	b.TargetRef = value
	return b
}
