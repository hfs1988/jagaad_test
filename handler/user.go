package handler

import (
	"fmt"

	"github.com/hfs1988/jagaad_test/adapter"
	"github.com/hfs1988/jagaad_test/service"
)

type userHandler struct {
	userService service.UserService
	logger      adapter.LogAdapter
}

func GetUserHandler(userService service.UserService, logger adapter.LogAdapter) *userHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *userHandler) SaveUsers() {
	users, err := h.userService.FetchUsers()
	if err != nil {
		h.logger.Error(err)
		return
	}

	err = h.userService.WriteUsers(users)
	if err != nil {
		h.logger.Error(err)
		return
	}

	h.logger.Info("Success save users to CSV file!")
}

func (h *userHandler) GetUsers(tags []string) {
	users, err := h.userService.GetUsers(tags)
	if err != nil {
		h.logger.Error(err)
		return
	}
	for k, v := range users {
		fmt.Printf("%d. ID: %s, Salary: %s\n", k+1, v.ID, v.Balance)
	}
}
