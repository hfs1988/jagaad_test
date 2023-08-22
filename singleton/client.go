package singleton

import "github.com/hfs1988/jagaad_test/handler"

type Handlers interface {
	GetUserHandler() handler.UserHandler
}
