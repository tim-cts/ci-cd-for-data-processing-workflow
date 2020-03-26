// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and

package main

import (
	"flag"
	"path/filepath"
	"source.cloud.google.com/datapipelines-ci/composer/cloudbuild/go/dagsdeployer/internal/composerdeployer"
)

func main() {

	var repoRoot, projectID, composerRegion, composerEnvName, dagBucketPrefix string
	var replace bool

	flag.StringVar(&repoRoot, "repo", "", "path to root of repo")
	flag.StringVar(&projectID, "project", "", "gcp project id")
	flag.StringVar(&composerRegion, "region", "", "project")
	flag.StringVar(&composerEnvName, "composerEnv", "", "Composer environment name")
	flag.StringVar(&dagBucketPrefix, "dagBucketPrefix", "", "Composer DAGs bucket prefix")
	flag.BoolVar(&replace, "replace", false, "Boolean flag to indicatae if source dag mismatches the object of same name in GCS delte the old version and deploy over it")

	flag.Parse()

	DagListFile := filepath.Join(repoRoot, "composer", "config", "running_dags.txt")

	c := composerdeployer.ComposerEnv{
		Name:                composerEnvName,
		Location:            composerRegion,
		DagBucketPrefix:     dagBucketPrefix,
		LocalComposerPrefix: filepath.Join(repoRoot, "composer")}

	dagsToStop, dagsToStart := c.GetStopAndStartDags(DagListFile, replace)
	c.StopDags(dagsToStop)
	c.StartDags(repoRoot, dagsToStart)
}