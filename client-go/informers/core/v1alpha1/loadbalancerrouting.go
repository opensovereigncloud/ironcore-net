/*
 * Copyright (c) 2022 by the OnMetal authors.
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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	corev1alpha1 "github.com/onmetal/onmetal-api-net/api/core/v1alpha1"
	internalinterfaces "github.com/onmetal/onmetal-api-net/client-go/informers/internalinterfaces"
	v1alpha1 "github.com/onmetal/onmetal-api-net/client-go/listers/core/v1alpha1"
	onmetalapinet "github.com/onmetal/onmetal-api-net/client-go/onmetalapinet"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LoadBalancerRoutingInformer provides access to a shared informer and lister for
// LoadBalancerRoutings.
type LoadBalancerRoutingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LoadBalancerRoutingLister
}

type loadBalancerRoutingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLoadBalancerRoutingInformer constructs a new informer for LoadBalancerRouting type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLoadBalancerRoutingInformer(client onmetalapinet.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLoadBalancerRoutingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLoadBalancerRoutingInformer constructs a new informer for LoadBalancerRouting type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLoadBalancerRoutingInformer(client onmetalapinet.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().LoadBalancerRoutings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().LoadBalancerRoutings(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.LoadBalancerRouting{},
		resyncPeriod,
		indexers,
	)
}

func (f *loadBalancerRoutingInformer) defaultInformer(client onmetalapinet.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLoadBalancerRoutingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *loadBalancerRoutingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.LoadBalancerRouting{}, f.defaultInformer)
}

func (f *loadBalancerRoutingInformer) Lister() v1alpha1.LoadBalancerRoutingLister {
	return v1alpha1.NewLoadBalancerRoutingLister(f.Informer().GetIndexer())
}
