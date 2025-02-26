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

// FakeStreams implements StreamInterface
type FakeStreams struct {
	Fake *FakeStreamingV1alpha1
	ns   string
}

var streamsResource = schema.GroupVersionResource{Group: "streaming.projectriff.io", Version: "v1alpha1", Resource: "streams"}

var streamsKind = schema.GroupVersionKind{Group: "streaming.projectriff.io", Version: "v1alpha1", Kind: "Stream"}

// Get takes name of the stream, and returns the corresponding stream object, and an error if there is any.
func (c *FakeStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.Stream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(streamsResource, c.ns, name), &v1alpha1.Stream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stream), err
}

// List takes label and field selectors, and returns the list of Streams that match those selectors.
func (c *FakeStreams) List(opts v1.ListOptions) (result *v1alpha1.StreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(streamsResource, streamsKind, c.ns, opts), &v1alpha1.StreamList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamList{ListMeta: obj.(*v1alpha1.StreamList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streams.
func (c *FakeStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(streamsResource, c.ns, opts))

}

// Create takes the representation of a stream and creates it.  Returns the server's representation of the stream, and an error, if there is any.
func (c *FakeStreams) Create(stream *v1alpha1.Stream) (result *v1alpha1.Stream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(streamsResource, c.ns, stream), &v1alpha1.Stream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stream), err
}

// Update takes the representation of a stream and updates it. Returns the server's representation of the stream, and an error, if there is any.
func (c *FakeStreams) Update(stream *v1alpha1.Stream) (result *v1alpha1.Stream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(streamsResource, c.ns, stream), &v1alpha1.Stream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stream), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreams) UpdateStatus(stream *v1alpha1.Stream) (*v1alpha1.Stream, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(streamsResource, "status", c.ns, stream), &v1alpha1.Stream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stream), err
}

// Delete takes name of the stream and deletes it. Returns an error if one occurs.
func (c *FakeStreams) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(streamsResource, c.ns, name), &v1alpha1.Stream{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(streamsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamList{})
	return err
}

// Patch applies the patch and returns the patched stream.
func (c *FakeStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Stream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(streamsResource, c.ns, name, data, subresources...), &v1alpha1.Stream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stream), err
}
