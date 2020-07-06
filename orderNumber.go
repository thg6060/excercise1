package main

import (
	"errors"
	"fmt"
	"sort"
)

func checkOutofRange(begin int, end int, sl []int) error {
	if end > len(sl) || begin < 0 {
		return errors.New("out of range")
	}
	return nil
}

func copySlice(begin int, end int, sl []int) ([]int, error) {
	err := checkOutofRange(begin, end, sl)
	if err != nil {
		return nil, err
	}
	return sl[begin:end], nil
}

func orderNumber() {
	slNumber := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	sort.Ints(slNumber)
	fmt.Println(slNumber)
	//copySl, err1 := copySlice(1, 8, slNumber)
	//fmt.Println(copySl)

	copySl1, err2 := copySlice(1, 16, slNumber)

	fmt.Println(copySl1)
	//Xu ly panic
	//fmt.Println(err1)
	fmt.Println(err2)

}
