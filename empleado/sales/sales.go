package sales

import (
	"fmt"
	"strconv"
)

type SalesLog struct {
	Sales [12]float64
}

func (salesLog *SalesLog) String() string {
	var total float64
	for _, sale := range salesLog.Sales {
		total += sale
	}
	return fmt.Sprintf("Tus ventas totales fueron: $%f", total)
}

func NewSalesLog() *SalesLog {
	return &SalesLog{}
}

func (salesLog *SalesLog) InsertSale() {

	for {
		var input string
		fmt.Println("Ingresa el numero del mes cuyas ventas quieres registrar")
		fmt.Println("O `x` para salir")
		fmt.Scanln(&input)
		if input == "x" {
			break
		}
		month, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Has ingresado caracteres invalidos, omitiendo")
			continue
		}

		if month <= 0 || month > 12 {
			fmt.Println("Mes invalido, no puede ser menor a 0 o mayor a 12")
			fmt.Println()
			continue
		}

		fmt.Println("Ingresa el total de ventas")
		fmt.Scanln(&input)

		salesInMonth, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Has ingresado caracteres invalidos, omitiendo")
			continue
		}

		if salesInMonth < 0 {
			fmt.Println("Monto invalido, no puede ser menor a 0")
			continue
		}
		salesLog.Sales[month-1] = salesInMonth
	}
}
