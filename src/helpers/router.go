package helpers

import (
	webHookData "../routes/webHook"
)

func SetAllRoutes() map[string][]func() {
	routerSchema := make(map[string][]func())

	routerSchema["/api"] = append(routerSchema["/api"], webHookData.HandleOrderWebHook)

	return routerSchema

}
