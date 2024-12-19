// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	corev1alpha1 "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/core/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore-net/client-go/ironcorenet/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NetworkPolicyRulesGetter has a method to return a NetworkPolicyRuleInterface.
// A group's client should implement this interface.
type NetworkPolicyRulesGetter interface {
	NetworkPolicyRules(namespace string) NetworkPolicyRuleInterface
}

// NetworkPolicyRuleInterface has methods to work with NetworkPolicyRule resources.
type NetworkPolicyRuleInterface interface {
	Create(ctx context.Context, networkPolicyRule *v1alpha1.NetworkPolicyRule, opts v1.CreateOptions) (*v1alpha1.NetworkPolicyRule, error)
	Update(ctx context.Context, networkPolicyRule *v1alpha1.NetworkPolicyRule, opts v1.UpdateOptions) (*v1alpha1.NetworkPolicyRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NetworkPolicyRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NetworkPolicyRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkPolicyRule, err error)
	Apply(ctx context.Context, networkPolicyRule *corev1alpha1.NetworkPolicyRuleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NetworkPolicyRule, err error)
	NetworkPolicyRuleExpansion
}

// networkPolicyRules implements NetworkPolicyRuleInterface
type networkPolicyRules struct {
	*gentype.ClientWithListAndApply[*v1alpha1.NetworkPolicyRule, *v1alpha1.NetworkPolicyRuleList, *corev1alpha1.NetworkPolicyRuleApplyConfiguration]
}

// newNetworkPolicyRules returns a NetworkPolicyRules
func newNetworkPolicyRules(c *CoreV1alpha1Client, namespace string) *networkPolicyRules {
	return &networkPolicyRules{
		gentype.NewClientWithListAndApply[*v1alpha1.NetworkPolicyRule, *v1alpha1.NetworkPolicyRuleList, *corev1alpha1.NetworkPolicyRuleApplyConfiguration](
			"networkpolicyrules",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.NetworkPolicyRule { return &v1alpha1.NetworkPolicyRule{} },
			func() *v1alpha1.NetworkPolicyRuleList { return &v1alpha1.NetworkPolicyRuleList{} }),
	}
}
