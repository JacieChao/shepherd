/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1alpha1.AddToScheme)
}

type Interface interface {
	Bundle() BundleController
	Cluster() ClusterController
	ClusterGroup() ClusterGroupController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) Bundle() BundleController {
	return NewBundleController(schema.GroupVersionKind{Group: "fleet.cattle.io", Version: "v1alpha1", Kind: "Bundle"}, "bundles", true, c.controllerFactory)
}
func (c *version) Cluster() ClusterController {
	return NewClusterController(schema.GroupVersionKind{Group: "fleet.cattle.io", Version: "v1alpha1", Kind: "Cluster"}, "clusters", true, c.controllerFactory)
}
func (c *version) ClusterGroup() ClusterGroupController {
	return NewClusterGroupController(schema.GroupVersionKind{Group: "fleet.cattle.io", Version: "v1alpha1", Kind: "ClusterGroup"}, "clustergroups", true, c.controllerFactory)
}
