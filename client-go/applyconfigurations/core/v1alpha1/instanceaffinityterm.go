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
	v1 "github.com/onmetal/onmetal-api-net/client-go/applyconfigurations/meta/v1"
)

// InstanceAffinityTermApplyConfiguration represents an declarative configuration of the InstanceAffinityTerm type for use
// with apply.
type InstanceAffinityTermApplyConfiguration struct {
	LabelSelector *v1.LabelSelectorApplyConfiguration `json:"labelSelector,omitempty"`
	TopologyKey   *string                             `json:"topologyKey,omitempty"`
}

// InstanceAffinityTermApplyConfiguration constructs an declarative configuration of the InstanceAffinityTerm type for use with
// apply.
func InstanceAffinityTerm() *InstanceAffinityTermApplyConfiguration {
	return &InstanceAffinityTermApplyConfiguration{}
}

// WithLabelSelector sets the LabelSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelSelector field is set to the value of the last call.
func (b *InstanceAffinityTermApplyConfiguration) WithLabelSelector(value *v1.LabelSelectorApplyConfiguration) *InstanceAffinityTermApplyConfiguration {
	b.LabelSelector = value
	return b
}

// WithTopologyKey sets the TopologyKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TopologyKey field is set to the value of the last call.
func (b *InstanceAffinityTermApplyConfiguration) WithTopologyKey(value string) *InstanceAffinityTermApplyConfiguration {
	b.TopologyKey = &value
	return b
}
