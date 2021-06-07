/*
Copyright 2019 The Tekton Authors

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

package pipeline

import (
	"fmt"
	"sort"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/pod"
)

// Images holds the images reference for a number of container images used
// across tektoncd pipelines.
type Images struct {
	// EntrypointImage is container image containing our entrypoint binary.
	EntrypointImage        string
	EntrypointImageWindows string
	// NopImage is the container image used to kill sidecars.
	NopImage        string
	NopImageWindows string
	// GitImage is the container image with Git that we use to implement the Git source step.
	GitImage string
	// KubeconfigWriterImage is the container image containing our kubeconfig writer binary.
	KubeconfigWriterImage string
	// ShellImage is the container image containing bash shell.
	ShellImage string
	// GsutilImage is the container image containing gsutil.
	GsutilImage string
	// PRImage is the container image that we use to implement the PR source step.
	PRImage string
	// ImageDigestExporterImage is the container image containing our image digest exporter binary.
	ImageDigestExporterImage string

	// NOTE: Make sure to add any new images to Validate below!
}

// Validate returns an error if any image is not set.
func (i Images) Validate() error {
	var unset []string
	for _, f := range []struct {
		v, name string
	}{
		{i.EntrypointImage, "entrypoint"},
		{i.EntrypointImageWindows, "entrypoint-windows"},
		{i.NopImage, "nop"},
		{i.NopImageWindows, "nop-windows"},
		{i.GitImage, "git"},
		{i.KubeconfigWriterImage, "kubeconfig-writer"},
		{i.ShellImage, "shell"},
		{i.GsutilImage, "gsutil"},
		{i.PRImage, "pr"},
		{i.ImageDigestExporterImage, "imagedigest-exporter"},
	} {
		if f.v == "" {
			unset = append(unset, f.name)
		}
	}
	if len(unset) > 0 {
		sort.Strings(unset)
		return fmt.Errorf("found unset image flags: %s", unset)
	}
	return nil
}

// GetEntrypointImage returns the entrypoint image reference for the
// platform defined by a PodTemplate
func (i Images) GetEntrypointImage(podTemplate *pod.Template) string {
	if os, ok := podTemplate.NodeSelector["kubernetes.io/os"]; ok {
		if os == "windows" {
			return i.EntrypointImageWindows
		}
	}

	return i.EntrypointImage
}

// GetNopImage returns the Nop image reference for the platform
// defined by a PodTemplate
func (i Images) GetNopImage(podTemplate *pod.Template) string {
	if os, ok := podTemplate.NodeSelector["kubernetes.io/os"]; ok {
		if os == "windows" {
			return i.NopImageWindows
		}
	}

	return i.NopImage
}
