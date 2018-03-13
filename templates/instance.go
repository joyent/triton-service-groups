//  Copyright (c) 2018, Joyent, Inc. All rights reserved.
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package templates_v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyent/triton-service-groups/server/handlers"
	"github.com/rs/zerolog/log"
)

type InstanceTemplate struct {
	ID                 int64             `json:"id"`
	TemplateName       string            `json:"template_name"`
	AccountID          string            `json:"account_id"`
	Package            string            `json:"package"`
	ImageID            string            `json:"image_id"`
	InstanceNamePrefix string            `json:"instance_name_prefix"`
	FirewallEnabled    bool              `json:"firewall_enabled"`
	Networks           []string          `json:"networks"`
	UserData           string            `json:"userdata"`
	MetaData           map[string]string `json:"metadata"`
	Tags               map[string]string `json:"tags"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, got := handlers.GetAuthSession(ctx)
	if !got {
		log.Fatal().Err(handlers.ErrNoSession)
		http.Error(w, handlers.ErrNoSession.Error(), http.StatusUnauthorized)
	}

	vars := mux.Vars(r)
	identifier := vars["identifier"]

	var template *InstanceTemplate

	id, err := strconv.Atoi(identifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template, ok := FindTemplateByID(ctx, int64(id), session.AccountID)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := json.Marshal(template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

// func Create(session *session.TsgSession) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		var template *InstanceTemplate
// 		err = json.Unmarshal(body, &template)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)

// 		}

// 		SaveTemplate(session.DbPool, session.AccountId, template)

// 		w.Header().Set("Location", r.URL.Path+"/"+template.TemplateName)

// 		com, ok := FindTemplateByName(session.DbPool, template.TemplateName, session.AccountId)
// 		if !ok {
// 			http.NotFound(w, r)
// 			return
// 		}

// 		bytes, err := json.Marshal(com)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		writeJsonResponse(w, bytes)
// 	}
// }

// func Delete(session *session.TsgSession) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		identifier := vars["identifier"]

// 		var template *InstanceTemplate

// 		id, err := strconv.Atoi(identifier)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		template, ok := FindTemplateByID(session.DbPool, int64(id), session.AccountId)
// 		if !ok {
// 			http.NotFound(w, r)
// 			return
// 		}

// 		RemoveTemplate(session.DbPool, template.ID, session.AccountId)
// 		w.WriteHeader(http.StatusNoContent)
// 	}
// }

func List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := handlers.GetAuthSession(ctx)
	if !ok {
		log.Fatal().Err(handlers.ErrNoSession)
		http.Error(w, handlers.ErrNoSession.Error(), http.StatusUnauthorized)
	}

	rows, err := FindTemplates(ctx, session.AccountID)
	if err != nil {
		log.Fatal().Err(err)
		http.NotFound(w, r)
		return
	}

	bytes, err := json.Marshal(rows)
	if err != nil {
		log.Fatal().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if n, err := w.Write(bytes); err != nil {
		log.Printf("%v", err)
	} else if n != len(bytes) {
		log.Printf("short write: %d/%d", n, len(bytes))
	}
}
