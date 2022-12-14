/*
 * Copyright © 2022 Atomist, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/atomist-skills/go-skill"
)

// OnCommit maps the incoming commit of the on_push and on_commit_signature to a Go struct
type OnCommit struct {
	Sha     string `edn:"git.commit/sha"`
	Message string `edn:"git.commit/message"`
	Repo    struct {
		Org struct {
			Url string `edn:"git.provider/url"`
		} `edn:"git.repo/org"`
		SourceId string `edn:"git.repo/source-id"`
		Name     string `edn:"git.repo/name"`
	} `edn:"git.commit/repo"`
}

// GitRepoEntity provides mappings for a :git/repo entity
type GitRepoEntity struct {
	skill.Entity
	SourceId string `edn:"git.repo/source-id,omitempty"`
	Url      string `edn:"git.provider/url,omitempty"`
}

// GitCommitEntity provides mappings for a :git/commit entity
type GitCommitEntity struct {
	skill.Entity
	Sha  string        `edn:"git.commit/sha,omitempty"`
	Repo GitRepoEntity `edn:"git.commit/repo,omitempty"`
	Url  string        `edn:"git.provider/url,omitempty"`
}

type BugEntity struct {
	skill.Entity
	ID      string `edn:"bug/id,omitempty"`
	System  string `edn:"bug/system,omitempty"`
	Summary string `edn:"bug/summary,omitempty"`
	State   string `edn:"bug/state,omitempty"`
}

type LinkedBugEntity struct {
	skill.Entity
	Commit  GitCommitEntity `edn:"bug/commit,omitempty"`
	ID      string          `edn:"bug/id,omitempty"`
	System  string          `edn:"bug/system,omitempty"`
	Summary string          `edn:"bug/summary,omitempty"`
	State   string          `edn:"bug/state,omitempty"`
}

type MistyBug struct {
	Id      string
	Summary string
	State   string
}
