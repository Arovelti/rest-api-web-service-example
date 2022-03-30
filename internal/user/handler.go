package user

import (
	"net/http"

	"rest-api-tutorial/cmd/pkg/logging"
	"rest-api-tutorial/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

//var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/user/:uuid" //Universal Unique Identifier
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/usersURL", h.GetList)
	router.GET("/userURL", h.GetUserByUUID)
	router.POST("/usersURL", h.CreateUser)
	router.PUT("/userURL", h.UpdateUser)
	router.PATCH("/userURL", h.PartiallyUpdateUser)
	router.DELETE("/userURL", h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("ZAGLUSKA"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("ZAGLUSKA"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("ZAGLUSKA"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("ZAGLUSKA"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("ZAGLUSKA"))
}
