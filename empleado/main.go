package main

// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez

import (
	"fmt"
	"strings"

	"github.com/MarcosIgnacioo/employee"
	"github.com/MarcosIgnacioo/sales"
)

func main() {
	for {
		fmt.Println("Bienvenido vendedor, desea registrar sus ventas?")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "no" || input == "n" {
			break
		}
		salesLog := sales.NewSalesLog()
		salesLog.InsertSale()
		fmt.Println(salesLog)
	}
}

func testEmployee() {
	em, err := employee.NewEmployee("123", "pancho")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(em)
	}
	em, err = employee.NewEmployee("1234567", "sopaknor")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(em)
	}
}
