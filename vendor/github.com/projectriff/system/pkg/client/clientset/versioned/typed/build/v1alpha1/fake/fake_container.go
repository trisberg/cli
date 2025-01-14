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
	v1alpha1 "github.com/projectriff/system/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainers implements ContainerInterface
type FakeContainers struct {
	Fake *FakeBuildV1alpha1
	ns   string
}

var containersResource = schema.GroupVersionResource{Group: "build.projectriff.io", Version: "v1alpha1", Resource: "containers"}

var containersKind = schema.GroupVersionKind{Group: "build.projectriff.io", Version: "v1alpha1", Kind: "Container"}

// Get takes name of the container, and returns the corresponding container object, and an error if there is any.
func (c *FakeContainers) Get(name string, options v1.GetOptions) (result *v1alpha1.Container, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containersResource, c.ns, name), &v1alpha1.Container{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Container), err
}

// List takes label and field selectors, and returns the list of Containers that match those selectors.
func (c *FakeContainers) List(opts v1.ListOptions) (result *v1alpha1.ContainerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(containersResource, containersKind, c.ns, opts), &v1alpha1.ContainerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ContainerList{ListMeta: obj.(*v1alpha1.ContainerList).ListMeta}
	for _, item := range obj.(*v1alpha1.ContainerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containers.
func (c *FakeContainers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containersResource, c.ns, opts))

}

// Create takes the representation of a container and creates it.  Returns the server's representation of the container, and an error, if there is any.
func (c *FakeContainers) Create(container *v1alpha1.Container) (result *v1alpha1.Container, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(containersResource, c.ns, container), &v1alpha1.Container{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Container), err
}

// Update takes the representation of a container and updates it. Returns the server's representation of the container, and an error, if there is any.
func (c *FakeContainers) Update(container *v1alpha1.Container) (result *v1alpha1.Container, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(containersResource, c.ns, container), &v1alpha1.Container{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Container), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainers) UpdateStatus(container *v1alpha1.Container) (*v1alpha1.Container, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(containersResource, "status", c.ns, container), &v1alpha1.Container{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Container), err
}

// Delete takes name of the container and deletes it. Returns an error if one occurs.
func (c *FakeContainers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(containersResource, c.ns, name), &v1alpha1.Container{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(containersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ContainerList{})
	return err
}

// Patch applies the patch and returns the patched container.
func (c *FakeContainers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Container, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(containersResource, c.ns, name, data, subresources...), &v1alpha1.Container{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Container), err
}
