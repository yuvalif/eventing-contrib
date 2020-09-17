/*
Copyright 2020 The Knative Authors

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "knative.dev/eventing-contrib/ceph/pkg/apis/sources/v1alpha1"
)

// CephSourceLister helps list CephSources.
type CephSourceLister interface {
	// List lists all CephSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CephSource, err error)
	// CephSources returns an object that can list and get CephSources.
	CephSources(namespace string) CephSourceNamespaceLister
	CephSourceListerExpansion
}

// cephSourceLister implements the CephSourceLister interface.
type cephSourceLister struct {
	indexer cache.Indexer
}

// NewCephSourceLister returns a new CephSourceLister.
func NewCephSourceLister(indexer cache.Indexer) CephSourceLister {
	return &cephSourceLister{indexer: indexer}
}

// List lists all CephSources in the indexer.
func (s *cephSourceLister) List(selector labels.Selector) (ret []*v1alpha1.CephSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CephSource))
	})
	return ret, err
}

// CephSources returns an object that can list and get CephSources.
func (s *cephSourceLister) CephSources(namespace string) CephSourceNamespaceLister {
	return cephSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CephSourceNamespaceLister helps list and get CephSources.
type CephSourceNamespaceLister interface {
	// List lists all CephSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CephSource, err error)
	// Get retrieves the CephSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CephSource, error)
	CephSourceNamespaceListerExpansion
}

// cephSourceNamespaceLister implements the CephSourceNamespaceLister
// interface.
type cephSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CephSources in the indexer for a given namespace.
func (s cephSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CephSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CephSource))
	})
	return ret, err
}

// Get retrieves the CephSource from the indexer for a given namespace and name.
func (s cephSourceNamespaceLister) Get(name string) (*v1alpha1.CephSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cephsource"), name)
	}
	return obj.(*v1alpha1.CephSource), nil
}