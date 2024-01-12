// Copyright 2023 The xxx Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configserver

import (
	configv1alpha1 "github.com/iptecharch/config-server/apis/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewConfigSetTableConvertor(gr schema.GroupResource) tableConvertor {
	return tableConvertor{
		resource: gr,
		cells: func(obj runtime.Object) []interface{} {
			config, ok := obj.(*configv1alpha1.ConfigSet)
			if !ok {
				return nil
			}
			return []interface{}{
				config.Name,
				config.GetCondition(configv1alpha1.ConditionTypeReady).Status,
				len(config.Status.Targets),
				//config.ObjectMeta.CreationTimestamp,
			}
		},
		columns: []metav1.TableColumnDefinition{
			{Name: "Name", Type: "string"},
			{Name: "Ready", Type: "string"},
			{Name: "Targets", Type: "integer"},
			//{Name: "Created at", Type: "string"},
		},
	}
}
