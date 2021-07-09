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
	"path"
	"testing"

	mf "github.com/manifestival/manifestival"
	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestReplaceNamespaceInClusterInterceptor(t *testing.T) {
	testData := path.Join("testdata", "test-namespace-remove-label.yaml")
	manifest, err := mf.ManifestFrom(mf.Recursive(testData))
	assert.NilError(t, err)

	manifest, err = manifest.Transform(RemoveMonitoringLabel())
	assert.NilError(t, err)

	n := &corev1.Namespace{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(manifest.Resources()[0].Object, n)
	if err != nil {
		assert.NilError(t, err)
	}
	assert.Equal(t, 1, len(n.GetLabels()))
	// will remove monitoring label but keep the other one
	assert.Equal(t, n.GetLabels()["app.kubernetes.io/part-of"], "tekton-pipelines")
}
