package main

import (
	"errors"
	"fmt"
	"strings"
)

func filterString(str string, sl []Serve) error {
	str1 := strings.ToLower(str)
	if sl != nil {
		for _, item := range sl {
			str2 := strings.ToLower(item.Class)
			if strings.Contains(str2, str1) {
				fmt.Println(item)
			}
		}
		return nil
	}
	return errors.New("empty slice ")

}
