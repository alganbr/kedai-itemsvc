package routes

type Routes struct {
	routes []Route
}

type Route interface {
	Setup()
}

func NewRoutes(
	swaggerRoutes SwaggerRoutes,
	homeRoutes HomeRoutes,
	itemRoutes ItemRoutes,
) Routes {
	return Routes{
		routes: []Route{
			&swaggerRoutes,
			&homeRoutes,
			&itemRoutes,
		},
	}
}

func (r *Routes) Setup() {
	for _, route := range r.routes {
		route.Setup()
	}
}
