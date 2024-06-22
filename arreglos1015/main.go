package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/MarcosIgnacio/utils"
)

// Practica 1
// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez

// Realizar un programa utilizando el paradigma imperativo, que permita ingresar 10 números enteros en un arreglo y 15 números enteros en otro arreglo.

// Se debe revisar que los números estén ordenados en forma creciente, de no ser asi,
// se deben ordenar. Mostrar los arreglos ordenados.
// Ejemplo
// arreglo [3,6,4,1]
// resultado de esta instruccion [1,3,4,6]
// Una vez que se están ordenados ambos arreglos, obtener un tercer arreglo ordenado en forma creciente intercalando los valores de los dos primeros arreglos.
// Repetir tantas veces como el usuario desee hacerlo.

// [1,2,3,4,5]
// [6,7,8,9,10,23,12,32,45,34]
// [1,7,3,4,5]
// [6,2,8,4,5]

func main() {
	Begin()
}

func Begin() {
	fmt.Println("Ingrese los elementos del arreglo de 10 elementos")
	tenLenghtArray := fillArray(10)
	fmt.Println("Ingrese los elementos del arreglo de 15 elementos")
	fiftTeenLenghtArray := fillArray(15)
	fmt.Println("No Ordenados: ")
	printArrays(tenLenghtArray, fiftTeenLenghtArray)
	fmt.Println("Ordenados: ")
	utils.QuickSort(&tenLenghtArray)
	utils.QuickSort(&fiftTeenLenghtArray)
	printArrays(tenLenghtArray, fiftTeenLenghtArray)
	fmt.Println("Arreglo intercalado")
	thirdArray := intercalateArrays(&tenLenghtArray, &fiftTeenLenghtArray)
	fmt.Println(thirdArray)
	res, err := Read("Quieres hacerlo de nuevo? [Escribe Si]", os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
	if *res == "Si" {
		Begin()
	}
}

func intercalateArrays(firstArr *[]int, secondArr *[]int) []int {
	firstArrLen := float64(len(*firstArr))
	secondArrLen := float64(len(*secondArr))
	limit := int(math.Min(firstArrLen, secondArrLen))
	thirdArr := make([]int, limit)
	for i := 0; i < limit; i++ {
		var intercalatedElement int
		if i%2 == 0 {
			intercalatedElement = (*firstArr)[i]
		} else {
			intercalatedElement = (*secondArr)[i]
		}
		thirdArr[i] = intercalatedElement
	}
	return thirdArr
}

func printArrays(firstArr []int, secondArr []int) {
	firstMsg := fmt.Sprintf("Array with lenght %d: ", len(firstArr))
	secondMsg := fmt.Sprintf("Array with lenght %d: ", len(secondArr))
	fmt.Println(firstMsg, firstArr)
	fmt.Println(secondMsg, secondArr)
}

func generateFillerFunction(limit int) func() (bool, *[]int) {
	count := 0
	limiter := limit
	in := os.Stdin
	out := os.Stdout
	arr := make([]int, limit)
	return func() (bool, *[]int) {
		if count < limiter {
			msg := fmt.Sprintf("Ingrese el numero para el indice: %d", count)
			input, err := Read(msg, in, out)
			if err != nil {
				fmt.Println(err)
			}
			number, err := strconv.Atoi(*input)
			if err != nil {
				fmt.Println(err)
			}
			arr[count] = number
		} else {
			return true, &arr
		}
		count++
		return false, nil
	}
}
func fillArray(limit int) []int {
	fillArrayCLI := generateFillerFunction(limit)
	isFilled, arr := fillArrayCLI()
	for !isFilled {
		isFilled, arr = fillArrayCLI()
	}
	return *arr
}

func Start(in io.Reader, out io.Writer, condition bool) {
	read, err := Read("Hola", in, out)
	if err != nil {
		fmt.Println(err)
		return
	}
	for *read != "x" {
		read, err = Read("Este es un mensaje al inicio del ciclo", in, out)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(*read)
	}
}

func Read(message string, in io.Reader, out io.Writer) (*string, error) {
	const PROMPT = ">> "
	fmt.Println(message)
	scanner := bufio.NewScanner(in)
	scanned := scanner.Scan()
	if !scanned {
		return nil, errors.New("Ocurrio un error")
	}
	input := scanner.Text()
	return &input, nil
}
