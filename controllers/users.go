package controllers

import (
	"fmt"
	"net/http"
)

func (h *httpServer)CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)CreateUsersWithListInput(w http.ResponseWriter, r *http.Request){}


func (h *httpServer)LoginUser(w http.ResponseWriter, r *http.Request, params LoginUserParams){}

func (h *httpServer)LogoutUser(w http.ResponseWriter, r *http.Request){}


func (h *httpServer) DeleteUser(w http.ResponseWriter, r *http.Request, username string){}


func (h *httpServer)GetUserByName(w http.ResponseWriter, r *http.Request, username string){}


func (h *httpServer)UpdateUser(w http.ResponseWriter, r *http.Request, username string){}

