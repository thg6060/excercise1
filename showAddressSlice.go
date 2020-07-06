package main

import "fmt"

func showAddressSlice(sl []Serve) {
	for i, _ := range sl {
		fmt.Printf("\n%p",&sl[i])
	}
}
