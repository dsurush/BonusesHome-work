package main

import "fmt"

func BonusAmount(percent int, salary []int) int {
	SumOfBonuses := 0
	for _, employee := range salary{
		if employee > 10_000{
			bonus := employee - 10_000
			bonus = bonus * percent / 100
			SumOfBonuses = SumOfBonuses + bonus
		}
	}
	return SumOfBonuses
}

func main() {
	salary := []int{12000, 8000, 15000, 8000}
	fmt.Println(BonusAmount(5, salary))
}
