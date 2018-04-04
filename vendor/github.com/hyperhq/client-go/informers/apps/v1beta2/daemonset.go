/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1beta2

import (
	apps_v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/hyperhq/client-go/informers/internalinterfaces"
	kubernetes "github.com/hyperhq/client-go/kubernetes"
	v1beta2 "github.com/hyperhq/client-go/listers/apps/v1beta2"
	cache "github.com/hyperhq/client-go/tools/cache"
	time "time"
)

// DaemonSetInformer provides access to a shared informer and lister for
// DaemonSets.
type DaemonSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.DaemonSetLister
}

type daemonSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDaemonSetInformer constructs a new informer for DaemonSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDaemonSetInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDaemonSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDaemonSetInformer constructs a new informer for DaemonSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDaemonSetInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta2().DaemonSets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta2().DaemonSets(namespace).Watch(options)
			},
		},
		&apps_v1beta2.DaemonSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *daemonSetInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDaemonSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *daemonSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps_v1beta2.DaemonSet{}, f.defaultInformer)
}

func (f *daemonSetInformer) Lister() v1beta2.DaemonSetLister {
	return v1beta2.NewDaemonSetLister(f.Informer().GetIndexer())
}
