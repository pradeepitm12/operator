/*
Copyright 2021 The Tekton Authors

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

package common

import (
	mf "github.com/manifestival/manifestival"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func RemoveMonitoringLabel() mf.Transformer {
	return func(u *unstructured.Unstructured) error {
		if u.GetKind() != "Namespace" {
			return nil
		}

		n := &corev1.Namespace{}
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, n)
		if err != nil {
			return err
		}

		labels := n.GetLabels()

		for key := range labels {
			if key == "openshift.io/cluster-monitoring" {
				delete(labels, key)
			}
		}

		unstrObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(n)
		if err != nil {
			return err
		}
		u.SetUnstructuredContent(unstrObj)

		return nil
	}
}
