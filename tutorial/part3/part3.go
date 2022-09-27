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
	"context"
	"encoding/json"
	"fmt"

	"github.com/atomist-skills/go-skill"
)

func handleBugWebhook(ctx context.Context, req skill.RequestContext) skill.Status {

	fmt.Println("Handling a webhook")
	bugJson := req.Event.Context.Webhook.Request.Body
	var mistyBug MistyBug
	json.Unmarshal([]byte(bugJson), &mistyBug)

	fmt.Printf("Received a MistyBug with ID [%s] in state [%s] and Summary [%s]\n", mistyBug.Id, mistyBug.State, mistyBug.Summary)

	transactMistyBug(ctx, req, mistyBug)

	return skill.Status{
		State:  skill.Completed,
		Reason: "Processed MistyBug webhook",
	}
}

func transactMistyBug(ctx context.Context, req skill.RequestContext, mistyBug MistyBug) error {

	transaction := skill.NewTransaction()

	transaction.AddEntities(BugEntity{
		Entity:  transaction.MakeEntity("bug"),
		ID:      mistyBug.Id,
		System:  "mistybugz",
		Summary: mistyBug.Summary,
		State:   mistyBug.State,
	})

	err := req.Transact(transaction.Entities())

	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Transacted MistyBug with id %s", mistyBug.Id)
	fmt.Println(msg)
	req.Log.Infof(msg)
	return err
}