package server

import (
	"github.com/alganbr/kedai-itemsvc/configs"
	"github.com/alganbr/kedai-itemsvc/internal/controllers"
	"github.com/alganbr/kedai-itemsvc/internal/databases"
	"github.com/alganbr/kedai-itemsvc/internal/managers"
	"github.com/alganbr/kedai-itemsvc/internal/repos"
	"github.com/alganbr/kedai-itemsvc/internal/routes"
	"go.uber.org/fx"
)

var controller = fx.Options(
	fx.Provide(controllers.NewHomeController),
	fx.Provide(controllers.NewItemController),
)

var manager = fx.Options(
	fx.Provide(managers.NewItemManager),
)

var repo = fx.Options(
	fx.Provide(repos.NewItemRepository),
)

var database = fx.Options(
	fx.Provide(databases.NewDB),
)

var router = fx.Options(
	fx.Provide(routes.NewRouter),
	fx.Provide(routes.NewRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewHomeRoutes),
	fx.Provide(routes.NewItemRoutes),
)

var server = fx.Options(
	fx.Provide(configs.NewConfig),
)

var Module = fx.Options(
	server,
	database,
	router,
	controller,
	manager,
	repo,
	fx.Invoke(StartApplication),
)
