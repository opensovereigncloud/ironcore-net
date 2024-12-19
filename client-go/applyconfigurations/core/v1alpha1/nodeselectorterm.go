// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// NodeSelectorTermApplyConfiguration represents a declarative configuration of the NodeSelectorTerm type for use
// with apply.
type NodeSelectorTermApplyConfiguration struct {
	MatchExpressions []NodeSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
	MatchFields      []NodeSelectorRequirementApplyConfiguration `json:"matchFields,omitempty"`
}

// NodeSelectorTermApplyConfiguration constructs a declarative configuration of the NodeSelectorTerm type for use with
// apply.
func NodeSelectorTerm() *NodeSelectorTermApplyConfiguration {
	return &NodeSelectorTermApplyConfiguration{}
}

// WithMatchExpressions adds the given value to the MatchExpressions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MatchExpressions field.
func (b *NodeSelectorTermApplyConfiguration) WithMatchExpressions(values ...*NodeSelectorRequirementApplyConfiguration) *NodeSelectorTermApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchExpressions")
		}
		b.MatchExpressions = append(b.MatchExpressions, *values[i])
	}
	return b
}

// WithMatchFields adds the given value to the MatchFields field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MatchFields field.
func (b *NodeSelectorTermApplyConfiguration) WithMatchFields(values ...*NodeSelectorRequirementApplyConfiguration) *NodeSelectorTermApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchFields")
		}
		b.MatchFields = append(b.MatchFields, *values[i])
	}
	return b
}
