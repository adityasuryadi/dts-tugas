package main

import (
	"fmt"
	"os"
)

type student struct {
	noAbsen string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	students := []student{
		student{
			noAbsen:"1",
			nama: "adit",
			alamat:"bandung",
			pekerjaan: "pelajar",
			alasan: "idk",
		},
		student{
			noAbsen:"2",
			nama: "budi",
			alamat: "semarang",
			pekerjaan: "marketing",
			alasan: "alasan",
		},
	}
	
	no_absen := os.Args[1]
	result := make([]student,0)
	for _, v := range students {
		if no_absen == v.noAbsen {
			result = append(result,v)
		}
	}
	fmt.Println(result)
}