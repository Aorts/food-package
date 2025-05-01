package main

import (
	"fmt"

	"github.com/Aorts/food-package/calculator"
)

func main() {
	for {
		var option int
		fmt.Println("what would you like to do?")
		fmt.Println("1. Ordering a meal?")
		fmt.Println("2. Checking out?")
		fmt.Println("3. Canceling an order?")
		fmt.Println("4. Exit")
		fmt.Println("Please enter your option: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			orderFood()
		case 2:
			checkOut()
		case 4:
			fmt.Println("Exiting the program")
			return
		default:
			fmt.Println("Invalid option, please try again")
		}
	}

}

func orderFood() {
	for {
		fmt.Println("Here you can order a meal")
		fmt.Println("1. Red set - 50 THB/set")
		fmt.Println("2. Green set - 40 THB/set")
		fmt.Println("3. Blue set - 30 THB/set")
		fmt.Println("4. Yellow set - 50 THB/set")
		fmt.Println("5. Pink set - 80 THB/set")
		fmt.Println("6. Purple set - 90 THB/set")
		fmt.Println("7. Orange set - 120 THB/set")
		fmt.Println("8. Back to main menu")
		fmt.Println("Please select No the set you want to order: ")
		var mealOption int
		fmt.Scan(&mealOption)
		switch mealOption {
		case 1:
			fmt.Println("You have selected Red set")
			q := quantity()
			if err := calculator.OrderFood("red", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 2:
			fmt.Println("You have selected Green set")
			q := quantity()
			if err := calculator.OrderFood("green", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 3:
			fmt.Println("You have selected Blue set")
			q := quantity()
			if err := calculator.OrderFood("blue", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 4:
			fmt.Println("You have selected Yellow set")
			q := quantity()
			if err := calculator.OrderFood("yellow", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 5:
			fmt.Println("You have selected Pink set")
			q := quantity()
			if err := calculator.OrderFood("pink", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 6:
			fmt.Println("You have selected Purple set")
			q := quantity()
			if err := calculator.OrderFood("purple", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 7:
			fmt.Println("You have selected Orange set")
			q := quantity()
			if err := calculator.OrderFood("orange", q); err != nil {
				fmt.Println(err)
				continue
			}
		case 8:
			fmt.Println("Returning to main menu")
			return
		default:
			fmt.Println("Invalid option, please try again")
		}
	}
}

func quantity() int {
	var quantity int
	fmt.Println("Please enter the quantity: ")
	fmt.Scan(&quantity)
	if quantity < 1 {
		fmt.Println("Invalid quantity, please try again")
		return 0
	}
	return quantity
}

func checkOut() {
	var isMember bool
	fmt.Println("Are you a member? (true/false)")
	fmt.Scan(&isMember)
	total := calculator.CheckOut(isMember)
	fmt.Printf("Total amount: %.2f\n", total)
}
