package main

import (
	"testing"
	"testapihoroes/internal"
	"testapihoroes/helpers"
	
	"net/http"
		
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestGetId(t *testing.T) {
	client := internal.NewAPIClient(internal.Base_url)
	runner.Run(t, "Get hero with id=66", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по id 66
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response, _ := client.GetHero("66", "")
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
		response, _ := client. GetHero("500", internal.Powerstats)
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
		response, _ := client.GetHero("489", internal.Biography)
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
		response, _ := client.GetHero("202", internal.Appearance)
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
		response, _ := client.GetHero("156", internal.Work)
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
		response, _ := client.GetHero("28", internal.Connections)
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
		response, _ := client.GetHero("707", "/image")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
}

func TestGetSearch(t *testing.T){
	client := internal.NewAPIClient(internal.Url_search)	
	runner.Run(t, "Get search hero with name=batman", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 200 при запросе к api по name batman
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response, _ := client.GetHero("", "batman")
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusOK
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
}

func TestGetError(t *testing.T) {
	client := internal.NewAPIClient(internal.Base_url)
	runner.Run(t, "Get not found hero id", func(t provider.T) {
		t.Parallel()
		t.Description(`
		Получение кода ответа 404 при запросе к api по не существующему id
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response, _ := client.GetHero("1000", "")
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
		response, _ := client.GetHero("", internal.Image)
		t.NewStep("Шаг 2. Указываем реальный ответ запроса и ожидаемый")
		got := response.StatusCode
		want := http.StatusNotFound
		t.NewStep("Шаг 3. Сравниваем ожидаемый ответ запроса и реальный")
		assertStatus(t, got, want)
	})
	
	runner.Run(t, "Access denied", func(t provider.T) {
		t.Parallel()
		client := internal.NewAPIClient(internal.Url_without_access)
		t.Description(`
		Получение кода ответа 403 при запросе к api без указания токена доступа
		`)
		t.NewStep("Шаг 1. Получаем ответ на запрос")
		response, _ := client.GetHero("1", internal.Image)
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
		response, _ := client.GetHero("321", "/artwe")
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
		got := helpers.MatchHeroes("644", "332")
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
		got := helpers.MatchHeroes("532", "346")
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
		got := helpers.MatchHeroes("241", "730")
		want := "Hero One equal in power Hero Two"
		t.NewStep("Шаг 2. Сравниваем реальный результат с ожидаемым")
		assertMatchingHero(t, got, want)
	})
}

func assertStatus(t provider.T, got, want int) {
	t.Assert().Equal(got, want, "Did not get correct status")
}

func assertMatchingHero(t provider.T, got, want string) {
	t.Assert().Equal(got, want, "Assertion Failed")
}