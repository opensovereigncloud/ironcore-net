// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	v1 "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/meta/v1"
)

// TopologySpreadConstraintApplyConfiguration represents an declarative configuration of the TopologySpreadConstraint type for use
// with apply.
type TopologySpreadConstraintApplyConfiguration struct {
	MaxSkew           *int32                                  `json:"maxSkew,omitempty"`
	TopologyKey       *string                                 `json:"topologyKey,omitempty"`
	WhenUnsatisfiable *v1alpha1.UnsatisfiableConstraintAction `json:"whenUnsatisfiable,omitempty"`
	LabelSelector     *v1.LabelSelectorApplyConfiguration     `json:"labelSelector,omitempty"`
}

// TopologySpreadConstraintApplyConfiguration constructs an declarative configuration of the TopologySpreadConstraint type for use with
// apply.
func TopologySpreadConstraint() *TopologySpreadConstraintApplyConfiguration {
	return &TopologySpreadConstraintApplyConfiguration{}
}

// WithMaxSkew sets the MaxSkew field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxSkew field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithMaxSkew(value int32) *TopologySpreadConstraintApplyConfiguration {
	b.MaxSkew = &value
	return b
}

// WithTopologyKey sets the TopologyKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TopologyKey field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithTopologyKey(value string) *TopologySpreadConstraintApplyConfiguration {
	b.TopologyKey = &value
	return b
}

// WithWhenUnsatisfiable sets the WhenUnsatisfiable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WhenUnsatisfiable field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithWhenUnsatisfiable(value v1alpha1.UnsatisfiableConstraintAction) *TopologySpreadConstraintApplyConfiguration {
	b.WhenUnsatisfiable = &value
	return b
}

// WithLabelSelector sets the LabelSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelSelector field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithLabelSelector(value *v1.LabelSelectorApplyConfiguration) *TopologySpreadConstraintApplyConfiguration {
	b.LabelSelector = value
	return b
}
