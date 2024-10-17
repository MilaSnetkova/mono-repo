package main

import (
	"fmt"
	"time"
)

 var basePrice = 100

func getPrice(city string, date string, celsius int) int {
	price := basePrice
	//если больше 10 градусов 
	if celsius > 10 {
		price = basePrice
	}
	//если равно 10 градусов
	if celsius == 10 {
		price = basePrice +10 
	} else {
		// если температура ниже 10 градусов
		// прибавляем к стоимости по 5  за каждые -2 градуса ниже 10
		price += 10 + (5*((10 - celsius)/2))
	}
	return price
}

func main() {
	city:= "Moscow"
	date := time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC)
	celsius := 2 
	price := getPrice(city, date.Format("2006-01-02"), celsius)
	fmt.Printf("Цена за кофе в %s на %s будет %d\n", city, date.Format("2006-01-02"), price)
}