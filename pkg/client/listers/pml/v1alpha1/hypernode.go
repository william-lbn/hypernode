/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/william-lbn/hypernode/pkg/apis/pml/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HyperNodeLister helps list HyperNodes.
// All objects returned here must be treated as read-only.
type HyperNodeLister interface {
	// List lists all HyperNodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HyperNode, err error)
	// HyperNodes returns an object that can list and get HyperNodes.
	HyperNodes(namespace string) HyperNodeNamespaceLister
	HyperNodeListerExpansion
}

// hyperNodeLister implements the HyperNodeLister interface.
type hyperNodeLister struct {
	indexer cache.Indexer
}

// NewHyperNodeLister returns a new HyperNodeLister.
func NewHyperNodeLister(indexer cache.Indexer) HyperNodeLister {
	return &hyperNodeLister{indexer: indexer}
}

// List lists all HyperNodes in the indexer.
func (s *hyperNodeLister) List(selector labels.Selector) (ret []*v1alpha1.HyperNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HyperNode))
	})
	return ret, err
}

// HyperNodes returns an object that can list and get HyperNodes.
func (s *hyperNodeLister) HyperNodes(namespace string) HyperNodeNamespaceLister {
	return hyperNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HyperNodeNamespaceLister helps list and get HyperNodes.
// All objects returned here must be treated as read-only.
type HyperNodeNamespaceLister interface {
	// List lists all HyperNodes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HyperNode, err error)
	// Get retrieves the HyperNode from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HyperNode, error)
	HyperNodeNamespaceListerExpansion
}

// hyperNodeNamespaceLister implements the HyperNodeNamespaceLister
// interface.
type hyperNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HyperNodes in the indexer for a given namespace.
func (s hyperNodeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HyperNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HyperNode))
	})
	return ret, err
}

// Get retrieves the HyperNode from the indexer for a given namespace and name.
func (s hyperNodeNamespaceLister) Get(name string) (*v1alpha1.HyperNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hypernode"), name)
	}
	return obj.(*v1alpha1.HyperNode), nil
}
