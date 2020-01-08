package main

import "fmt"

func BonusAmount(percent int, salary []int) int {
	sumOfBonuses := 0
	for _, employee := range salary{
		const bonusBorder = 10_000
		if employee > bonusBorder{
			bonus := employee - bonusBorder
			bonus = bonus * percent / 100
			sumOfBonuses = sumOfBonuses + bonus
		}
	}
	return sumOfBonuses
}


func main() {
	salary := []int{12000, 8000, 15000, 8000}
	fmt.Println(BonusAmount(5, salary))
}
