package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type FormData struct {
	Email   string
	Message string
}

// FormPOST
func FormPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"), os.Getenv("MAILGUN_PUBLIC_API_KEY"))
	message := mailgun.NewMessage(
		d.Email,
		"Hello from brynn.se!",
		d.Message,
		os.Getenv("MAILGUN_EMAIL"))

	resp, id, err := mg.Send(message)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
