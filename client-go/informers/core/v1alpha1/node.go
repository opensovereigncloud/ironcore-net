// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	corev1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	internalinterfaces "github.com/ironcore-dev/ironcore-net/client-go/informers/internalinterfaces"
	ironcorenet "github.com/ironcore-dev/ironcore-net/client-go/ironcorenet"
	v1alpha1 "github.com/ironcore-dev/ironcore-net/client-go/listers/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeInformer provides access to a shared informer and lister for
// Nodes.
type NodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodeLister
}

type nodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeInformer constructs a new informer for Node type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeInformer(client ironcorenet.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeInformer constructs a new informer for Node type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeInformer(client ironcorenet.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().Nodes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().Nodes().Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.Node{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeInformer) defaultInformer(client ironcorenet.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.Node{}, f.defaultInformer)
}

func (f *nodeInformer) Lister() v1alpha1.NodeLister {
	return v1alpha1.NewNodeLister(f.Informer().GetIndexer())
}
