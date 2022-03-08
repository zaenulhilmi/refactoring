package main

import (
    "fmt"
    "time"
)

func printOwing(invoice Invoice) {
    printBanner()
    outstanding := calculateOutstanding(invoice)
    recordDueDate(&invoice)
    printDetails(invoice, outstanding)
}

func calculateOutstanding(invoice Invoice) int{
    outstanding := 0
    for _, v := range invoice.orders {
        outstanding += v.amount
    }
    return outstanding
}

func printBanner() {
    fmt.Println("*************************")
    fmt.Println("***** Customer Owes *****")
    fmt.Println("*************************")
}

func printDetails(invoice Invoice, outstanding int) {
    fmt.Println("name:", invoice.customer)
    fmt.Println("amount:", outstanding)
    fmt.Println("due:", invoice.dueDate.Format("2006-01-02"))
}

func recordDueDate(invoice *Invoice) {
    today := time.Now()
    dueDate := today.AddDate(0, 0, 30)
    invoice.dueDate = dueDate
}

/*----------------------------------------------------*/
// Tidak perlu diupdate
/*----------------------------------------------------*/
func main() {
    invoice := Invoice{
        customer: "John",
        orders: []*Order{
            {amount: 100},
            {amount: 200},
        },
    }
    printOwing(invoice)
}

type Invoice struct {
    customer string
    dueDate time.Time
    orders []*Order
}

type Order struct {
    amount int
}
