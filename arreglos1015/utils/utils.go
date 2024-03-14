package utils

func qS(array *[]int, low int, high int) {
	// Checamos si el low ya ha alcanzado al high, es decir, que ya sean particiones tipo 1,1 7,7 asi
	if low >= high {
		return
	}
	// Si no es asi realizamos la particion
	pivotIdx := partition(array, low, high)
	qS(array, low, pivotIdx-1)
	qS(array, pivotIdx+1, high)
}

func partition(array *[]int, low int, high int) int {
	// Agarramos a nuestro pivot, que en el caso de la primera vez que hagamos esto con el arreglo 9, 7, 2, 3, 4, 1, 5 el pivot sera el 5 porque es el elemento en el high que nosotros le pasamos

	pivot := (*array)[high]
	idx := low - 1
	// El idx es el low - 1 porque vamos a ser inclusivos con nuestro low por lo que vamos a incluir en el rango a nuestro elemento de low, en la primera vez que hagamos esto el low va a valer 0 por lo que el idx sera -1
	for i := low; i < high; i++ {
		if (*array)[i] <= pivot {
			idx++
			tmp := (*array)[i]
			(*array)[i] = (*array)[idx]
			(*array)[idx] = tmp
		}
	}

	// Hacemos un recorrido desde nuestro low hasta nuestro high (0 hasta 7 pero no 7), si nuestro elemento actual en ele ciclo es menor al elemento que se encuentra en nuestro pivot aumentamos en 1 nuestro idx y swapeamos el elemento que ocupa nuestro idx actualmente con el elemento que este ene el ciclo actualmente por ejemplo el primer elemento de nuestro arreglo es 9 por lo qeu esta condicion no se cumple, no es hasta que llegamos al elemento 2 en el indice 2 que esta condicion pues si se cumple, por lo que vamos a intercambiar el primer elemento de nuestro arreglo recordemos que nuestro idx solo aumentara cuando hagamos los cambios porque queremos estar colocando los elementos en el orden adecuado, entonces cambiamos el 9 por el 2, luego el 3 es menor que 5 por lo que entra tambhien dentro de nuestro if por lo que haremos el swap entre el segundo elemento de nuestro arreglo con este que es el elemento en el indice 3, por lo que cambiamos a array[i] por array[idx], luego lo mismo con el 4 y el 1

	// Luego idx habra quedado en el ultimo elemento al que le hicimos swap por lo que estara al lado de nuestro pivot asi que lo que haremos sera actualizar el idx sumandole 1 para obtener el indice del pivot y luego swapeamos nuestro high por el pivot
	idx++
	(*array)[high] = (*array)[idx]
	(*array)[idx] = pivot
	return idx
}

func QuickSort(array *[]int) {
	// Le pasamos nuestro arreglo 9, 7, 2, 3, 4, 1, 5, el low que como este es la primera vez que vamos a hacer el proceso sera el primer elemento y el high sera el ultimo
	qS(array, 0, len(*array)-1)
}
