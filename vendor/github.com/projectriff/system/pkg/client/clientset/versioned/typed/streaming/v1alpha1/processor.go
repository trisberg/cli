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

// ProcessorsGetter has a method to return a ProcessorInterface.
// A group's client should implement this interface.
type ProcessorsGetter interface {
	Processors(namespace string) ProcessorInterface
}

// ProcessorInterface has methods to work with Processor resources.
type ProcessorInterface interface {
	Create(*v1alpha1.Processor) (*v1alpha1.Processor, error)
	Update(*v1alpha1.Processor) (*v1alpha1.Processor, error)
	UpdateStatus(*v1alpha1.Processor) (*v1alpha1.Processor, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Processor, error)
	List(opts v1.ListOptions) (*v1alpha1.ProcessorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Processor, err error)
	ProcessorExpansion
}

// processors implements ProcessorInterface
type processors struct {
	client rest.Interface
	ns     string
}

// newProcessors returns a Processors
func newProcessors(c *StreamingV1alpha1Client, namespace string) *processors {
	return &processors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the processor, and returns the corresponding processor object, and an error if there is any.
func (c *processors) Get(name string, options v1.GetOptions) (result *v1alpha1.Processor, err error) {
	result = &v1alpha1.Processor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("processors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Processors that match those selectors.
func (c *processors) List(opts v1.ListOptions) (result *v1alpha1.ProcessorList, err error) {
	result = &v1alpha1.ProcessorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("processors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested processors.
func (c *processors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("processors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a processor and creates it.  Returns the server's representation of the processor, and an error, if there is any.
func (c *processors) Create(processor *v1alpha1.Processor) (result *v1alpha1.Processor, err error) {
	result = &v1alpha1.Processor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("processors").
		Body(processor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a processor and updates it. Returns the server's representation of the processor, and an error, if there is any.
func (c *processors) Update(processor *v1alpha1.Processor) (result *v1alpha1.Processor, err error) {
	result = &v1alpha1.Processor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("processors").
		Name(processor.Name).
		Body(processor).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *processors) UpdateStatus(processor *v1alpha1.Processor) (result *v1alpha1.Processor, err error) {
	result = &v1alpha1.Processor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("processors").
		Name(processor.Name).
		SubResource("status").
		Body(processor).
		Do().
		Into(result)
	return
}

// Delete takes name of the processor and deletes it. Returns an error if one occurs.
func (c *processors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("processors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *processors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("processors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched processor.
func (c *processors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Processor, err error) {
	result = &v1alpha1.Processor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("processors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
