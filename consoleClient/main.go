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
	ID        string `json:"ID"`
	Name      string `json:"name"`
	ShelfLife int    `json:"shelfLife"`
	Units     string `json:"units"`
	Image     string `json:"imageURL"`
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

func getAllRecipes() {
	resp, err := http.Get(URL + "/allRecipes")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	var recipes []Recepie
	json.NewDecoder(resp.Body).Decode(&recipes)
	fmt.Printf("|%-2s|%-10s|%-15s|%-15s|%-25s|%-20s\n", "No", "Name", "Description", "Complexity", "Coocking time in minutes", "Ingredients")
	for k, v := range recipes {
		fmt.Printf("|%-2d|%-10s|%-15s|%-15s|%-25d|%-20s\n", k+1, v.RecName, v.Description, v.Complexity, v.CoockingTimeMin, v.Ingred)
	}
}

func getProductByName(name string) {
	req, err := http.NewRequest("GET", URL+"/products/getByName/"+name, nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode == http.StatusOK {
		var p Product
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("|%-10s|%-10s|%-7s|%-36s|%-36s\n", "Name", "Shelf life", "Units", "ID", "Image URL")
		fmt.Printf("|%-10s|%-10d|%-7s|%-36s|%-250s\n", p.Name, p.ShelfLife, p.Units, p.ID, p.Image)
	} else {
		fmt.Println(resp.Status)
	}
}

func getProductByID(id string) (*Product, error) {
	req, err := http.NewRequest("GET", URL+"/products/getByID/"+id, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		var p Product
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Printf("|%-10s|%-10s|%-7s|%-36s|%-36s\n", "Name", "Shelf life", "Units", "ID", "Image URL")
		fmt.Printf("|%-10s|%-10d|%-7s|%-36s|%-250s\n", p.Name, p.ShelfLife, p.Units, p.ID, p.Image)
		return &p, nil
	} else {
		fmt.Println(resp.Status)
		return nil, err
	}
}

func deleteProductByID(id string, session string) {
	req, err := http.NewRequest("DELETE", URL+"/products/remove/"+id, nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("response Status:", resp.Status)
}

func getAllProducts() {
	req, err := http.NewRequest("GET", URL+"/getProducts", nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	var products []Product
	json.NewDecoder(resp.Body).Decode(&products)
	fmt.Printf("|%-2s|%-10s|%-10s|%-7s|%-36s|%-36s\n", "No", "Name", "Shelf life", "Units", "ID", "Image URL")
	for k, v := range products {
		fmt.Printf("|%-2d|%-10s|%-10d|%-7s|%-36s|%-250s\n", k+1, v.Name, v.ShelfLife, v.Units, v.ID, v.Image)
	}
}

func addProduct(name string, shelfLife int, image string, unit string) {
	product := &Product{Name: name, ShelfLife: shelfLife, Image: strings.Trim(image, " "), Units: unit}
	url := URL + "/addProduct"
	jsonStr, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	cookie := "sessionName" + "=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("response Status:", resp.Status)
}

func updateProduct(product *Product) {
	url := URL + "/updateProduct"
	jsonStr, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("response Status:", resp.Status)
}

func recepies() {
	req, err := http.NewRequest("GET", URL+"/searchRecipes", nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	var recipes []Recepie
	json.NewDecoder(resp.Body).Decode(&recipes)
	fmt.Printf("|%-2s|%-10s|%-15s|%-15s|%-25s|%-20s\n", "No", "Name", "Description", "Complexity", "Coocking time in minutes", "Ingredients")
	for k, v := range recipes {
		fmt.Printf("|%-2d|%-10s|%-15s|%-15s|%-25d|%-20s\n", k+1, v.RecName, v.Description, v.Complexity, v.CoockingTimeMin, v.Ingred)
	}
}

func loginFunc(login string, pass string) {
	var jsonStr = []byte(`{"login":"` + login + `","pass":"` + pass + `"}`)
	req, err := http.NewRequest("POST", URL+"/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("response Status:", resp.Status)
	if resp.StatusCode == http.StatusOK {
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
		log.Println(errors.New("not authenticated"))
		return false
	}
	return true
}

func getRecepiesroductName(productName string) {
	req, err := http.NewRequest("GET", URL+"/recipes/getByProductName/"+productName, nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie := "sessionName=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode == http.StatusOK {
		var recipes []Recepie
		json.NewDecoder(resp.Body).Decode(&recipes)
		fmt.Printf("|%-2s|%-10s|%-15s|%-15s|%-25s|%-20s\n", "No", "Name", "Description", "Complexity", "Coocking time in minutes", "Ingredients")
		for k, v := range recipes {
			fmt.Printf("|%-2d|%-10s|%-15s|%-15s|%-25d|%-20s\n", k+1, v.RecName, v.Description, v.Complexity, v.CoockingTimeMin, v.Ingred)
		}
	} else {
		fmt.Println(resp.Status)
		return
	}
}

func getRecepiesroductNames(productNames []string) {
	jsonStr, err := json.Marshal(productNames)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", URL+"/recipes/recipesByProductNames", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	cookie := "sessionName" + "=" + session
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return
	}
	var recipes []Recepie
	json.NewDecoder(resp.Body).Decode(&recipes)
	fmt.Printf("|%-2s|%-10s|%-15s|%-15s|%-25s|%-20s\n", "No", "Name", "Description", "Complexity", "Coocking time in minutes", "Ingredients")
	for k, v := range recipes {
		fmt.Printf("|%-2d|%-10s|%-15s|%-15s|%-25d|%-20s\n", k, v.RecName, v.Description, v.Complexity, v.CoockingTimeMin, v.Ingred)
	}
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
	fmt.Println("8 find recepies by product name (authentication needed)")
	fmt.Println("9 find recepies by many product names (authentication needed)")
	fmt.Println("10 login")
	fmt.Println("11 logout")
	fmt.Println("12 quit\n====================================================")
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

	for choice != 12 {
		printMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			getAllRecipes()
		case 2:
			if isAuthorized() {
				getAllProducts()
			}
		case 3:
			if isAuthorized() {
				fmt.Println("Enter product name")
				var name string
				fmt.Scan(&name)
				getProductByName(name)
			}
		case 4:
			if isAuthorized() {
				var name, unit, image string
				var shelfLife int
				fmt.Println("Enter product name")
				fmt.Scan(&name)
				fmt.Println("Enter product shelf life")
				fmt.Scan(&shelfLife)
				fmt.Println("Enter product image link")
				fmt.Scan(&image)
				fmt.Println("Enter product measuring unit")
				fmt.Scan(&unit)
				addProduct(name, shelfLife, image, unit)
			}
		case 5:
			if isAuthorized() {
				fmt.Println("Enter product ID")
				var id string
				fmt.Scan(&id)
				deleteProductByID(id, session)
			}
		case 6:
			if isAuthorized() {
				fmt.Println("Enter product ID")
				var id, name, shelfLife, image, unit string
				fmt.Scan(&id)
				p, err := getProductByID(id)
				if err != nil {
					log.Println(err)
					continue
				}

				fmt.Println("Old product name is " + p.Name + "; Enter new product name (enter n if you do not wont to change it)")
				fmt.Scan(&name)
				if name != "n" {
					p.Name = name
				}
				fmt.Println("Old product shelf life is " + strconv.Itoa(p.ShelfLife) + "; Enter new product shelf life (enter n if you do not wont to change it)")
				fmt.Scan(&shelfLife)
				if shelfLife != "n" {
					shelfLifeInt, err := strconv.Atoi(shelfLife)
					if err != nil {
						fmt.Println(err)
						continue
					}
					p.ShelfLife = shelfLifeInt
				}
				fmt.Println("Old image link is " + p.Image + "; Enter new product image link (enter n if you do not wont to change it)")
				fmt.Scan(&image)
				if image != "n" {
					p.Image = strings.Trim(image, " ")
				}
				fmt.Println("Old product measuring unit is " + p.Units + "; Enter new product measuring unit (enter n if you do not wont to change it)")
				fmt.Scan(&unit)
				if unit != "n" {
					p.Units = unit
				}
				updateProduct(p)
			}
		case 7:
			if isAuthorized() {
				recepies()
			}
		case 8:
			if isAuthorized() {
				var name string
				fmt.Println("Enter product name")
				fmt.Scan(&name)
				getRecepiesroductName(name)
			}
		case 9:
			if isAuthorized() {
				var name string
				nameSlice := make([]string, 0)
				for {
					fmt.Println("Enter product name or n case you finished with entering products")
					fmt.Scan(&name)
					if name == "n" {
						break
					}
					nameSlice = append(nameSlice, name)
				}
				getRecepiesroductNames(nameSlice)
			}
		case 10:
			if session != "" {
				fmt.Println("You are logged in")
				continue
			}
			var login, pass string
			fmt.Println("Enter login")
			fmt.Scan(&login)
			fmt.Println("Enter password")
			fmt.Scan(&pass)
			loginFunc(login, pass)
		case 11:
			if isAuthorized() {
				logout()
				fmt.Println("Logged out")
			}
		case 12:
			fmt.Println("Bye!")
		default:
			fmt.Println("Are you sure you have chosen properly?")
		}
	}
}
