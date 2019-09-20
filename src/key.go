package main

import "fmt"

func main() {
	str := "xrlvf23xfqwsxsqf"
	strsli := []byte(str)
	//fmt.Println(strsli)
	//fmt.Println(string(strsli))
	for i := 1; i < 26; i++ {
		newSli := make([]byte, len(strsli))
		for j := 0; j < len(strsli); j++ {
			if strsli[j]+byte(i) < 123 {
				newSli[j] = strsli[j] + byte(i)
			} else {
				newSli[j] = strsli[j] + byte(i) - 26
			}
		}
		fmt.Println(string(newSli))
	}
}
