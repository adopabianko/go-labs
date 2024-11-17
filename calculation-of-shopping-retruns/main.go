package main

import (
	"fmt"
)

func main() {
	var totalPurchase, totalPayment int
	fmt.Print("Total belanja seorang customer: ")
	fmt.Scan(&totalPurchase)
	fmt.Print("Pembeli membayar: ")
	fmt.Scan(&totalPayment)

	result := calculateChange(totalPurchase, totalPayment)

	if result == false {
		fmt.Println("False, kurang bayar")
	} else {
		fmt.Println("Pecahan uang:")
		for fraction, total := range result.(map[int]int) {
			moneyType := "Lembar"
			if fraction < 1000 {
				moneyType = "Koin"
			}

			fmt.Printf("%d %s Rp.%d\n", total, moneyType, fraction)
		}
	}
}

func calculateChange(totalPurchase, totalPayment int) interface{} {
	if totalPayment < totalPurchase {
		return false
	}

	fraction := []int{100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000}
	totalReturn := (totalPayment - totalPurchase) / 100 * 100
	result := make(map[int]int)

	for _, p := range fraction {
		if totalReturn >= p {
			result[p] = totalReturn / p
			totalReturn %= p
		}
	}

	return result
}
