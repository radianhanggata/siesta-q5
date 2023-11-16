package main

import (
	"fmt"

	"github.com/radianhanggata/siesta-coding-test/q5/model"
)

func main() {
	data := &model.Data{
		Classes: 14,
		Students: []model.Student{
			{ID: 12, Name: "Fikri Naufal", NISN: "43567200883", Class: "VII", School: "MTs Ali Maksum"},
			{ID: 13, Name: "Moh Rasya", NISN: "4356720042", Class: "X", School: "MA Ali Maksum"},
			{ID: 30, Name: "Naufal Kurniawan", NISN: "63567200868", Class: "XI", School: "MA Ali Maksum"},
			{ID: 39, Name: "Aminuddin Abdallah", NISN: "43567200827", Class: "XII", School: "MA Ali Maksum"},
		},
		Sicks: []model.Student{
			{ID: 39, Name: "Aminuddin Abdallah", NISN: "43567200827", Class: "XII", School: "MA Ali Maksum"},
			{ID: 39, Name: "Aminuddin Abdallah", NISN: "43567200827", Class: "XII", School: "MA Ali Maksum"},
			{ID: 13, Name: "Moh Rasya", NISN: "4356720042", Class: "X", School: "MA Ali Maksum"},
		},
		PermittedLeaves: []model.Student{
			{ID: 12, Name: "Fikri Naufal", NISN: "43567200883", Class: "VII", School: "MTs Ali Maksum"},
		},
	}

	list := CalculateAbsen(data)
	for _, obj := range list.Data {
		fmt.Printf("Student: %v, Sicks: %d, Leaves: %d, Present: %d\n", obj.Student, obj.TotalSick, obj.TotalPermittedLeaves, obj.TotalPresent)
	}
}

func CalculateAbsen(data *model.Data) model.ResponseList {
	list := model.ResponseList{}
	for _, student := range data.Students {
		sicks := total(student.ID, data.Sicks)
		leaves := total(student.ID, data.PermittedLeaves)
		present := data.Classes - (sicks + leaves)
		obj := model.ResponseObject{
			Student:              student,
			TotalSick:            sicks,
			TotalPermittedLeaves: leaves,
			TotalPresent:         present,
		}
		list.Data = append(list.Data, obj)
	}

	return list
}

func total(id int, category []model.Student) int {
	count := 0
	for _, s := range category {
		if s.ID == id {
			count++
		}
	}

	return count
}
