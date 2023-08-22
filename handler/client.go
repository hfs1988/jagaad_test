package handler

type UserHandler interface {
	SaveUsers()
	GetUsers(tags []string)
}
