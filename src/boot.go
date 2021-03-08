package src

import (
	"ansme/src/config"
	"ansme/src/router"
	"ansme/src/utils/logger"
	"fmt"
	"net/http"
)

type App struct{}

func (app *App) Start() {

	app.loadConfig()
	logger.Info("server config loaded")

	app.loadRouterGroup()
	logger.Info("server router loaded")

	app.startServer()
}

// 加载配置信息
func (app *App) loadConfig() {
	error := config.Load()

	if error != nil {
		logger.Critical(fmt.Sprintf("Error loadConfig: %s", error.Error()))
		return
	}
}

// 启动 web 监听服务
func (app *App) startServer() {

	serverAddress := getServerAddr()

	logger.Info(fmt.Sprintf("server started at: %s", serverAddress))
	error := http.ListenAndServe(serverAddress, nil)

	if error != nil {
		logger.Critical(fmt.Sprintf("Error startServer: %s", error.Error()))
		return
	}

}

func getServerAddr() string {
	host := config.Get("server_host")
	if host == "" {
		host = "127.0.0.1"
	}

	port := config.Get("server_port")
	if port == "" {
		port = "2000"
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	return addr
}

// 加载路由
func (app *App) loadRouterGroup() {

	routerGroup := router.RouterGroup

	for _, router := range routerGroup {
		app.Handler(router.Method, router.Path, router.Handler)

	}
}

func (app *App) Handler(method, path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(path, handler)
}
