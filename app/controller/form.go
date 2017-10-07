package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larsha/brynn.se-go/app/shared/config"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type FormData struct {
	Email   string
	Message string
}

// FormPOST
func FormPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config := config.Get()
	decoder := json.NewDecoder(r.Body)
	d := FormData{}
	err := decoder.Decode(&d)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if d.Email == "" || d.Message == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	mg := mailgun.NewMailgun(
		config.Mailgun.Domain,
		config.Mailgun.ApiKey,
		config.Mailgun.PublicApiKey)

	message := mailgun.NewMessage(
		d.Email,
		config.Mailgun.Subject,
		d.Message,
		config.Mailgun.Email)

	resp, id, err := mg.Send(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
