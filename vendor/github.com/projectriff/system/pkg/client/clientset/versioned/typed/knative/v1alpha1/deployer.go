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

// DeployersGetter has a method to return a DeployerInterface.
// A group's client should implement this interface.
type DeployersGetter interface {
	Deployers(namespace string) DeployerInterface
}

// DeployerInterface has methods to work with Deployer resources.
type DeployerInterface interface {
	Create(*v1alpha1.Deployer) (*v1alpha1.Deployer, error)
	Update(*v1alpha1.Deployer) (*v1alpha1.Deployer, error)
	UpdateStatus(*v1alpha1.Deployer) (*v1alpha1.Deployer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Deployer, error)
	List(opts v1.ListOptions) (*v1alpha1.DeployerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Deployer, err error)
	DeployerExpansion
}

// deployers implements DeployerInterface
type deployers struct {
	client rest.Interface
	ns     string
}

// newDeployers returns a Deployers
func newDeployers(c *KnativeV1alpha1Client, namespace string) *deployers {
	return &deployers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployer, and returns the corresponding deployer object, and an error if there is any.
func (c *deployers) Get(name string, options v1.GetOptions) (result *v1alpha1.Deployer, err error) {
	result = &v1alpha1.Deployer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployers that match those selectors.
func (c *deployers) List(opts v1.ListOptions) (result *v1alpha1.DeployerList, err error) {
	result = &v1alpha1.DeployerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deployers.
func (c *deployers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deployers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a deployer and creates it.  Returns the server's representation of the deployer, and an error, if there is any.
func (c *deployers) Create(deployer *v1alpha1.Deployer) (result *v1alpha1.Deployer, err error) {
	result = &v1alpha1.Deployer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deployers").
		Body(deployer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deployer and updates it. Returns the server's representation of the deployer, and an error, if there is any.
func (c *deployers) Update(deployer *v1alpha1.Deployer) (result *v1alpha1.Deployer, err error) {
	result = &v1alpha1.Deployer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployers").
		Name(deployer.Name).
		Body(deployer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deployers) UpdateStatus(deployer *v1alpha1.Deployer) (result *v1alpha1.Deployer, err error) {
	result = &v1alpha1.Deployer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployers").
		Name(deployer.Name).
		SubResource("status").
		Body(deployer).
		Do().
		Into(result)
	return
}

// Delete takes name of the deployer and deletes it. Returns an error if one occurs.
func (c *deployers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deployer.
func (c *deployers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Deployer, err error) {
	result = &v1alpha1.Deployer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
