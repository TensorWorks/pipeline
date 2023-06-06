/*
Copyright 2022 The Tekton Authors

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

package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	for i, d := range os.Args {
		if i == 0 {
			// os.Args[0] is the path to this executable, so we should skip it
			continue
		}

		ws := "/workspace/"
		p := filepath.Clean(d)

		if runtime.GOOS == "windows" {
			ws = "C:\\workspace\\"

			if strings.HasPrefix(p, "\\") && !strings.HasPrefix(p, "\\\\") {
				p = "C:" + p
			}
		}

		if !filepath.IsAbs(p) || strings.HasPrefix(p, ws) {
			if err := os.MkdirAll(p, 0755); err != nil {
				log.Fatalf("Failed to mkdir %q: %v", p, err)
			}
		}
	}
}
