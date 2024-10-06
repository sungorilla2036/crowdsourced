//go:build wireinject
// +build wireinject

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// The build tag makes sure the stub is not built in the final build.

package answercmd

import (
	"github.com/google/wire"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/log"
	"github.com/sungorilla2036/crowdsourced/internal/base/conf"
	"github.com/sungorilla2036/crowdsourced/internal/base/cron"
	"github.com/sungorilla2036/crowdsourced/internal/base/data"
	"github.com/sungorilla2036/crowdsourced/internal/base/middleware"
	"github.com/sungorilla2036/crowdsourced/internal/base/server"
	"github.com/sungorilla2036/crowdsourced/internal/base/translator"
	"github.com/sungorilla2036/crowdsourced/internal/controller"
	"github.com/sungorilla2036/crowdsourced/internal/controller/template_render"
	"github.com/sungorilla2036/crowdsourced/internal/controller_admin"
	"github.com/sungorilla2036/crowdsourced/internal/repo"
	"github.com/sungorilla2036/crowdsourced/internal/router"
	"github.com/sungorilla2036/crowdsourced/internal/service"
	"github.com/sungorilla2036/crowdsourced/internal/service/service_config"
)

// initApplication init application.
func initApplication(
	debug bool,
	serverConf *conf.Server,
	dbConf *data.Database,
	cacheConf *data.CacheConf,
	i18nConf *translator.I18n,
	swaggerConf *router.SwaggerConfig,
	serviceConf *service_config.ServiceConfig,
	uiConf *server.UI,
	logConf log.Logger) (*pacman.Application, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		controller.ProviderSetController,
		controller_admin.ProviderSetController,
		templaterender.ProviderSetTemplateRenderController,
		service.ProviderSetService,
		cron.ProviderSetService,
		repo.ProviderSetRepo,
		translator.ProviderSet,
		middleware.ProviderSetMiddleware,
		newApplication,
	))
}
