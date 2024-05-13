package main

import (
	"fmt"
	"strings"
	"time"
)

var (
	productName        []string
	quantity           []int64
	itemPrice          []float64
	total              []float64
	discountPercentage float64
	amountPaid         float64
	subTotal           float64
	discountedPrice    float64
	vat                float64
	billTotal          float64
	itemBalance        float64
	cashierName        string
	customerName       string
	date               = time.Now()
)

func collectInput() {
	fmt.Println("What is your name?")
	_, err := fmt.Scanln(&customerName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {
		fmt.Println("What did the user buy?")
		var name string
		_, err = fmt.Scanln(&name)
		if err != nil {
			return
		}
		productName = append(productName, name)

		fmt.Println("How many pieces?")
		var piece int64
		_, err = fmt.Scanln(&piece)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		quantity = append(quantity, piece)

		fmt.Println("How much per unit?")
		var unitPrice float64
		_, err = fmt.Scanln(&unitPrice)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		itemPrice = append(itemPrice, unitPrice)

		fmt.Println("Add more item? (yes/no)")
		var status string
		_, err := fmt.Scanln(&status)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if strings.ToLower(status) != "yes" {
			break
		}
	}

	fmt.Println("What is your name?")
	_, err = fmt.Scanln(&cashierName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("How much discount will he/she get?")
	_, err = fmt.Scanln(&discountPercentage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func calculateTotalForEachProduct() {
	for index := 0; index < len(quantity); index++ {
		totalPrice := itemPrice[index] * float64(quantity[index])
		total = append(total, totalPrice)
	}
}

func calculateSubTotal() {
	for index := 0; index < len(total); index++ {
		subTotal += total[index]
	}
}

func calculateDiscountedPrice() {
	discountedPrice = (discountPercentage / 100) * subTotal
}

func calculateVat() {
	vat = (17.50 / 100) * subTotal
}

func calculateBillTotal() {
	billTotal = subTotal - discountedPrice + vat
}

func calculateBalance() {
	fmt.Println("How much did the customer give to you?")
	_, err := fmt.Scanln(&amountPaid)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	itemBalance = amountPaid - billTotal
}
func printReceipt() {
	fmt.Printf("SEMICOLON STORES\nMAIN BRANCH\nLOCATION: 312, HERBERT MACAULAY WAY, SABO YABA LAGOS.\nTEL: 03293828343\nDate: %s\nCashier: %s\nCustomer Name: %s\n", date.Format("02-01-2006"), cashierName, customerName)
	fmt.Println("===================================")
	fmt.Println("ITEM\tQTY\tPRICE\tTOTAL(NGN)")
	fmt.Println("====================================")
	for index := 0; index < len(productName); index++ {
		fmt.Printf("%s\t%d\t%.2f\t%.2f\n", productName[index], quantity[index], itemPrice[index], total[index])
	}
	fmt.Println("=====================================")
	fmt.Printf("Sub total: %.2f\n", subTotal)
	fmt.Printf("Discount: %.2f\n", discountedPrice)
	fmt.Printf("VAT @ 17.50%%: %.2f\n", vat)
	fmt.Println("======================================")
}

func printReceipt1() {
	fmt.Printf("SEMICOLON STORES\nMAIN BRANCH\nLOCATION: 312, HERBERT MACAULAY WAY, SABO YABA LAGOS.\nTEL: 03293828343\nDate: %s\nCashier: %s\nCustomer Name: %s\n", date.Format("02-01-2006"), cashierName, customerName)
	fmt.Println("========================================")
	fmt.Println("ITEM\tQTY\tPRICE\tTOTAL(NGN)")
	fmt.Println("=========================================")
	for index := 0; index < len(productName); index++ {
		fmt.Printf("%s\t%d\t%.2f\t%.2f\n", productName[index], quantity[index], itemPrice[index], total[index])
	}
	fmt.Println("=========================================")
	fmt.Printf("Sub total: %.2f\n", subTotal)
	fmt.Printf("Discount: %.2f\n", discountedPrice)
	fmt.Printf("VAT @ 17.50%%: %.2f\n", vat)
	fmt.Println("======================================")
	fmt.Printf("Bill Total: %.2f\n", billTotal)
	fmt.Printf("Amount paid: %.2f\n", amountPaid)
	fmt.Printf("Balance: %.2f\n", itemBalance)
	fmt.Println("================================")
	fmt.Println("THANK YOU FOR YOUR PATRONAGE")
	fmt.Println("================================")
}

func main() {
	collectInput()
	calculateTotalForEachProduct()
	calculateSubTotal()
	calculateDiscountedPrice()
	calculateVat()
	calculateBillTotal()
	printReceipt()
	calculateBalance()
	printReceipt1()
}
