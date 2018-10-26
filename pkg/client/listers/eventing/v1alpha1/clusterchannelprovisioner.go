/*
Copyright 2018 The Knative Authors

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
	v1alpha1 "github.com/knative/eventing/pkg/apis/eventing/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterChannelProvisionerLister helps list ClusterChannelProvisioners.
type ClusterChannelProvisionerLister interface {
	// List lists all ClusterChannelProvisioners in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterChannelProvisioner, err error)
	// Get retrieves the ClusterChannelProvisioner from the index for a given name.
	Get(name string) (*v1alpha1.ClusterChannelProvisioner, error)
	ClusterChannelProvisionerListerExpansion
}

// clusterChannelProvisionerLister implements the ClusterChannelProvisionerLister interface.
type clusterChannelProvisionerLister struct {
	indexer cache.Indexer
}

// NewClusterChannelProvisionerLister returns a new ClusterChannelProvisionerLister.
func NewClusterChannelProvisionerLister(indexer cache.Indexer) ClusterChannelProvisionerLister {
	return &clusterChannelProvisionerLister{indexer: indexer}
}

// List lists all ClusterChannelProvisioners in the indexer.
func (s *clusterChannelProvisionerLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterChannelProvisioner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterChannelProvisioner))
	})
	return ret, err
}

// Get retrieves the ClusterChannelProvisioner from the index for a given name.
func (s *clusterChannelProvisionerLister) Get(name string) (*v1alpha1.ClusterChannelProvisioner, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterchannelprovisioner"), name)
	}
	return obj.(*v1alpha1.ClusterChannelProvisioner), nil
}
