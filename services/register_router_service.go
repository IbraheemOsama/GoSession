package services

import (
	"github.com/gorilla/mux"
)

type IRegisterRouterService interface {
	InitAndRegisterService(router *mux.Router)
}
