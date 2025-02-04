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

package handlers

import (
	"apiserver/pkg/database/models"
	utils "apiserver/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
)

// GetHealth checks the API server's connection to the database and reports this status.
// In the future, there may be more health checks to implement here. For now, the main error case
// inside the API server is the connection to the database.
func GetHealth(ctx context.Context, handlerConfig *ApiHandlerConfig, w http.ResponseWriter, r *http.Request) *utils.ServerError {
	if err := handlerConfig.Database.ProbeConnection(); err != nil {
		return utils.NewError(err, "failed to connect to database", 500, false)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")

	resp := models.ServerResp{
		Results:   nil,
		Status:    "OK",
		RequestId: "0",
		Error:     nil,
	}

	if err := enc.Encode(resp); err != nil {
		return utils.NewError(err, "error encoding data as json", 500, false)
	}
	return nil
}
