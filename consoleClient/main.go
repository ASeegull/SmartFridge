package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID        string `json:"productID"`
	Name      string `json:"productName"`
	ShelfLife int    `json:"productShelfLife"`
	Units     string `json:"productUnit"`
	Image     string `json:"image"`
}

type Recepie struct {
	ID              string   `json:"-"`
	RecName         string   `json:"title"`
	Description     string   `json:"description"`
	CoockingTimeMin int      `json:"coockingTimeMin"`
	Complexity      string   `json:"complexity"`
	Ingred          []string `json:"ingredients"`
}

var session string
var URL string

//GetAllRecipes() displays all stored recipes with ingredients, do need authentification
func GetAllRecipes() {
	resp, err := http.Get(URL + "/allRecipes")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	var recipes []Recepie
	json.NewDecoder(resp.Body).Decode(&recipes)
	for k, v := range recipes {
		fmt.Printf("|%-2d|%-10s|%s|%-7s|%-2d|%-20s|\n", k, v.RecName, v.Description, v.Complexity, v.CoockingTimeMin, v.Ingred)
	}
}

func GetProductByName(name string) {
	req, err := http.NewRequest("GET", URL+"/products/getByName/"+name, nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	if resp.Status == "200 OK" {
		var p Product
		json.NewDecoder(resp.Body).Decode(&p)
		fmt.Printf("|%-10s|%-2d|%-7s|%s|%-50s|\n", p.Name, p.ShelfLife, p.Units, p.ID, p.Image)
	} else {
		fmt.Println(err)
	}
}

func GetProductByID(id string) (*Product, error) {
	req, err := http.NewRequest("GET", URL+"/products/getByID/"+id, nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	if resp.Status == "200 OK" {
		var p Product
		json.NewDecoder(resp.Body).Decode(&p)
		fmt.Printf("|%-10s|%-2d|%-7s|%s|%-50s|\n", p.Name, p.ShelfLife, p.Units, p.ID, p.Image)
		return &p, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

func DeleteProductByID(id string, session string) {
	req, err := http.NewRequest("DELETE", URL+"/products/remove/"+id, nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response Status:", resp.Status)
}

func GetAllProducts() {
	req, err := http.NewRequest("GET", URL+"/getProducts", nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	var products []Product
	json.NewDecoder(resp.Body).Decode(&products)
	for k, v := range products {
		fmt.Printf("|%-2d|%-10s|%-2d|%-7s|%s|%-50s|\n", k, v.Name, v.ShelfLife, v.Units, v.ID, v.Image)
	}
}

func AddProduct(name string, shelfLife int, image string, unit string) {
	product := &Product{}
	product.Name = name
	product.ShelfLife = shelfLife
	product.Image = image
	product.Units = unit
	url := URL + "/addProduct"
	jsonStr, err := json.Marshal(product)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	cookie := "sessionName" + "=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response Status:", resp.Status)
}

func UpdateProduct(product *Product) {
	url := URL + "/updateProduct"
	jsonStr, err := json.Marshal(product)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response Status:", resp.Status)
}

func Recepies() {
	req, err := http.NewRequest("GET", URL+"/searchRecipes", nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	var recipes []Recepie
	json.NewDecoder(resp.Body).Decode(&recipes)
	for k, v := range recipes {
		fmt.Printf("|%-2d|%-10s|%s|%-7s|%-2d|%-20s|\n", k, v.RecName, v.Description, v.Complexity, v.CoockingTimeMin, v.Ingred)
	}
}

func loginFunc(login string, pass string) {
	var jsonStr = []byte(`{"login":"` + login + `","pass":"` + pass + `"}`)
	req, err := http.NewRequest("POST", URL+"/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("response Status:", resp.Status)
	if resp.Status == "200 OK" {
		str := resp.Header.Get("Set-Cookie")
		r := strings.NewReplacer("=", " ", ";", " ")
		str = r.Replace(str)
		strArray := strings.Fields(str)
		session = strArray[1]
	}
}

func logout() {
	session = ""
}

func isAuthorized() bool {
	if session == "" {
		fmt.Println(errors.New("not authenticated"))
		return false
	}
	return true
}

func printMenu() {
	fmt.Println("====================================================")
	fmt.Println("Please choose:")
	fmt.Println("1 all recepies")
	fmt.Println("2 all products (authentication needed)")
	fmt.Println("3 get product dy name (authentication needed)")
	fmt.Println("4 add product (authentication needed)")
	fmt.Println("5 delete product dy ID (authentication needed)")
	fmt.Println("6 update product (authentication needed)")
	fmt.Println("7 recomended recepies (authentication needed)")
	fmt.Println("8 login")
	fmt.Println("9 logout")
	fmt.Println("10 quit")
	fmt.Println("====================================================")
}

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	URL = config.ServerAddress

	fmt.Println("====================================================")
	fmt.Println("Welcome to the smart frige console client")
	var choice int

	for choice != 10 {
		printMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			GetAllRecipes()
		case 2:
			if !isAuthorized() {
				continue
			}
			GetAllProducts()
		case 3:
			if !isAuthorized() {
				continue
			}
			fmt.Println("Enter produc name")
			var name string
			fmt.Scan(&name)
			GetProductByName(name)
		case 4:
			if !isAuthorized() {
				continue
			}
			var name, unit, image string
			var shelfLife int
			fmt.Println("Enter produc name")
			fmt.Scan(&name)
			fmt.Println("Enter produc shelf life")
			fmt.Scan(&shelfLife)
			fmt.Println("Enter produc image link")
			fmt.Scan(&image)
			fmt.Println("Enter produc measuring unit")
			fmt.Scan(&unit)
			AddProduct(name, shelfLife, image, unit)
		case 5:
			if !isAuthorized() {
				continue
			}
			fmt.Println("Enter produc ID")
			var id string
			fmt.Scan(&id)
			DeleteProductByID(id, session)
		case 6:
			if !isAuthorized() {
				continue
			}
			fmt.Println("Enter produc ID")
			var id, name, shelfLife, image, unit string
			fmt.Scan(&id)
			p, err := GetProductByID(id)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Old product name is " + p.Name + "; Enter new produc name (enter n if you do not wont to change it)")
			fmt.Scan(&name)
			if name != "n" {
				p.Name = name
			}
			fmt.Println("Old product shelf life is " + strconv.Itoa(p.ShelfLife) + "; Enter new produc shelf life (enter n if you do not wont to change it)")
			fmt.Scan(&shelfLife)
			if shelfLife != "n" {
				shelfLifeInt, err := strconv.Atoi(shelfLife)
				if err != nil {
					fmt.Println(err)
					continue
				}
				p.ShelfLife = shelfLifeInt
			}
			fmt.Println("Old image link is " + p.Image + "; Enter new produc image link (enter n if you do not wont to change it)")
			fmt.Scan(&image)
			if image != "n" {
				p.Image = image
			}
			fmt.Println("Old product measuring unit is " + p.Units + "; Enter new produc measuring unit (enter n if you do not wont to change it)")
			fmt.Scan(&unit)
			if unit != "n" {
				p.Units = unit
			}
			UpdateProduct(p)
		case 7:
			if !isAuthorized() {
				continue
			}
			Recepies()
		case 8:
			var login, pass string
			fmt.Println("Enter login")
			fmt.Scan(&login)
			fmt.Println("Enter password")
			fmt.Scan(&pass)
			loginFunc(login, pass)
		case 9:
			if !isAuthorized() {
				continue
			}
			logout()
			fmt.Println("Logged out")
		case 10:
			fmt.Println("Bye!")
		default:
			fmt.Printf("Are you sure you have chosen properly?")
		}
	}
}
