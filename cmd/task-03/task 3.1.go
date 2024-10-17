package main

import "fmt"

var basePrice = 100  

func getTemp(city string, date string) float64 {
	if city == "Moscow" && date == "2024-12-01" {
	return 2.0
}
return 10.0 
} 

func getPrice(temp float64) float64 {
	price := basePrice 
	//если больше 10 градусов 
	if temp > 10 {
		return float64(price)
	//если ровно 10 градусов 
	} else if temp == 10 {
		return float64(price + 10)
	} else {
		// если температура ниже 10 градусов
		// прибавляем к стоимости по 5  за каждые -2 градуса ниже 10
		extra := 10 + ((10 - temp) / 2) * 5 
		return float64(price) + extra		 
}
}
func main() {
	city := "Moscow"
	date := "2024-12-01"

	temp := getTemp(city,date)
	finalPrice := getPrice(temp)

	fmt.Printf("Цена за кофе в %s на %s будет %.2f", city, date, finalPrice) 
}
