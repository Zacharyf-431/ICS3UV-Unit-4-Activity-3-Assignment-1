/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-12-10
	* This file calculates total of items and price in a shopping cart
	*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
		// assign variables
		var numItems int
		var subtotal float64 = 0
		var discount float64 = 0
		var hst float64
		var total float64

		// assign array 
		var itemNames [100]string

		// reader for user input
		reader := bufio.NewReader(os.Stdin)

		// INPUT
		fmt.Print("Enter the number of items you are going to purchase: ")
		numStr, _ := reader.ReadString('\n')
		numStr = strings.TrimSpace(numStr)
		numItems, _ = strconv.Atoi(numStr)


		
		for i := 0; i < numItems; i++ {
			// prompts for name of the items 
			fmt.Printf("Enter the name of item #%d: ", i+1)
			name, _ := reader.ReadString('\n')
			itemNames[i] = strings.TrimSpace(name)

			// prompts for the price of the previous input 
			fmt.Printf("Enter the price of %s: ", itemNames[i])
			priceStr, _ := reader.ReadString('\n')
			priceStr = strings.TrimSpace(priceStr)
			price, _ := strconv.ParseFloat(priceStr, 64)

			subtotal += price // add to subtotal
		}

		// PROCESS:
		// discount
		if subtotal > 350 {
			discount = subtotal * 0.10
		}

		// PROCESS:
		// tax 
		hst = (subtotal - discount) * 0.13
		total = (subtotal - discount) + hst

		// OUTPUT
		fmt.Println("\nYour shopping cart includes:")
		for i := 0; i < numItems; i++ {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(itemNames[i])
		}

		fmt.Printf("\nThe total amount of items in your shopping cart is %d\n", numItems)
		fmt.Printf("The subtotal cost of your shopping trip was $%.2f\n", subtotal)

		if discount > 0 {
			fmt.Printf("You are eligible for a discount of $%.2f\n", discount)
		} else {
			fmt.Println("You are not eligible for a discount.")
		}

		fmt.Printf("The HST is $%.2f\n", hst)
		fmt.Printf("The total is $%.2f\n", total)

		fmt.Println("\nDone.")
}