package main

import (
	"testing"
	"net/http"
	"fmt"
	"strconv"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

const (
	base_url 			= 	"https://superheroapi.com/api/4b3e7de93f96e6c75ce7e09a504a7c6b/"
	url_search 			= 	"https://superheroapi.com/api/4b3e7de93f96e6c75ce7e09a504a7c6b/search/"
	url_without_access 	= 	"https://superheroapi.com/api/"
	powerstats 			= 	"/powerstats"
	biography 			= 	"/biography"
	appearance 			= 	"/appearance"
	work 				= 	"/work"
	connections 		= 	"/connections"
	image 				= 	"/image"
)

func TestGetId(t *testing.T) {
	runner.Run(t, "Get hero with id=66", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 66
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "66", "")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Get powerstats hero with id=500", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 500 и powerstats
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "500", powerstats)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	}) 
	
	runner.Run(t, "Get biography hero with id=489", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 489 и biography
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "489", biography)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Get appearance hero with id=202", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 202 и appearance
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "202", appearance)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Get work hero with id=156", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 156 и work
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "156", work)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Get connections hero with id=28", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 28 и connections
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "28", connections)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Get image hero with id=707", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 707 и image
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "707", "/image")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Get search hero with name=batman", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по name batman
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(url_search, "", "batman")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
}

func TestGetError(t *testing.T) {
	runner.Run(t, "Get not found hero id", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 404 при запросе к api по не существующему id
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "1000", "")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusNotFound
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Without id", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 404 при запросе к api без указания id
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "", image)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusNotFound
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Access denied", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 403 при запросе к api без указания токена доступа
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(url_without_access, "1", image)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusForbidden 
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
		
	})
	
	runner.Run(t, "Bad request", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 400 при запросе к api c указанием некоректного endpoint
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response := GetHero(base_url, "321", "/artwe")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusBadRequest
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
}

func TestMatchingHeroes(t *testing.T) {
	runner.Run(t, "Superman stronger Hulk", func(t provider.T) {	
		t.Parallel()
		t.Description(`
		Сравниваем power Superman с Hulk
		`)
		t.NewStep("Шаг 1. Записываем в got результат сравнения двух героев")
		got := matchHeroes("644", "332")
		want := "Hero One stronger Hero Two"
		t.NewStep("Шаг 2. Сравниваем реальный результат с ожидаемым")
		assertMatchingHero(t, got, want)
	})
	
	runner.Run(t, "Iron Man stronger Pyro", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Сравниваем power Iron Man с Pyro
		`)
		t.NewStep("Шаг 1. Записываем в got результат сравнения двух героев")
		got := matchHeroes("532", "346")
		want := "Hero Two stronger Hero One"
		t.NewStep("Шаг 2. Сравниваем реальный результат с ожидаемым")
		assertMatchingHero(t, got, want)
	})
	
	runner.Run(t, "Emma Frost equal in power Zatanna", func(t provider.T) {	
		t.Parallel()
		t.Description(`
		Сравниваем power Emma Frost с Zatanna
		`)
		t.NewStep("Шаг 1. Записываем в got результат сравнения двух героев")
		got := matchHeroes("241", "730")
		want := "Hero One equal in power Hero Two"
		t.NewStep("Шаг 2. Сравниваем реальный результат с ожидаемым")
		assertMatchingHero(t, got, want)
	})
}

func matchHeroes(id1, id2 string) string{
	number1 := strInint(id1)
	number2 := strInint(id2)
	
	got := ""
	if number1 > number2 {
		got = "Hero One stronger Hero Two"
	} else if number1 < number2{
		got = "Hero Two stronger Hero One"
	}else {
		got = "Hero One equal in power Hero Two"
	}
	return got
}

func strInint(id string) int {
	var item Item
	s := GetPowerHeroes(id, item)
	
	number, err := strconv.Atoi(s.Power)
	if err != nil {
		fmt.Println(err)
	}
	return number
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertMatchingHero(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}