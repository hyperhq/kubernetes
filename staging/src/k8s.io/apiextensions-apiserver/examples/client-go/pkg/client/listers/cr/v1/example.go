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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/hyperhq/client-go/tools/cache"
)

// ExampleLister helps list Examples.
type ExampleLister interface {
	// List lists all Examples in the indexer.
	List(selector labels.Selector) (ret []*v1.Example, err error)
	// Examples returns an object that can list and get Examples.
	Examples(namespace string) ExampleNamespaceLister
	ExampleListerExpansion
}

// exampleLister implements the ExampleLister interface.
type exampleLister struct {
	indexer cache.Indexer
}

// NewExampleLister returns a new ExampleLister.
func NewExampleLister(indexer cache.Indexer) ExampleLister {
	return &exampleLister{indexer: indexer}
}

// List lists all Examples in the indexer.
func (s *exampleLister) List(selector labels.Selector) (ret []*v1.Example, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Example))
	})
	return ret, err
}

// Examples returns an object that can list and get Examples.
func (s *exampleLister) Examples(namespace string) ExampleNamespaceLister {
	return exampleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExampleNamespaceLister helps list and get Examples.
type ExampleNamespaceLister interface {
	// List lists all Examples in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Example, err error)
	// Get retrieves the Example from the indexer for a given namespace and name.
	Get(name string) (*v1.Example, error)
	ExampleNamespaceListerExpansion
}

// exampleNamespaceLister implements the ExampleNamespaceLister
// interface.
type exampleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Examples in the indexer for a given namespace.
func (s exampleNamespaceLister) List(selector labels.Selector) (ret []*v1.Example, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Example))
	})
	return ret, err
}

// Get retrieves the Example from the indexer for a given namespace and name.
func (s exampleNamespaceLister) Get(name string) (*v1.Example, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("example"), name)
	}
	return obj.(*v1.Example), nil
}
