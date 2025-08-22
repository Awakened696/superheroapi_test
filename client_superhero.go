package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"strings"
)

type Item struct {
	ID   			string   	`json:"id"`
	Name 			string 		`json:"name"`
	Intelligence 	string 		`json:"intelligence"`
	Strength 		string 		`json:"strength"`
	Speed 			string 		`json:"speed"`
	Durability 		string 		`json:"durability"`
	Power 			string 		`json:"power"`
	Combat 			string 		`json:"combat"`
}

//параметр ref в запросе отвечает за выбор между: image, connections, work, appearance, biography, powerstats. Например, выбрав work мы получим следующий запрос:
//https://superheroapi.com/api.php/4b3e7de93f96e6c75ce7e09a504a7c6b/500/work  
//если не указывать ref, а только id, то произойдет поиск по id героя
func GetHero(base_url, id, ref string) *http.Response {
	client := &http.Client{}
	
	req, _ := http.NewRequest(http.MethodGet, base_url + id + ref, nil)	
	resp, _ := client.Do(req)
	
	defer resp.Body.Close()
	
	data, _ := io.ReadAll(resp.Body)
	requestBody := string(data)
	
	if strings.Contains(requestBody, "invalid id") {
		resp.StatusCode = http.StatusNotFound
	}else if strings.Contains(requestBody, "access denied"){
		resp.StatusCode = http.StatusForbidden
	}else if strings.Contains(requestBody, "invalid property request") {
		resp.StatusCode = http.StatusBadRequest
	}
		
	return resp
}

func GetPowerHeroes(id string, item Item) Item{
	client := &http.Client{}
	
	req, _ := http.NewRequest(http.MethodGet, base_url + id + powerstats, nil)	
	resp, _ := client.Do(req)
	
	defer resp.Body.Close()
	
	data, _ := io.ReadAll(resp.Body)
	
	err := json.Unmarshal([]byte(data), &item)
	if err != nil {
		log.Fatal(err)
	}

	return item
}

func main() {
	
	resp, err := http.Get("https://superheroapi.com/api/4b3e7de93f96e6c75ce7e09a504a7c6b/66")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close() 

	data, err := io.ReadAll(resp.Body) 
	if err != nil {
		log.Println(err)
		return
	}
	
	fmt.Printf("%s", data)
}