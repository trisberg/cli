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
	v1alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProcessors implements ProcessorInterface
type FakeProcessors struct {
	Fake *FakeStreamingV1alpha1
	ns   string
}

var processorsResource = schema.GroupVersionResource{Group: "streaming.projectriff.io", Version: "v1alpha1", Resource: "processors"}

var processorsKind = schema.GroupVersionKind{Group: "streaming.projectriff.io", Version: "v1alpha1", Kind: "Processor"}

// Get takes name of the processor, and returns the corresponding processor object, and an error if there is any.
func (c *FakeProcessors) Get(name string, options v1.GetOptions) (result *v1alpha1.Processor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(processorsResource, c.ns, name), &v1alpha1.Processor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Processor), err
}

// List takes label and field selectors, and returns the list of Processors that match those selectors.
func (c *FakeProcessors) List(opts v1.ListOptions) (result *v1alpha1.ProcessorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(processorsResource, processorsKind, c.ns, opts), &v1alpha1.ProcessorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProcessorList{ListMeta: obj.(*v1alpha1.ProcessorList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProcessorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested processors.
func (c *FakeProcessors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(processorsResource, c.ns, opts))

}

// Create takes the representation of a processor and creates it.  Returns the server's representation of the processor, and an error, if there is any.
func (c *FakeProcessors) Create(processor *v1alpha1.Processor) (result *v1alpha1.Processor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(processorsResource, c.ns, processor), &v1alpha1.Processor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Processor), err
}

// Update takes the representation of a processor and updates it. Returns the server's representation of the processor, and an error, if there is any.
func (c *FakeProcessors) Update(processor *v1alpha1.Processor) (result *v1alpha1.Processor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(processorsResource, c.ns, processor), &v1alpha1.Processor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Processor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProcessors) UpdateStatus(processor *v1alpha1.Processor) (*v1alpha1.Processor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(processorsResource, "status", c.ns, processor), &v1alpha1.Processor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Processor), err
}

// Delete takes name of the processor and deletes it. Returns an error if one occurs.
func (c *FakeProcessors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(processorsResource, c.ns, name), &v1alpha1.Processor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProcessors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(processorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProcessorList{})
	return err
}

// Patch applies the patch and returns the patched processor.
func (c *FakeProcessors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Processor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(processorsResource, c.ns, name, data, subresources...), &v1alpha1.Processor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Processor), err
}
