package controllers

import (
	"fmt"
	"net/http"
)

func (h *httpServer)Addmeeting(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)FindmeetingsByStatus(w http.ResponseWriter, r *http.Request, params FindmeetingsByStatusParams){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)Updatemeeting(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)FindmeetingsByTags(w http.ResponseWriter, r *http.Request, params FindmeetingsByStatusParams){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)FindmeetingsByStatus(w http.ResponseWriter, r *http.Request, params FindmeetingsByStatusParams){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)Deletemeeting(w http.ResponseWriter, r *http.Request, meetingId int64, params DeletemeetingParams){
	fmt.Fprintln(w, "hello world")
}
	
func (h *httpServer)GetmeetingById(w http.ResponseWriter, r *http.Request, meetingId int64){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)UpdatemeetingWithForm(w http.ResponseWriter, r *http.Request, meetingId int64, params UpdatemeetingWithFormParams){
	fmt.Fprintln(w, "hello world")
}

func (h *httpServer)UploadFile(w http.ResponseWriter, r *http.Request, meetingId int64, params UploadFileParams){
    fmt.Fprintln(w, "hello world")
}
