package controller

import (
	"encoding/json"
	"fmt"
	"log"
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
	var d FormData
	err := decoder.Decode(&d)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
		return
	}

	defer r.Body.Close()

	if d.Email == "" || d.Message == "" {
		http.Error(w, "", 422)
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
		fmt.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
