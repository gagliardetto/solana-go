// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package static

//go:generate rice embed-go

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const IndexFilename = "graphiql.html"
const (
	MIME_TYPE_CSS  = "text/css"
	MIME_TYPE_HTML = "text/html"
	MIME_TYPE_JS   = "application/javascript"
	MIME_TYPE_JSON = "application/json"
)

func RegisterStaticRoutes(router *mux.Router) {
	zlog.Info("registering static route")
	box := rice.MustFindBox("build")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/graphiql/", 302)
	})

	router.HandleFunc("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/graphiql/", 302)
	})

	serveIndexHTML(router, box, "/graphiql/")
	serveFileAsset(router, box, "/graphiql/graphiql_dfuse_override.css", "graphiql_dfuse_override.css", MIME_TYPE_CSS)
	serveFileAsset(router, box, "/graphiql/helper.js", "helper.js", MIME_TYPE_JS)
	serveFileAsset(router, box, "/graphiql/favorites.json", "favorites.json", MIME_TYPE_JSON)

	// Redirects since it was supported at some point, redirects everyone to `GraphiQL` instead
	router.HandleFunc("/playground", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/graphiql/", 302)
	})

	router.HandleFunc("/playground/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/graphiql/", 302)
	})
}

func serveInMemoryAsset(router *mux.Router, path string, content string, contentType string) {
	router.HandleFunc(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write([]byte(content))
	}))
}

func serveIndexHTML(router *mux.Router, box *rice.Box, path string) {
	zlog.Info("setting up index http handler")
	router.HandleFunc(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zlog.Info("serving index",
			zap.String("path", path),
			zap.String("file_to_template", IndexFilename),
		)
		reader, err := templatedIndex(box)
		if err != nil {
			zlog.Error("unable to serve graphiql.html", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("unable to read asset"))
			return
		}

		w.Header().Set("Content-Type", MIME_TYPE_HTML)
		_, _ = io.Copy(w, reader)
	}))
}

func serveFileAsset(router *mux.Router, box *rice.Box, path string, asset string, contentType string, options ...interface{}) {
	router.HandleFunc(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reader, err := box.Open(asset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unable to read asset"))
			return
		}
		defer reader.Close()

		w.Header().Set("Content-Type", contentType)

		// We ignore the error since if we are unable to write to HTTP pipe, it's probably broken
		io.Copy(w, reader)
	}))
}

func templatedIndex(box *rice.Box) (*bytes.Reader, error) {
	zlog.Info("rendering templated index")
	indexContent, err := box.Bytes(IndexFilename)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New(IndexFilename).Funcs(template.FuncMap{
		"json": func(v interface{}) (template.JS, error) {
			cnt, err := json.Marshal(v)
			return template.JS(cnt), err
		},
	}).Delims("--==", "==--").Parse(string(indexContent))
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, nil); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}
