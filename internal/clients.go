package internal

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

type APIClient struct {
	client *http.Client
	baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		client: &http.Client{
		},
		baseURL: baseURL,
	}
}

func (a *APIClient) GetPowerHeroes(id string) (*Item, error){
	
	req, err := http.NewRequest(http.MethodGet, a.baseURL + id + Powerstats, nil)	
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

//параметр ref в запросе отвечает за выбор между: image, connections, work, appearance, biography, powerstats. Например, выбрав work мы получим следующий запрос:
//https://superheroapi.com/api.php/4b3e7de93f96e6c75ce7e09a504a7c6b/500/work  
//если не указывать ref, а только id, то произойдет поиск по id героя
func (a *APIClient) GetHero(id, ref string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, a.baseURL + id + ref, nil)
	if err != nil {
		return nil, err
	}	
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	
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
		
	return resp, err
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