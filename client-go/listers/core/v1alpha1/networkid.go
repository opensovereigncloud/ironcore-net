// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkIDLister helps list NetworkIDs.
// All objects returned here must be treated as read-only.
type NetworkIDLister interface {
	// List lists all NetworkIDs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkID, err error)
	// Get retrieves the NetworkID from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NetworkID, error)
	NetworkIDListerExpansion
}

// networkIDLister implements the NetworkIDLister interface.
type networkIDLister struct {
	indexer cache.Indexer
}

// NewNetworkIDLister returns a new NetworkIDLister.
func NewNetworkIDLister(indexer cache.Indexer) NetworkIDLister {
	return &networkIDLister{indexer: indexer}
}

// List lists all NetworkIDs in the indexer.
func (s *networkIDLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkID, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkID))
	})
	return ret, err
}

// Get retrieves the NetworkID from the index for a given name.
func (s *networkIDLister) Get(name string) (*v1alpha1.NetworkID, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkid"), name)
	}
	return obj.(*v1alpha1.NetworkID), nil
}
