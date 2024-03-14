package utils

func GetGreatestNumberInArray(arr []int) int {
	buffer := arr[1]
	if len(arr) == 0 {
		return arr[0]
	}
	if len(arr) == 2 {
		first := BoolToInt(arr[0] > arr[1])
		second := BoolToInt(arr[1] > arr[0])
		return first*arr[0] + second*arr[1]
	}
	for i := 2; i < len(arr); i++ {
		if buffer < arr[i] {
			tmp := buffer
			buffer = arr[i]
			arr[i] = tmp
		}
	}
	return buffer
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func CreateBidimensionalArray(dy int, dx int) [][]int {
	biArr := make([][]int, dy)
	for i := 0; i < len(biArr); i++ {
		biArr[i] = make([]int, dx)
	}
	return biArr
}
