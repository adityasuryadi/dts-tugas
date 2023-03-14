package main

import "fmt"

type Kalimat struct {
	huruf  string
	jumlah int
}

func FindLetter(needle string, haystack []Kalimat) (bool, int) {
	for index, v := range haystack {
		if v.huruf == needle {
			return true, index
		}
	}
	return false, -1
}

func main() {
	kata := "selamat malam"
	var totals []Kalimat
	for _, v := range kata {
		fmt.Println(string(v))
		isExist, index := FindLetter(string(v), totals)
		if !isExist {
			totals = append(totals, Kalimat{
				huruf:  string(v),
				jumlah: 1,
			})
		} else {
			totals[index].jumlah += 1
		}
	}

	tmp := make(map[string]int)
	for _, v := range totals {
		tmp[v.huruf] = v.jumlah
	}
	fmt.Println(tmp)
}
