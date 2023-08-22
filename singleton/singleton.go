package singleton

import (
	"github.com/hfs1988/jagaad_test/adapter"
	"github.com/hfs1988/jagaad_test/config"
	"github.com/hfs1988/jagaad_test/handler"
	"github.com/hfs1988/jagaad_test/service"
)

var instance *handlers

type handlers struct {
	userHandler handler.UserHandler
}

func GetHandlers() *handlers {
	if instance != nil {
		return instance
	}

	var (
		httpAdapter adapter.HTTPAdapter = adapter.GetHTTPAdapter()
		logAdapter  adapter.LogAdapter  = adapter.GetLogAdapter()
		csvAdapter  adapter.CSVAdapter  = adapter.GetCSVAdapter()
		conf        config.Config       = config.GetConfig()
		userService service.UserService = service.GetUserService(httpAdapter, csvAdapter, conf)
		userHandler handler.UserHandler = handler.GetUserHandler(userService, logAdapter)
	)

	instance = &handlers{
		userHandler: userHandler,
	}

	return instance
}

func (h *handlers) GetUserHandler() handler.UserHandler {
	return h.userHandler
}
