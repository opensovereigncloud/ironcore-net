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

// IPAddressInformer provides access to a shared informer and lister for
// IPAddresses.
type IPAddressInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IPAddressLister
}

type iPAddressInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIPAddressInformer constructs a new informer for IPAddress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIPAddressInformer(client ironcorenet.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIPAddressInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIPAddressInformer constructs a new informer for IPAddress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIPAddressInformer(client ironcorenet.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().IPAddresses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().IPAddresses().Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.IPAddress{},
		resyncPeriod,
		indexers,
	)
}

func (f *iPAddressInformer) defaultInformer(client ironcorenet.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIPAddressInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iPAddressInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.IPAddress{}, f.defaultInformer)
}

func (f *iPAddressInformer) Lister() v1alpha1.IPAddressLister {
	return v1alpha1.NewIPAddressLister(f.Informer().GetIndexer())
}
