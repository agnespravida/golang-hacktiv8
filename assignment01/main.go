package main

import (
	"fmt"
	"os"
)

type Student struct {
	id         int32
	name       string
	address    string
	occupation string
	reason     string
}

type StudentBio struct {
	address    string
	occupation string
	reason     string
}

func main() {
	var students = []string{"Agnes", "Roni", "Nada"}
	student := os.Args[1]
	var index *int = nil

	for i, value := range students {
		if value == student {
			var idx int = i
			index = &idx
		}
	}

	if index != nil {
		var studentBio = getStudentData(*index, students[*index])
		fmt.Println("ID: ", studentBio.id)
		fmt.Println("nama: ", studentBio.name)
		fmt.Println("alamat: ", studentBio.address)
		fmt.Println("pekerjaan: ", studentBio.occupation)
		fmt.Println("alasan: ", studentBio.reason)
	} else {
		fmt.Println("Error! Student not exist")
	}
}

func getStudentData(studentID int, name string) Student {
	var bio = []StudentBio{
		{
			address:    "Jakarta",
			occupation: "programmer",
			reason:     "Learning",
		},
		{
			address:    "Bandung",
			occupation: "programmer",
			reason:     "Learning",
		},
		{
			address:    "Yogyakarta",
			occupation: "programmer",
			reason:     "Learning",
		},
	}

	var student Student
	student.id = int32(studentID) + 1
	student.name = name
	student.address = bio[studentID].address
	student.occupation = bio[studentID].occupation
	student.reason = bio[studentID].reason
	return student
}
