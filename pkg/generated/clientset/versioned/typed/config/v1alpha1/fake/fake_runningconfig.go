/*
Copyright 2024 Nokia.

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

	v1alpha1 "github.com/sdcio/config-server/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRunningConfigs implements RunningConfigInterface
type FakeRunningConfigs struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var runningconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("runningconfigs")

var runningconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("RunningConfig")

// Get takes name of the runningConfig, and returns the corresponding runningConfig object, and an error if there is any.
func (c *FakeRunningConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RunningConfig, err error) {
	emptyResult := &v1alpha1.RunningConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(runningconfigsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RunningConfig), err
}

// List takes label and field selectors, and returns the list of RunningConfigs that match those selectors.
func (c *FakeRunningConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RunningConfigList, err error) {
	emptyResult := &v1alpha1.RunningConfigList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(runningconfigsResource, runningconfigsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RunningConfigList{ListMeta: obj.(*v1alpha1.RunningConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.RunningConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runningConfigs.
func (c *FakeRunningConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(runningconfigsResource, c.ns, opts))

}

// Create takes the representation of a runningConfig and creates it.  Returns the server's representation of the runningConfig, and an error, if there is any.
func (c *FakeRunningConfigs) Create(ctx context.Context, runningConfig *v1alpha1.RunningConfig, opts v1.CreateOptions) (result *v1alpha1.RunningConfig, err error) {
	emptyResult := &v1alpha1.RunningConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(runningconfigsResource, c.ns, runningConfig, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RunningConfig), err
}

// Update takes the representation of a runningConfig and updates it. Returns the server's representation of the runningConfig, and an error, if there is any.
func (c *FakeRunningConfigs) Update(ctx context.Context, runningConfig *v1alpha1.RunningConfig, opts v1.UpdateOptions) (result *v1alpha1.RunningConfig, err error) {
	emptyResult := &v1alpha1.RunningConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(runningconfigsResource, c.ns, runningConfig, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RunningConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRunningConfigs) UpdateStatus(ctx context.Context, runningConfig *v1alpha1.RunningConfig, opts v1.UpdateOptions) (result *v1alpha1.RunningConfig, err error) {
	emptyResult := &v1alpha1.RunningConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(runningconfigsResource, "status", c.ns, runningConfig, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RunningConfig), err
}

// Delete takes name of the runningConfig and deletes it. Returns an error if one occurs.
func (c *FakeRunningConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(runningconfigsResource, c.ns, name, opts), &v1alpha1.RunningConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRunningConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(runningconfigsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RunningConfigList{})
	return err
}

// Patch applies the patch and returns the patched runningConfig.
func (c *FakeRunningConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RunningConfig, err error) {
	emptyResult := &v1alpha1.RunningConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(runningconfigsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RunningConfig), err
}
