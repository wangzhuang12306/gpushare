/*
Copyright 2021 wz123456.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	gpusharev1 "github.com/wangzhuang12306/GpuShare/api/gpushare/v1"
	versioned "github.com/wangzhuang12306/GpuShare/pkg/client/clientset/versioned"
	internalinterfaces "github.com/wangzhuang12306/GpuShare/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/wangzhuang12306/GpuShare/pkg/client/listers/gpushare/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualPodInformer provides access to a shared informer and lister for
// VirtualPods.
type VirtualPodInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VirtualPodLister
}

type virtualPodInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualPodInformer constructs a new informer for VirtualPod type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualPodInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualPodInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualPodInformer constructs a new informer for VirtualPod type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualPodInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GpushareV1().VirtualPods(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GpushareV1().VirtualPods(namespace).Watch(context.TODO(), options)
			},
		},
		&gpusharev1.VirtualPod{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualPodInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualPodInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualPodInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gpusharev1.VirtualPod{}, f.defaultInformer)
}

func (f *virtualPodInformer) Lister() v1.VirtualPodLister {
	return v1.NewVirtualPodLister(f.Informer().GetIndexer())
}
