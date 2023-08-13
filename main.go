package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/bblj/totp/models"
	qrcode "github.com/skip2/go-qrcode"

	"encoding/base32"

	"github.com/gorilla/mux"
)

var config *models.Config

func main() {

	config = readConfig()
	r := mux.NewRouter()
	r.PathPrefix("/statics/").Handler(http.FileServer(http.Dir(".")))
	r.HandleFunc("/generate", generate).Methods("POST")
	r.HandleFunc("/verify", verify).Methods("GET")
	r.HandleFunc("/", index)
	fmt.Printf("http://localhost:%s", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}

func readConfig() *models.Config {
	var config models.Config = models.Config{
		Port:   "8080",
		Secret: "HELLO",
	}
	jsonFile, err := os.Open("config.json")
	fmt.Println(err)
	if err == nil {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		err := json.Unmarshal(byteValue, &config)
		fmt.Println(err)
		defer jsonFile.Close()
	}
	return &config
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/index.html"))
	tmpl.Execute(w, nil)
}

func generate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5242880)
	issuer := r.Form.Get("Issuer")
	email := r.Form.Get("Email")
	data := []byte(config.Secret + ":" + email)
	secret := base32.StdEncoding.EncodeToString(data)
	totpUrl := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, email, secret, issuer)
	png, err := qrcode.Encode(totpUrl, qrcode.Medium, 256)
	dataURI := ""
	if err == nil {
		dataURI = "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(png))
	} else {
		fmt.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("./views/generate.html"))
	tmpl.Execute(w, struct {
		Issuer  string
		Email   string
		Secret  string
		DataURI string
	}{
		Issuer:  issuer,
		Email:   email,
		Secret:  secret,
		DataURI: dataURI,
	})
}

func verify(w http.ResponseWriter, r *http.Request) {

}
