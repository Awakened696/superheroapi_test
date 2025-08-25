package helpers

import (
	"fmt"
	"strconv"
	
	"testapihoroes/internal"
)

func MatchHeroes(id1, id2 string) string{
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
	
	client := internal.NewAPIClient(internal.Base_url)
	s, err := client.GetPowerHeroes(id)
	if err != nil {
		fmt.Println(err)
	}
	
	number, err := strconv.Atoi(s.Power)
	if err != nil {
		fmt.Println(err)
	}
	return number
}