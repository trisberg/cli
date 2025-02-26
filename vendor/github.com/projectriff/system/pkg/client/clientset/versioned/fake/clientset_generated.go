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
	clientset "github.com/projectriff/system/pkg/client/clientset/versioned"
	buildv1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/build/v1alpha1"
	fakebuildv1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/build/v1alpha1/fake"
	corev1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/core/v1alpha1"
	fakecorev1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/core/v1alpha1/fake"
	knativev1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/knative/v1alpha1"
	fakeknativev1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/knative/v1alpha1/fake"
	streamingv1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/streaming/v1alpha1"
	fakestreamingv1alpha1 "github.com/projectriff/system/pkg/client/clientset/versioned/typed/streaming/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// BuildV1alpha1 retrieves the BuildV1alpha1Client
func (c *Clientset) BuildV1alpha1() buildv1alpha1.BuildV1alpha1Interface {
	return &fakebuildv1alpha1.FakeBuildV1alpha1{Fake: &c.Fake}
}

// Build retrieves the BuildV1alpha1Client
func (c *Clientset) Build() buildv1alpha1.BuildV1alpha1Interface {
	return &fakebuildv1alpha1.FakeBuildV1alpha1{Fake: &c.Fake}
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return &fakecorev1alpha1.FakeCoreV1alpha1{Fake: &c.Fake}
}

// Core retrieves the CoreV1alpha1Client
func (c *Clientset) Core() corev1alpha1.CoreV1alpha1Interface {
	return &fakecorev1alpha1.FakeCoreV1alpha1{Fake: &c.Fake}
}

// KnativeV1alpha1 retrieves the KnativeV1alpha1Client
func (c *Clientset) KnativeV1alpha1() knativev1alpha1.KnativeV1alpha1Interface {
	return &fakeknativev1alpha1.FakeKnativeV1alpha1{Fake: &c.Fake}
}

// Knative retrieves the KnativeV1alpha1Client
func (c *Clientset) Knative() knativev1alpha1.KnativeV1alpha1Interface {
	return &fakeknativev1alpha1.FakeKnativeV1alpha1{Fake: &c.Fake}
}

// StreamingV1alpha1 retrieves the StreamingV1alpha1Client
func (c *Clientset) StreamingV1alpha1() streamingv1alpha1.StreamingV1alpha1Interface {
	return &fakestreamingv1alpha1.FakeStreamingV1alpha1{Fake: &c.Fake}
}

// Streaming retrieves the StreamingV1alpha1Client
func (c *Clientset) Streaming() streamingv1alpha1.StreamingV1alpha1Interface {
	return &fakestreamingv1alpha1.FakeStreamingV1alpha1{Fake: &c.Fake}
}
