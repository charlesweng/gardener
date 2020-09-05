/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperatingSystemConfigs implements OperatingSystemConfigInterface
type FakeOperatingSystemConfigs struct {
	Fake *FakeExtensionsV1alpha1
	ns   string
}

var operatingsystemconfigsResource = schema.GroupVersionResource{Group: "extensions.gardener.cloud", Version: "v1alpha1", Resource: "operatingsystemconfigs"}

var operatingsystemconfigsKind = schema.GroupVersionKind{Group: "extensions.gardener.cloud", Version: "v1alpha1", Kind: "OperatingSystemConfig"}

// Get takes name of the operatingSystemConfig, and returns the corresponding operatingSystemConfig object, and an error if there is any.
func (c *FakeOperatingSystemConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OperatingSystemConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operatingsystemconfigsResource, c.ns, name), &v1alpha1.OperatingSystemConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperatingSystemConfig), err
}

// List takes label and field selectors, and returns the list of OperatingSystemConfigs that match those selectors.
func (c *FakeOperatingSystemConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OperatingSystemConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(operatingsystemconfigsResource, operatingsystemconfigsKind, c.ns, opts), &v1alpha1.OperatingSystemConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OperatingSystemConfigList{ListMeta: obj.(*v1alpha1.OperatingSystemConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.OperatingSystemConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operatingSystemConfigs.
func (c *FakeOperatingSystemConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(operatingsystemconfigsResource, c.ns, opts))

}

// Create takes the representation of a operatingSystemConfig and creates it.  Returns the server's representation of the operatingSystemConfig, and an error, if there is any.
func (c *FakeOperatingSystemConfigs) Create(ctx context.Context, operatingSystemConfig *v1alpha1.OperatingSystemConfig, opts v1.CreateOptions) (result *v1alpha1.OperatingSystemConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(operatingsystemconfigsResource, c.ns, operatingSystemConfig), &v1alpha1.OperatingSystemConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperatingSystemConfig), err
}

// Update takes the representation of a operatingSystemConfig and updates it. Returns the server's representation of the operatingSystemConfig, and an error, if there is any.
func (c *FakeOperatingSystemConfigs) Update(ctx context.Context, operatingSystemConfig *v1alpha1.OperatingSystemConfig, opts v1.UpdateOptions) (result *v1alpha1.OperatingSystemConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(operatingsystemconfigsResource, c.ns, operatingSystemConfig), &v1alpha1.OperatingSystemConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperatingSystemConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperatingSystemConfigs) UpdateStatus(ctx context.Context, operatingSystemConfig *v1alpha1.OperatingSystemConfig, opts v1.UpdateOptions) (*v1alpha1.OperatingSystemConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(operatingsystemconfigsResource, "status", c.ns, operatingSystemConfig), &v1alpha1.OperatingSystemConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperatingSystemConfig), err
}

// Delete takes name of the operatingSystemConfig and deletes it. Returns an error if one occurs.
func (c *FakeOperatingSystemConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(operatingsystemconfigsResource, c.ns, name), &v1alpha1.OperatingSystemConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperatingSystemConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(operatingsystemconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OperatingSystemConfigList{})
	return err
}

// Patch applies the patch and returns the patched operatingSystemConfig.
func (c *FakeOperatingSystemConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OperatingSystemConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(operatingsystemconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OperatingSystemConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperatingSystemConfig), err
}
