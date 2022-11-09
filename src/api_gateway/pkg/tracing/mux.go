// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracing

import (
	"fmt"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/api_gateway/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

// NewServeMux creates a new TracedServeMux.
func NewServeMux(tracer opentracing.Tracer, conf *configs.Configuration) *TracedServeMux {
	// First, we construct the mux and server. We don't want to start the server
	// until all handlers are registered.
	root := mux.NewRouter()
	root = root.PathPrefix("/api").Subrouter()
	root.Use(middleware.PanicRecovery)
	root.Use(middleware.Logging)
	casbinJWTRoleAuthorizer, err := middleware.NewCasbinJWTRoleAuthorizer(conf)
	if err != nil {
		fmt.Println("CasbinConfigPath:", conf.CasbinConfigPath)
		fmt.Println("MiddlewareRolesPath:", conf.MiddlewareRolesPath)
		panic(err)
	}
	root.Use(casbinJWTRoleAuthorizer.Middleware)

	return &TracedServeMux{
		mux:    root,
		tracer: tracer,
	}
}

// TracedServeMux is a wrapper around http.ServeMux that instruments handlers for tracing.
type TracedServeMux struct {
	mux    *mux.Router
	tracer opentracing.Tracer
}

// Handle implements http.ServeMux#Handle
func (tm *TracedServeMux) Handle(pattern string, handler http.Handler) {
	middleware := nethttp.Middleware(
		tm.tracer,
		handler,
		nethttp.OperationNameFunc(func(r *http.Request) string {
			return "HTTP " + r.Method + " " + pattern
		}))
	tm.mux.Handle(pattern, middleware)
}

// ServeHTTP implements http.ServeMux#ServeHTTP
func (tm *TracedServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm.mux.ServeHTTP(w, r)
}
