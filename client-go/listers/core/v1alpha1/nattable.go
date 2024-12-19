// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// NATTableLister helps list NATTables.
// All objects returned here must be treated as read-only.
type NATTableLister interface {
	// List lists all NATTables in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NATTable, err error)
	// NATTables returns an object that can list and get NATTables.
	NATTables(namespace string) NATTableNamespaceLister
	NATTableListerExpansion
}

// nATTableLister implements the NATTableLister interface.
type nATTableLister struct {
	listers.ResourceIndexer[*v1alpha1.NATTable]
}

// NewNATTableLister returns a new NATTableLister.
func NewNATTableLister(indexer cache.Indexer) NATTableLister {
	return &nATTableLister{listers.New[*v1alpha1.NATTable](indexer, v1alpha1.Resource("nattable"))}
}

// NATTables returns an object that can list and get NATTables.
func (s *nATTableLister) NATTables(namespace string) NATTableNamespaceLister {
	return nATTableNamespaceLister{listers.NewNamespaced[*v1alpha1.NATTable](s.ResourceIndexer, namespace)}
}

// NATTableNamespaceLister helps list and get NATTables.
// All objects returned here must be treated as read-only.
type NATTableNamespaceLister interface {
	// List lists all NATTables in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NATTable, err error)
	// Get retrieves the NATTable from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NATTable, error)
	NATTableNamespaceListerExpansion
}

// nATTableNamespaceLister implements the NATTableNamespaceLister
// interface.
type nATTableNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.NATTable]
}
