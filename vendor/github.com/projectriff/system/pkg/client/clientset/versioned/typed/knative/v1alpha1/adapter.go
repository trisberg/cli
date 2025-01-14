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
	v1alpha1 "github.com/projectriff/system/pkg/apis/knative/v1alpha1"
	scheme "github.com/projectriff/system/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AdaptersGetter has a method to return a AdapterInterface.
// A group's client should implement this interface.
type AdaptersGetter interface {
	Adapters(namespace string) AdapterInterface
}

// AdapterInterface has methods to work with Adapter resources.
type AdapterInterface interface {
	Create(*v1alpha1.Adapter) (*v1alpha1.Adapter, error)
	Update(*v1alpha1.Adapter) (*v1alpha1.Adapter, error)
	UpdateStatus(*v1alpha1.Adapter) (*v1alpha1.Adapter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Adapter, error)
	List(opts v1.ListOptions) (*v1alpha1.AdapterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Adapter, err error)
	AdapterExpansion
}

// adapters implements AdapterInterface
type adapters struct {
	client rest.Interface
	ns     string
}

// newAdapters returns a Adapters
func newAdapters(c *KnativeV1alpha1Client, namespace string) *adapters {
	return &adapters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the adapter, and returns the corresponding adapter object, and an error if there is any.
func (c *adapters) Get(name string, options v1.GetOptions) (result *v1alpha1.Adapter, err error) {
	result = &v1alpha1.Adapter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("adapters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Adapters that match those selectors.
func (c *adapters) List(opts v1.ListOptions) (result *v1alpha1.AdapterList, err error) {
	result = &v1alpha1.AdapterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("adapters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested adapters.
func (c *adapters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("adapters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a adapter and creates it.  Returns the server's representation of the adapter, and an error, if there is any.
func (c *adapters) Create(adapter *v1alpha1.Adapter) (result *v1alpha1.Adapter, err error) {
	result = &v1alpha1.Adapter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("adapters").
		Body(adapter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a adapter and updates it. Returns the server's representation of the adapter, and an error, if there is any.
func (c *adapters) Update(adapter *v1alpha1.Adapter) (result *v1alpha1.Adapter, err error) {
	result = &v1alpha1.Adapter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("adapters").
		Name(adapter.Name).
		Body(adapter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *adapters) UpdateStatus(adapter *v1alpha1.Adapter) (result *v1alpha1.Adapter, err error) {
	result = &v1alpha1.Adapter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("adapters").
		Name(adapter.Name).
		SubResource("status").
		Body(adapter).
		Do().
		Into(result)
	return
}

// Delete takes name of the adapter and deletes it. Returns an error if one occurs.
func (c *adapters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("adapters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *adapters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("adapters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched adapter.
func (c *adapters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Adapter, err error) {
	result = &v1alpha1.Adapter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("adapters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
