/*
Copyright 2021 The PlanetPulse Authors.

Planet Pulse is an API designed to serve climate data pulled from NOAA's
Global Monitoring Laboratory FTP server. This API is based on the
OpenAPI v3 specification.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

A copy of the GNU General Public License can be found here:
https://www.gnu.org/licenses/

API version: 0.1.0
Contact: planetpulse.api@gmail.com
*/

package server

import (
	"apiserver/pkg/server/handlers"
	"apiserver/pkg/server/handlers/co2"
	utils "apiserver/pkg/utils"
	"context"
	"strings"

	"github.com/gorilla/mux"
)

// NewRouter generates a new gorilla mux router to be used instead of the default golang http router.
func (apiserver *ApiServer) NewRouter(ctx context.Context, routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(utils.SetCSPHeaders)
	router.Use(utils.Gzip)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handlers.NewHandler(ctx, route.Handler, route.Name))
	}
	return router
}

// CreateRoutes returns a Routes list representing all routes on the server.
func (apiserver *ApiServer) CreateRoutes() Routes {
	return Routes{
		Route{
			"favicon",
			strings.ToUpper("Get"),
			"/favicon.ico",
			handlers.ApiHandler{
				Handler: handlers.GetFavicon,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},

		Route{
			"heatlh",
			strings.ToUpper("Get"),
			"/v1/health",
			handlers.ApiHandler{
				Handler: handlers.GetHealth,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},

		Route{
			"index",
			strings.ToUpper("Get"),
			"/",
			handlers.ApiHandler{
				Handler: handlers.GetIndex,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},

		Route{
			"v1",
			strings.ToUpper("Get"),
			"/v1",
			handlers.ApiHandler{
				Handler: handlers.GetIndex,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},

		Route{
			"co2",
			strings.ToUpper("Get"),
			"/v1/co2",
			handlers.ApiHandler{
				Handler: co2.Get,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},

		Route{
			"co2Weekly",
			strings.ToUpper("Get"),
			"/v1/co2/weekly",
			handlers.ApiHandler{
				Handler: co2.Get,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},
		Route{
			"co2Weekly",
			strings.ToUpper("Get"),
			"/v1/co2/weekly/increase",
			handlers.ApiHandler{
				Handler: co2.Get,
				Config: &handlers.ApiHandlerConfig{
					Database: apiserver.Database,
				},
			},
		},
		/*
			Route{
				"co2WeeklyId",
				strings.ToUpper("Get"),
				"/v1/co2/weekly/{id}",
				apiHandler.GetIndexHandler,
			},

			Route{
				"co2WeeklyIncrease",
				strings.ToUpper("Get"),
				"/v1/co2/weekly/increase",
				apiHandler.GetIndexHandler,
			},

			Route{
				"co2WeeklyPpm",
				strings.ToUpper("Get"),
				"/v1/co2/weekly/{ppm}",
				apiHandler.GetIndexHandler,
			},

			Route{
				"heatlh",
				strings.ToUpper("Get"),
				"/v1/health",
				apiHandler.GetHealthHandler,
			},*/
	}
}
