package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
)

// Francisco Alejandro Alcantar Aviles
// Marcos Ignacio Camacho Gonzalez

// Practica 2
// De 30 alumnos de licenciatura se tienen los siguientes datos:

// Numero de alumno (Entero)
// Unidad1, Unidad2, Unidad3, Unidad4 (Enteros)

// Guardar estos datos en un arreglo bidimensional

// 1
// Mostrar el numero del alumno que tuvo la mejor calificacion de todas y en que unidad la obtuvo

// 2
// Obtener un listado de los alumnos ordenados en forma creciente por el numero de alumno

// 3
// Ingresar un numero de alumno y realizar una busqueda ditomica, si se encuentra mostrar su informacion (numero y calificaciones de las 4 unidades) si no mostrar que el alumno no esta en esa lista
// Implementar este algoritmo de manera recursiva
// Implementar este algoritmo de manera no recursiva

func main() {
	students := createBidimensionalArray(30, 5)
	fillStudentsBidimensionalArray(&students)
	orderMatrix(&students)
	bestGrade := greatestGrade(students)
	fmt.Printf("El alumno con mejor calificacion es: %d con la calificacion de %d en la unidad: %d ", bestGrade[0], bestGrade[1], bestGrade[2])
	fmt.Println()
	printStudentsData(students)
	searchStudentAPI(students)
}

// Interfaz de consola que pide los datos necesarios para la busqueda, da la opcion a elegir del usuario de si quiere realizar su busqueda usando recursividad o de manera normal
func searchStudentAPI(students [][]int) {
	input, err := Read("Ingrese el id del alumno que quiere buscar", os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
	id, err := strconv.Atoi(*input)
	if err != nil {
		fmt.Println("Por favor ingrese un numero valido")
		searchStudentAPI(students)
		return
	}
	input, err = Read("Como quiere realizar la busqueda (R)ecursiva o (C)iclica", os.Stdin, os.Stdout)
	switch *input {
	case "R":
		BSR := generateBinarySearchFunction(students)
		student, err := BSR(id)
		if err != nil {
			fmt.Println(err)
			break
		}
		printStudent(student)
	case "C":
		student, err := BinarySearch(students, id)
		if err != nil {
			fmt.Println(err)
			break
		}
		printStudent(student)
	}
}

// Funcion para leer por consola el input de un usuario
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

// Ordena una matriz en base a comparar el elemento en el indice 0 de cada arreglo con cada uno, siendo que si el primero es menor se intercambian de lugar
func orderMatrix(matrix *[][]int) {
	for i := 0; i < len(*matrix); i++ {
		for j := 0; j < len(*matrix); j++ {
			if (*matrix)[i][0] < (*matrix)[j][0] {
				tmp := (*matrix)[j]
				(*matrix)[j] = (*matrix)[i]
				(*matrix)[i] = tmp
			}
		}
	}
}

func BinarySearch(haystack [][]int, needle int) ([]int, error) {
	low := 0
	high := len(haystack)

	// La razon por la que el high se mantiene como el mid y no se resta es porque
	// al ser un < el comparador ya esta excluyendo al mid, por ejemplo si el mid fuera 5
	// y el low 1, 1 < 5, ya no estariamos tocando de cualquier manera el 5, sin embargo si el que adquirio el mid fue el low como esta del otro lado del < no se va a excluir automaticamente lpor lo que se tiene que excluir manualmente con un + 1
	for low < high {
		mid := int(math.Floor(float64(low + (high-low)/2)))
		value := haystack[mid][0]
		if value == needle {
			return haystack[mid], nil
		} else if value > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nil, errors.New("No se encontro el alumno")
}

// Generamos la funcion para buscar por medio de binary search en una matriz (Se retorna una funcion para aprovechar la memoria local que ofrecen las closures lo que nos permite tener acceso al ultimo low y high cada vez que hagamos recursividad y asi evitamos tener que estarla pasando como parametro)

func generateBinarySearchFunction(haystack [][]int) func(int) ([]int, error) {
	low := 0
	high := len(haystack)
	var BSR func(int) ([]int, error)
	BSR = func(needle int) ([]int, error) {
		mid := int(math.Floor(float64(low + (high-low)/2)))
		if mid >= len(haystack) {
			return nil, errors.New("No se encontro el alumno")
		}
		value := haystack[mid][0]
		if value == needle {
			return haystack[mid], nil
		} else if value > needle {
			high = mid
			return BSR(needle)
		} else {
			low = mid + 1
			return BSR(needle)
		}
	}
	return BSR
}

// Funcion para imprimir en formato informacion del estudiante
func printStudent(student []int) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tU1\tU2\tU3\tu4\t")
	output := fmt.Sprintf("%v\t%v\t%v\t%v\t%v", student[0], student[1], student[2], student[3], student[4])
	fmt.Fprintln(w, output)
	w.Flush()
}

// Funcion para imprimir en formato informacion de todos los estudiantes
func printStudentsData(students [][]int) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tU1\tU2\tU3\tu4\t")
	for i := 0; i < len(students); i++ {
		id := students[i][0]
		u1 := students[i][1]
		u2 := students[i][2]
		u3 := students[i][3]
		u4 := students[i][4]
		output := fmt.Sprintf("%v\t%v\t%v\t%v\t%v", id, u1, u2, u3, u4)
		fmt.Fprintln(w, output)
	}
	w.Flush()
}

// Llenar el arreglo de los alumnos con califaciones del rango de 0 a 10 y generar sus ids
func fillStudentsBidimensionalArray(students *[][]int) {
	var grade int
	id := 200
	for i := 0; i < len(*students); i++ {
		(*students)[i][0] = id
		id += 76 + rand.Intn(10)
		for j := 1; j < len((*students)[0]); j++ {
			grade = rand.Intn(11)
			(*students)[i][j] = grade
		}
	}
}

// Sirve para crear un arreglo bidimensional
func createBidimensionalArray(dy int, dx int) [][]int {
	biArr := make([][]int, dy)
	for i := 0; i < len(biArr); i++ {
		biArr[i] = make([]int, dx)
	}
	return biArr
}

// Sirve para obtener la mejor calificacion del arreglo de students
func greatestGrade(students [][]int) []int {
	biGradesArray := createBidimensionalArray(len(students), 3)
	for i := 0; i < len(biGradesArray); i++ {
		grade, unit := getGreatestNumberInArray(students[i])
		biGradesArray[i][0] = students[i][0]
		biGradesArray[i][1] = grade
		biGradesArray[i][2] = unit
	}

	bestGrade := biGradesArray[0][1]
	bestAlumn := biGradesArray[0]

	for i := 1; i < len(biGradesArray); i++ {
		if bestGrade < biGradesArray[i][1] {
			bestAlumn = biGradesArray[i]
			bestGrade = biGradesArray[i][1]
		}
	}
	return bestAlumn
}

// Retorna el numero mas grande dentro de un arreglo
func getGreatestNumberInArray(arr []int) (int, int) {
	buffer := arr[1]
	var idx = 1
	for i := 2; i < len(arr); i++ {
		if buffer < arr[i] {
			buffer = arr[i]
			idx = i
		}
	}
	return buffer, idx
}
