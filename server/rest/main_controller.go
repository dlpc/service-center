//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package rest

import (
	"encoding/json"
	"github.com/ServiceComb/service-center/util"
	"github.com/ServiceComb/service-center/util/rest"
	"github.com/ServiceComb/service-center/version"
	"github.com/astaxie/beego"
	"net/http"
)

type Version struct {
	Version    string `json:"version"`
	ApiVersion string `json:"apiVersion"`
	BuildTag   string `json:"buildTag"`
}

type Result struct {
	Info   string `json:"info" description:"return info"`
	Status int    `json:"status" description:"http return code"`
}

type MainService struct {
	//
}

func (this *MainService) URLPatterns() []rest.Route {
	return []rest.Route{
		{rest.HTTP_METHOD_GET, "/version", this.GetVersion},
		{rest.HTTP_METHOD_GET, "/health", this.CluterHealth},
	}
}

func (this *MainService) CluterHealth(w http.ResponseWriter, r *http.Request) {
	resp, err := InstanceAPI.CluterHealth(r.Context())
	if err != nil {
		util.LOGGER.Error("health check failed", err)
		WriteText(http.StatusInternalServerError, "health check failed", w)
		return
	}

	respInternal := resp.Response
	resp.Response = nil
	WriteJsonResponse(respInternal, resp, err, w)
}

func (this *MainService) GetVersion(w http.ResponseWriter, r *http.Request) {
	buildTag := beego.AppConfig.String("build_tag")
	version := Version{
		version.Version,
		version.ApiVersion,
		buildTag,
	}
	versionJSON, _ := json.Marshal(version)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(versionJSON)
}
