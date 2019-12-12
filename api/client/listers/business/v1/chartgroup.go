/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tke/api/business/v1"
)

// ChartGroupLister helps list ChartGroups.
type ChartGroupLister interface {
	// List lists all ChartGroups in the indexer.
	List(selector labels.Selector) (ret []*v1.ChartGroup, err error)
	// ChartGroups returns an object that can list and get ChartGroups.
	ChartGroups(namespace string) ChartGroupNamespaceLister
	ChartGroupListerExpansion
}

// chartGroupLister implements the ChartGroupLister interface.
type chartGroupLister struct {
	indexer cache.Indexer
}

// NewChartGroupLister returns a new ChartGroupLister.
func NewChartGroupLister(indexer cache.Indexer) ChartGroupLister {
	return &chartGroupLister{indexer: indexer}
}

// List lists all ChartGroups in the indexer.
func (s *chartGroupLister) List(selector labels.Selector) (ret []*v1.ChartGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ChartGroup))
	})
	return ret, err
}

// ChartGroups returns an object that can list and get ChartGroups.
func (s *chartGroupLister) ChartGroups(namespace string) ChartGroupNamespaceLister {
	return chartGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChartGroupNamespaceLister helps list and get ChartGroups.
type ChartGroupNamespaceLister interface {
	// List lists all ChartGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ChartGroup, err error)
	// Get retrieves the ChartGroup from the indexer for a given namespace and name.
	Get(name string) (*v1.ChartGroup, error)
	ChartGroupNamespaceListerExpansion
}

// chartGroupNamespaceLister implements the ChartGroupNamespaceLister
// interface.
type chartGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ChartGroups in the indexer for a given namespace.
func (s chartGroupNamespaceLister) List(selector labels.Selector) (ret []*v1.ChartGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ChartGroup))
	})
	return ret, err
}

// Get retrieves the ChartGroup from the indexer for a given namespace and name.
func (s chartGroupNamespaceLister) Get(name string) (*v1.ChartGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("chartgroup"), name)
	}
	return obj.(*v1.ChartGroup), nil
}
