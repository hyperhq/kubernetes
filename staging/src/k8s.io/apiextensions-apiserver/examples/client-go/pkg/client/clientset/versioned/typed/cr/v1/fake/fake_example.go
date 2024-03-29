/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	cr_v1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "github.com/hyperhq/client-go/testing"
)

// FakeExamples implements ExampleInterface
type FakeExamples struct {
	Fake *FakeCrV1
	ns   string
}

var examplesResource = schema.GroupVersionResource{Group: "cr.client-go.k8s.io", Version: "v1", Resource: "examples"}

var examplesKind = schema.GroupVersionKind{Group: "cr.client-go.k8s.io", Version: "v1", Kind: "Example"}

// Get takes name of the example, and returns the corresponding example object, and an error if there is any.
func (c *FakeExamples) Get(name string, options v1.GetOptions) (result *cr_v1.Example, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(examplesResource, c.ns, name), &cr_v1.Example{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.Example), err
}

// List takes label and field selectors, and returns the list of Examples that match those selectors.
func (c *FakeExamples) List(opts v1.ListOptions) (result *cr_v1.ExampleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(examplesResource, examplesKind, c.ns, opts), &cr_v1.ExampleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cr_v1.ExampleList{}
	for _, item := range obj.(*cr_v1.ExampleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested examples.
func (c *FakeExamples) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(examplesResource, c.ns, opts))

}

// Create takes the representation of a example and creates it.  Returns the server's representation of the example, and an error, if there is any.
func (c *FakeExamples) Create(example *cr_v1.Example) (result *cr_v1.Example, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(examplesResource, c.ns, example), &cr_v1.Example{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.Example), err
}

// Update takes the representation of a example and updates it. Returns the server's representation of the example, and an error, if there is any.
func (c *FakeExamples) Update(example *cr_v1.Example) (result *cr_v1.Example, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(examplesResource, c.ns, example), &cr_v1.Example{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.Example), err
}

// Delete takes name of the example and deletes it. Returns an error if one occurs.
func (c *FakeExamples) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(examplesResource, c.ns, name), &cr_v1.Example{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExamples) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(examplesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cr_v1.ExampleList{})
	return err
}

// Patch applies the patch and returns the patched example.
func (c *FakeExamples) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cr_v1.Example, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(examplesResource, c.ns, name, data, subresources...), &cr_v1.Example{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.Example), err
}
