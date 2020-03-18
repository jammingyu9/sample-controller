/*
Copyright The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
)

// FakeCustomDeployments implements CustomDeploymentInterface
type FakeCustomDeployments struct {
	Fake *FakeSamplecontrollerV1alpha1
	ns   string
}

var customdeploymentsResource = schema.GroupVersionResource{Group: "samplecontroller.k8s.io", Version: "v1alpha1", Resource: "customdeployments"}

var customdeploymentsKind = schema.GroupVersionKind{Group: "samplecontroller.k8s.io", Version: "v1alpha1", Kind: "CustomDeployment"}

// Get takes name of the customDeployment, and returns the corresponding customDeployment object, and an error if there is any.
func (c *FakeCustomDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(customdeploymentsResource, c.ns, name), &v1alpha1.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomDeployment), err
}

// List takes label and field selectors, and returns the list of CustomDeployments that match those selectors.
func (c *FakeCustomDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(customdeploymentsResource, customdeploymentsKind, c.ns, opts), &v1alpha1.CustomDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomDeploymentList{ListMeta: obj.(*v1alpha1.CustomDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customDeployments.
func (c *FakeCustomDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(customdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a customDeployment and creates it.  Returns the server's representation of the customDeployment, and an error, if there is any.
func (c *FakeCustomDeployments) Create(ctx context.Context, customDeployment *v1alpha1.CustomDeployment, opts v1.CreateOptions) (result *v1alpha1.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(customdeploymentsResource, c.ns, customDeployment), &v1alpha1.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomDeployment), err
}

// Update takes the representation of a customDeployment and updates it. Returns the server's representation of the customDeployment, and an error, if there is any.
func (c *FakeCustomDeployments) Update(ctx context.Context, customDeployment *v1alpha1.CustomDeployment, opts v1.UpdateOptions) (result *v1alpha1.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(customdeploymentsResource, c.ns, customDeployment), &v1alpha1.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomDeployments) UpdateStatus(ctx context.Context, customDeployment *v1alpha1.CustomDeployment, opts v1.UpdateOptions) (*v1alpha1.CustomDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(customdeploymentsResource, "status", c.ns, customDeployment), &v1alpha1.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomDeployment), err
}

// Delete takes name of the customDeployment and deletes it. Returns an error if one occurs.
func (c *FakeCustomDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(customdeploymentsResource, c.ns, name), &v1alpha1.CustomDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(customdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched customDeployment.
func (c *FakeCustomDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(customdeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CustomDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomDeployment), err
}
