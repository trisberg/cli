/*
 * Copyright 2019 The original author or authors
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
package fake

import (
	v1alpha1 "github.com/projectriff/system/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeployers implements DeployerInterface
type FakeDeployers struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var deployersResource = schema.GroupVersionResource{Group: "core.projectriff.io", Version: "v1alpha1", Resource: "deployers"}

var deployersKind = schema.GroupVersionKind{Group: "core.projectriff.io", Version: "v1alpha1", Kind: "Deployer"}

// Get takes name of the deployer, and returns the corresponding deployer object, and an error if there is any.
func (c *FakeDeployers) Get(name string, options v1.GetOptions) (result *v1alpha1.Deployer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deployersResource, c.ns, name), &v1alpha1.Deployer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployer), err
}

// List takes label and field selectors, and returns the list of Deployers that match those selectors.
func (c *FakeDeployers) List(opts v1.ListOptions) (result *v1alpha1.DeployerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deployersResource, deployersKind, c.ns, opts), &v1alpha1.DeployerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeployerList{ListMeta: obj.(*v1alpha1.DeployerList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeployerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployers.
func (c *FakeDeployers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deployersResource, c.ns, opts))

}

// Create takes the representation of a deployer and creates it.  Returns the server's representation of the deployer, and an error, if there is any.
func (c *FakeDeployers) Create(deployer *v1alpha1.Deployer) (result *v1alpha1.Deployer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deployersResource, c.ns, deployer), &v1alpha1.Deployer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployer), err
}

// Update takes the representation of a deployer and updates it. Returns the server's representation of the deployer, and an error, if there is any.
func (c *FakeDeployers) Update(deployer *v1alpha1.Deployer) (result *v1alpha1.Deployer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deployersResource, c.ns, deployer), &v1alpha1.Deployer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeployers) UpdateStatus(deployer *v1alpha1.Deployer) (*v1alpha1.Deployer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deployersResource, "status", c.ns, deployer), &v1alpha1.Deployer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployer), err
}

// Delete takes name of the deployer and deletes it. Returns an error if one occurs.
func (c *FakeDeployers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deployersResource, c.ns, name), &v1alpha1.Deployer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deployersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeployerList{})
	return err
}

// Patch applies the patch and returns the patched deployer.
func (c *FakeDeployers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Deployer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deployersResource, c.ns, name, data, subresources...), &v1alpha1.Deployer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployer), err
}
