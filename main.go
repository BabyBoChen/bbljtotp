package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/bblj/totp/models"
	"github.com/bblj/totp/secrets"
	"github.com/bblj/totp/utils"
	qrcode "github.com/skip2/go-qrcode"

	"github.com/gorilla/mux"
)

var config *models.Config

func main() {

	config = readConfig()
	r := mux.NewRouter()
	r.PathPrefix("/statics/").Handler(http.FileServer(http.Dir(".")))
	r.HandleFunc("/generate", generate).Methods("POST")
	r.HandleFunc("/generateApi", generateApi).Methods("GET")
	r.HandleFunc("/generateTOTP", generateTOTP).Methods("GET")
	r.HandleFunc("/verify", verify).Methods("GET")
	r.HandleFunc("/test", test).Methods("GET")
	r.HandleFunc("/apiUrl", apiUrl).Methods("GET")
	r.HandleFunc("/", index)
	fmt.Printf("http://localhost:%s", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}

func readConfig() *models.Config {
	var config models.Config = models.Config{
		Port:      "8080",
		Github:    "",
		SecretKey: secrets.SECRET_KEY,
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
	tmpl.Execute(w, struct {
		Github string
	}{
		Github: config.Github,
	})
}

func generate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5242880)
	secret := r.Form.Get("Secret")
	if len(secret) == 0 {
		t := utils.DateTimeString()
		data := []byte(config.SecretKey + ":" + t)
		secret = base32.StdEncoding.EncodeToString(data)
	}
	digit := r.Form.Get("Digit")
	period := r.Form.Get("Period")
	issuer := r.Form.Get("Issuer")
	username := r.Form.Get("Username")
	totpUrl := fmt.Sprintf("otpauth://totp/%s:%s?issuer=%s&secret=%s&algorithm=SHA1&digits=%s&period=%s", issuer, username, issuer, url.QueryEscape(secret), digit, period)
	png, err := qrcode.Encode(totpUrl, qrcode.Medium, 256)
	dataURI := ""
	if err == nil {
		dataURI = "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(png))
	} else {
		fmt.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("./views/generate.html"))
	tmpl.Execute(w, struct {
		Secret   string
		Digit    string
		Period   string
		Issuer   string
		Username string
		DataURI  string
		TotpUrl  string
	}{
		Secret:   secret,
		Digit:    digit,
		Period:   period,
		Issuer:   issuer,
		Username: username,
		DataURI:  dataURI,
		TotpUrl:  totpUrl,
	})
}

// return json => { secret: string, qrcode: string (e.g. "data:image/png;base64,... ")}
func generateApi(w http.ResponseWriter, r *http.Request) {
	queryVals := r.URL.Query()

	secret := queryVals.Get("secret")
	if len(secret) == 0 {
		t := utils.DateTimeString()
		data := []byte(config.SecretKey + ":" + t)
		secret = base32.StdEncoding.EncodeToString(data)
	}
	//digit := queryVals.Get("digit")
	digit := "6"
	//period := queryVals.Get("period")
	period := "30"
	issuer := queryVals.Get("issuer")
	username := queryVals.Get("username")
	totpUrl := fmt.Sprintf("otpauth://totp/%s:%s?issuer=%s&secret=%s&algorithm=SHA1&digits=%s&period=%s", issuer, username, issuer, url.QueryEscape(secret), digit, period)
	png, err := qrcode.Encode(totpUrl, qrcode.Medium, 256)
	dataURI := ""
	if err == nil {
		dataURI = "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(png))
	} else {
		fmt.Println(err)
	}
	key := struct {
		Secret string
		Qrcode string
	}{
		Secret: secret,
		Qrcode: dataURI,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(key)
}

// return json => true || false
func verify(w http.ResponseWriter, r *http.Request) {
	queryVals := r.URL.Query()
	secret := queryVals.Get("secret")
	passcode := queryVals.Get("totp")
	isValid := utils.Verify(secret, passcode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isValid)
}

func generateTOTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/generateTOTP.html"))
	tmpl.Execute(w, nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/test.html"))
	tmpl.Execute(w, nil)
}

func apiUrl(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/apiUrl.html"))
	tmpl.Execute(w, nil)
}
