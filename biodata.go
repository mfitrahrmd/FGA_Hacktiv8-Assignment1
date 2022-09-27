package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type Student struct {
	No                              int
	Nama, Alamat, Pekerjaan, Alasan string
}

type DataFGA struct {
	Students []Student `json:"students"`
}

// Add new Student to existing DataFGA
func (d DataFGA) AddStudent(Student ...Student) {
	d.Students = append(d.Students, Student...)
}

// Find Student struct field with the longest value
func (d DataFGA) getLongestStudentsValue(n int) int {
	reflectValue := reflect.ValueOf(d.Students[n])

	longest := len(reflectValue.Field(0).String())

	for i := 1; i < reflectValue.NumField(); i++ {
		if len(reflectValue.Field(i).String()) > longest {
			longest = len(reflectValue.Field(i).String())
		}
	}

	if longest%2 != 0 {
		longest += 1
	}

	return longest
}

// display student(s) bio for given No(s)
func (d DataFGA) Display(search ...int) {
	for i := 0; i < len(search); i++ {
		for j := 0; j < len(d.Students); j++ {
			if j == len(d.Students)-1 && d.Students[j].No != search[i] {
				fmt.Printf("!!Error : Student dengan nomor absen %d tidak ditemukan.\n", search[i])
			}
			if d.Students[j].No == search[i] {
				fmt.Printf(`
~~~~~%s~~~~~~
%s Student

Nama      : %s
Alamat    : %s
Pekerjaan : %s
Alasan    : %s
~~~~~%s~~~~~~
`, RepeatChar("~", d.getLongestStudentsValue(j)), RepeatChar(" ", d.getLongestStudentsValue(j)/2), d.Students[j].Nama, d.Students[j].Alamat, d.Students[j].Pekerjaan, d.Students[j].Alasan, RepeatChar("~", d.getLongestStudentsValue(j)))

				break
			}
		}
	}
}

var Biodata DataFGA

func init() {
	// read students data from json file and bind to DataFGA struct
	fileName := "kelas8.json"
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error saat membaca file : %s\n", err.Error())
		fmt.Println("membuat default data student...")
		// if reading file error, use this data instead
		Biodata = *CreateNewDataFGA(Student{
			No:   1,
			Nama: "Rafli Abi Kusuma",
		}, Student{
			No:        2,
			Nama:      "M Fitrah Ramadhan",
			Alamat:    "Palembang",
			Pekerjaan: "Software Developer",
			Alasan:    "Belajar",
		}, Student{
			No:   3,
			Nama: "Agung chumaidi",
		})
	}

	err = json.Unmarshal([]byte(file), &Biodata)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Create new Student, returning Student interface
func CreateNewDataFGA(Student ...Student) *DataFGA {
	return &DataFGA{
		Students: Student,
	}
}

func main() {
	// get arguments from CLI
	args, err := GetArgs()
	if err != nil {
		log.Fatalln(err.Error())
	}

	Biodata.Display(args...)
}
