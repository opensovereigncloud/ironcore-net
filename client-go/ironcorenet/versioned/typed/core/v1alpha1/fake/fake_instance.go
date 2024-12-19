// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	corev1alpha1 "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstances implements InstanceInterface
type FakeInstances struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var instancesResource = v1alpha1.SchemeGroupVersion.WithResource("instances")

var instancesKind = v1alpha1.SchemeGroupVersion.WithKind("Instance")

// Get takes name of the instance, and returns the corresponding instance object, and an error if there is any.
func (c *FakeInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Instance, err error) {
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(instancesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}

// List takes label and field selectors, and returns the list of Instances that match those selectors.
func (c *FakeInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstanceList, err error) {
	emptyResult := &v1alpha1.InstanceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(instancesResource, instancesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstanceList{ListMeta: obj.(*v1alpha1.InstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instances.
func (c *FakeInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(instancesResource, c.ns, opts))

}

// Create takes the representation of a instance and creates it.  Returns the server's representation of the instance, and an error, if there is any.
func (c *FakeInstances) Create(ctx context.Context, instance *v1alpha1.Instance, opts v1.CreateOptions) (result *v1alpha1.Instance, err error) {
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(instancesResource, c.ns, instance, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}

// Update takes the representation of a instance and updates it. Returns the server's representation of the instance, and an error, if there is any.
func (c *FakeInstances) Update(ctx context.Context, instance *v1alpha1.Instance, opts v1.UpdateOptions) (result *v1alpha1.Instance, err error) {
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(instancesResource, c.ns, instance, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstances) UpdateStatus(ctx context.Context, instance *v1alpha1.Instance, opts v1.UpdateOptions) (result *v1alpha1.Instance, err error) {
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(instancesResource, "status", c.ns, instance, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}

// Delete takes name of the instance and deletes it. Returns an error if one occurs.
func (c *FakeInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(instancesResource, c.ns, name, opts), &v1alpha1.Instance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(instancesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstanceList{})
	return err
}

// Patch applies the patch and returns the patched instance.
func (c *FakeInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Instance, err error) {
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(instancesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied instance.
func (c *FakeInstances) Apply(ctx context.Context, instance *corev1alpha1.InstanceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Instance, err error) {
	if instance == nil {
		return nil, fmt.Errorf("instance provided to Apply must not be nil")
	}
	data, err := json.Marshal(instance)
	if err != nil {
		return nil, err
	}
	name := instance.Name
	if name == nil {
		return nil, fmt.Errorf("instance.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(instancesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeInstances) ApplyStatus(ctx context.Context, instance *corev1alpha1.InstanceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Instance, err error) {
	if instance == nil {
		return nil, fmt.Errorf("instance provided to Apply must not be nil")
	}
	data, err := json.Marshal(instance)
	if err != nil {
		return nil, err
	}
	name := instance.Name
	if name == nil {
		return nil, fmt.Errorf("instance.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Instance{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(instancesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Instance), err
}
