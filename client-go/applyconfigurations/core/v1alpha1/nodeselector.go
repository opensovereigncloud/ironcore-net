// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// NodeSelectorApplyConfiguration represents a declarative configuration of the NodeSelector type for use
// with apply.
type NodeSelectorApplyConfiguration struct {
	NodeSelectorTerms []NodeSelectorTermApplyConfiguration `json:"nodeSelectorTerms,omitempty"`
}

// NodeSelectorApplyConfiguration constructs a declarative configuration of the NodeSelector type for use with
// apply.
func NodeSelector() *NodeSelectorApplyConfiguration {
	return &NodeSelectorApplyConfiguration{}
}

// WithNodeSelectorTerms adds the given value to the NodeSelectorTerms field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeSelectorTerms field.
func (b *NodeSelectorApplyConfiguration) WithNodeSelectorTerms(values ...*NodeSelectorTermApplyConfiguration) *NodeSelectorApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNodeSelectorTerms")
		}
		b.NodeSelectorTerms = append(b.NodeSelectorTerms, *values[i])
	}
	return b
}
