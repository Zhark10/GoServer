package users

import (
	"github.com/julienschmidt/httprouter"
	"goServer/internal/handlers"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.getUsers)
	router.GET(userURL, h.getUserById)
	router.POST(usersURL, h.createUser)
	router.PUT(userURL, h.updateUser)
	router.PATCH(userURL, h.updatePartOfUser)
	router.DELETE(userURL, h.deleteUser)
}

func (h *handler) getUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("all users"))
}

func (h *handler) getUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("user by id"))
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("user created"))
}

func (h *handler) updateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("user updated"))
}

func (h *handler) updatePartOfUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("part of user updated"))
}

func (h *handler) deleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("user deleted"))
}
