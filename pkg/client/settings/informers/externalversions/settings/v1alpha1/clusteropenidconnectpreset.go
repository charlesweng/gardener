/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1alpha1

import (
	"context"
	time "time"

	settingsv1alpha1 "github.com/gardener/gardener/pkg/apis/settings/v1alpha1"
	versioned "github.com/gardener/gardener/pkg/client/settings/clientset/versioned"
	internalinterfaces "github.com/gardener/gardener/pkg/client/settings/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/gardener/pkg/client/settings/listers/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterOpenIDConnectPresetInformer provides access to a shared informer and lister for
// ClusterOpenIDConnectPresets.
type ClusterOpenIDConnectPresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterOpenIDConnectPresetLister
}

type clusterOpenIDConnectPresetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterOpenIDConnectPresetInformer constructs a new informer for ClusterOpenIDConnectPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterOpenIDConnectPresetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterOpenIDConnectPresetInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterOpenIDConnectPresetInformer constructs a new informer for ClusterOpenIDConnectPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterOpenIDConnectPresetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().ClusterOpenIDConnectPresets().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().ClusterOpenIDConnectPresets().Watch(context.TODO(), options)
			},
		},
		&settingsv1alpha1.ClusterOpenIDConnectPreset{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterOpenIDConnectPresetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterOpenIDConnectPresetInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterOpenIDConnectPresetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&settingsv1alpha1.ClusterOpenIDConnectPreset{}, f.defaultInformer)
}

func (f *clusterOpenIDConnectPresetInformer) Lister() v1alpha1.ClusterOpenIDConnectPresetLister {
	return v1alpha1.NewClusterOpenIDConnectPresetLister(f.Informer().GetIndexer())
}
