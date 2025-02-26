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
package v1alpha1

import (
	v1alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
	scheme "github.com/projectriff/system/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StreamsGetter has a method to return a StreamInterface.
// A group's client should implement this interface.
type StreamsGetter interface {
	Streams(namespace string) StreamInterface
}

// StreamInterface has methods to work with Stream resources.
type StreamInterface interface {
	Create(*v1alpha1.Stream) (*v1alpha1.Stream, error)
	Update(*v1alpha1.Stream) (*v1alpha1.Stream, error)
	UpdateStatus(*v1alpha1.Stream) (*v1alpha1.Stream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Stream, error)
	List(opts v1.ListOptions) (*v1alpha1.StreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Stream, err error)
	StreamExpansion
}

// streams implements StreamInterface
type streams struct {
	client rest.Interface
	ns     string
}

// newStreams returns a Streams
func newStreams(c *StreamingV1alpha1Client, namespace string) *streams {
	return &streams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the stream, and returns the corresponding stream object, and an error if there is any.
func (c *streams) Get(name string, options v1.GetOptions) (result *v1alpha1.Stream, err error) {
	result = &v1alpha1.Stream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("streams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Streams that match those selectors.
func (c *streams) List(opts v1.ListOptions) (result *v1alpha1.StreamList, err error) {
	result = &v1alpha1.StreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("streams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested streams.
func (c *streams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("streams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a stream and creates it.  Returns the server's representation of the stream, and an error, if there is any.
func (c *streams) Create(stream *v1alpha1.Stream) (result *v1alpha1.Stream, err error) {
	result = &v1alpha1.Stream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("streams").
		Body(stream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a stream and updates it. Returns the server's representation of the stream, and an error, if there is any.
func (c *streams) Update(stream *v1alpha1.Stream) (result *v1alpha1.Stream, err error) {
	result = &v1alpha1.Stream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("streams").
		Name(stream.Name).
		Body(stream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *streams) UpdateStatus(stream *v1alpha1.Stream) (result *v1alpha1.Stream, err error) {
	result = &v1alpha1.Stream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("streams").
		Name(stream.Name).
		SubResource("status").
		Body(stream).
		Do().
		Into(result)
	return
}

// Delete takes name of the stream and deletes it. Returns an error if one occurs.
func (c *streams) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("streams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *streams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("streams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched stream.
func (c *streams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Stream, err error) {
	result = &v1alpha1.Stream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("streams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
