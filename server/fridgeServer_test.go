package server

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/satori/go.uuid"
)

const (
	adminName     = "admin"
	adminPassword = "admin"
)

var (
	server          *httptest.Server
	newUserLogin    string
	newUserPassword string
	cli             *http.Client
	loginCookie     []*http.Cookie
)

func getServer() *httptest.Server {
	if server != nil {
		return server
	}
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Cannot read config from file with error : %v", err)
	}

	if err = database.InitiateMongoDB(cfg.Mongo); err != nil {
		log.Fatal(err)
	}

	if err = database.InitPostgersDB(cfg.Postgres); err != nil {
		log.Fatal(err)
	}

	handler := newRouter()
	server = httptest.NewServer(handler)
	return server
}

func TestGetRecipes(t *testing.T) {
	server := getServer()
	res, err := http.Get(fmt.Sprintf("%s/client/allRecipes", server.URL))
	if err != nil {
		t.Fatalf("could not send Get: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status must be OK, got : %d", res.StatusCode)
	}
}

func TestSignUp(t *testing.T) {
	server := getServer()
	client := getClientWithoutCookie()
	handlerURL := fmt.Sprintf("%s/client/signup", server.URL)
	newUserLogin = uuid.NewV4().String()[:10]
	newUserPassword = "qwerty"

	tt := []struct {
		name   string
		url    string
		json   string
		status int
	}{
		{"sign up with admin credentials", handlerURL, `{"login":"` + adminName + `","pass":"` + adminPassword + `"}`, http.StatusInternalServerError},
		{"empty json", handlerURL, `""`, http.StatusInternalServerError},
		{"normal sign up", handlerURL, `{"login":"` + newUserLogin + `","pass":"` + newUserPassword + `"}`, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Post(tc.url, "application/json; charset=utf-8", bytes.NewReader([]byte(tc.json)))
			if err != nil {
				t.Errorf("could not send POST: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}

			if response.StatusCode == http.StatusOK {
				loginCookie = response.Cookies()
			}
		})
	}
}

func getClientWithCookie() *http.Client {
	if cli != nil {
		return cli
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	cookieJar, _ := cookiejar.New(nil)
	url, _ := url.Parse(server.URL)
	cookieJar.SetCookies(url, loginCookie)

	return &http.Client{Transport: tr, Jar: cookieJar}
}

func getClientWithoutCookie() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

func TestLogOut(t *testing.T) {
	server := getServer()
	handlerURL := fmt.Sprintf("%s/client/logout", server.URL)
	tt := []struct {
		name   string
		client *http.Client
		url    string
		json   string
		status int
	}{
		{"log out without credentials", getClientWithoutCookie(), handlerURL, `""`, http.StatusUnauthorized},
		{"normal log out", getClientWithCookie(), handlerURL, `{"login":"` + adminName + `","pass":"` + adminPassword + `"}`, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tc.client.Post(tc.url, "application/json; charset=utf-8", bytes.NewReader([]byte(tc.json)))
			if err != nil {
				t.Errorf("could not send POST: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	server := getServer()
	handlerURL := fmt.Sprintf("%s/client/login", server.URL)
	tt := []struct {
		name   string
		url    string
		json   string
		status int
	}{
		{"bad json", handlerURL, `""""`, http.StatusInternalServerError},
		{"bad login", handlerURL, `{"login":"` + adminName + `some other name","pass":"` + adminPassword + `"}`, http.StatusUnauthorized},
		{"bad password", handlerURL, `{"login":"` + adminName + `","pass":"` + adminPassword + `other pass"}`, http.StatusUnauthorized},
		{"normal login", handlerURL, `{"login":"` + adminName + `","pass":"` + adminPassword + `"}`, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := http.Post(tc.url, "byte", bytes.NewReader([]byte(tc.json)))
			if err != nil {
				t.Errorf("could not send Get: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}

			if response.StatusCode == http.StatusOK {
				loginCookie = response.Cookies()
			}
		})
	}
}

func TestGetFoodInfo(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	handlerURL := fmt.Sprintf("%s/client/fridgeContent", server.URL)

	tt := []struct {
		name   string
		url    string
		status int
	}{
		{"normal", handlerURL, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Get(tc.url)
			if err != nil {
				t.Errorf("could not send Get: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestSearchRecipes(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()

	response, err := client.Get(fmt.Sprintf("%s/client/searchRecipes", server.URL))
	if err != nil {
		t.Errorf("could not send GET: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status must be %d, got : %d", http.StatusOK, response.StatusCode)
	}
}

func TestGetProductByName(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	URL := fmt.Sprintf("%s/client/products/getByName/", server.URL)

	tt := []struct {
		name        string
		productName string
		status      int
	}{
		{"normal product name", "onion", http.StatusOK},
		{"bad product name", "noname", http.StatusInternalServerError},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Get(URL + tc.productName)
			if err != nil {
				t.Errorf("could not send Get: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	URL := fmt.Sprintf("%s/client/products/getByID/", server.URL)

	tt := []struct {
		name      string
		productID string
		status    int
	}{
		{"normal product ID", "48c08fa4-19b5-478c-a30c-8f65f5ca87f6", http.StatusOK},
		{"bad product ID", "noname", http.StatusInternalServerError},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Get(URL + tc.productID)
			if err != nil {
				t.Errorf("could not send Get: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestGetRecipesByProductName(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	URL := fmt.Sprintf("%s/client/recipes/getByProductName/", server.URL)

	tt := []struct {
		name        string
		productName string
		status      int
	}{
		{"normal product name", "tomato", http.StatusOK},
		{"bad product name", "noname", http.StatusInternalServerError},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Get(URL + tc.productName)
			if err != nil {
				t.Errorf("could not send Get: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestGetRecipesByProductNames(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	URL := fmt.Sprintf("%s/client/recipes/recipesByProductNames", server.URL)

	tt := []struct {
		name         string
		productNames []string
		status       int
	}{
		{"normal products", []string{"butter", "bread", "sausage"}, http.StatusOK},
		{"no products", []string{}, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.productNames)
			if err != nil {
				t.Errorf("could not marshal products: %v", err)
			}

			response, err := client.Post(URL, "byte", bytes.NewReader(data))
			if err != nil {
				t.Errorf("could not send Post: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()

	response, err := client.Get(fmt.Sprintf("%s/client/getProducts", server.URL))
	if err != nil {
		t.Errorf("could not send GET: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status must be %d, got : %d", http.StatusOK, response.StatusCode)
	}
}

func TestProductAdd(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	URL := fmt.Sprintf("%s/client/addProduct", server.URL)

	tt := []struct {
		name       string
		newProduct *database.Product
		status     int
	}{
		{"no product", &database.Product{}, http.StatusInternalServerError},
		{"normal product", &database.Product{Name: "name", ShelfLife: 2, Units: "grams", Image: "url"}, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.newProduct)
			if err != nil {
				t.Errorf("could not marshal products: %v", err)
			}

			response, err := client.Post(URL, "byte", bytes.NewReader(data))
			if err != nil {
				t.Errorf("could not send Post: %v", err)
			}

			if response.StatusCode != tc.status {
				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

func TestProductUpdate(t *testing.T) {
	server := getServer()
	client := getClientWithCookie()
	URL := fmt.Sprintf("%s/client/updateProduct", server.URL)

	tt := []struct {
		name       string
		newProduct *database.Product
		status     int
	}{
		{"the same product", &database.Product{Name: "name", ShelfLife: 2, Units: "gramm", Image: "url"}, http.StatusInternalServerError},
		{"empty product", &database.Product{}, http.StatusInternalServerError},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.newProduct)
			if err != nil {
				t.Fatalf("could not marshal products: %v", err)
			}

			req, err := http.NewRequest(http.MethodPut, URL, bytes.NewReader(data))
			if err != nil {
				t.Fatalf("could not send Put: %v", err)
			}
			response, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if response.StatusCode != tc.status {
				t.Fatalf("status must be %d, got : %d", tc.status, response.StatusCode)
			}
		})
	}
}

//func TestDeleteProduct(t *testing.T) {
//	server := getServer()
//	client := getClientWithCookie()
//	URL := fmt.Sprintf("%s/client/products/remove/", server.URL)
//
//	tt := []struct {
//		name        string
//		productName string
//		status      int
//	}{
//		{"bad product's name", "bad name", http.StatusInternalServerError},
//		{"normal product's name", "name", http.StatusOK},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			req, err := http.NewRequest("DELETE", URL+tc.productName, nil)
//			if err != nil {
//				t.Errorf("could not create request DELETE: %v", err)
//			}
//
//			response, err := client.Do(req)
//			if err != nil {
//				t.Errorf("could not send DELETE: %v", err)
//			}
//
//			if response.StatusCode != tc.status {
//				t.Errorf("status must be %d, got : %d", tc.status, response.StatusCode)
//			}
//		})
//	}
//}
