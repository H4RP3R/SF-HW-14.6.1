package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrWrongArrSize = fmt.Errorf("array size must be a not negative integer")

func getArrSize(msg string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return -1, err
	}

	size, err := strconv.Atoi(input)
	if err != nil {
		return -1, ErrWrongArrSize
	}

	if size < 0 {
		return size, ErrWrongArrSize
	}

	return size, nil
}

func fillInArray(size int, msg string) []string {
	arr := make([]string, size)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(msg)
	for i := 0; i < size; i++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		arr[i] = strings.TrimSpace(input)
	}

	return arr
}

func intersection(arr1, arr2 []string) []string {
	hashmap := make(map[string]bool, len(arr1))
	resArr := []string{}

	for _, el := range arr1 {
		hashmap[el] = false
	}

	for _, el := range arr2 {
		if processed, ok := hashmap[el]; ok && !processed {
			resArr = append(resArr, el)
			hashmap[el] = true
		}
	}

	return resArr
}

func main() {
	var (
		size1 int = -1
		size2 int = -1
		err   error
	)

	for size1 < 0 {
		size1, err = getArrSize("Enter first array size: ")
		if err != nil {
			fmt.Println(err)
		}
	}

	for size2 < 0 {
		size2, err = getArrSize("Enter second array size: ")
		if err != nil {
			fmt.Println(err)
		}
	}

	array1 := fillInArray(size1, "Enter first array:")
	array2 := fillInArray(size2, "Enter second array:")

	fmt.Printf("Intersection: %v\n", intersection(array1, array2))
}
