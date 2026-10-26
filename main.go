package main

import (
	"fmt"
)

// Объявляем константы конвертации
const (
	// Курс доллара к евро (примерный курс)
	UsdToEur float64 = 0.92
	
	// Курс доллара к рублю (примерный курс)
	UsdToRub float64 = 92.50
	
	// Рассчитываем курс евро к рублю на основании первых двух констант
	// Логика: Если 1 USD = 0.92 EUR и 1 USD = 92.50 RUB, 
	// то 0.92 EUR = 92.50 RUB. 
	// Следовательно, 1 EUR = 92.50 / 0.92 RUB
	EurToRub = UsdToRub / UsdToEur
)

func main() {
	fmt.Println("--- Калькулятор валют ---")
	
	// Выводим заданные константы
	fmt.Printf("Курс USD к EUR: %.2f\n", UsdToEur)
	fmt.Printf("Курс USD к RUB: %.2f\n", UsdToRub)
	
	// Выводим рассчитанный кросс-курс
	fmt.Printf("Рассчитанный курс EUR к RUB: %.2f\n", EurToRub)
	
	fmt.Println("-------------------------")
}
