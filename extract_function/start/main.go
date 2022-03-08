package main

import (
    "fmt"
    "time"
)


func printOwing(invoice Invoice) {

    outstanding := 0
    fmt.Println("*************************")
    fmt.Println("***** Customer Owes *****")
    fmt.Println("*************************")


    // calculate outstanding
    for _, v := range invoice.orders {
        outstanding += v.amount
    }

    // record due date
    today := time.Now()
    dueDate := today.AddDate(0, 0, 30)
    invoice.dueDate = dueDate

    // print details
    fmt.Println("name:", invoice.customer)
    fmt.Println("amount:", outstanding)
    fmt.Println("due:", invoice.dueDate.Format("2006-01-02"))
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
