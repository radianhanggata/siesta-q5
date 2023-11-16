package model

type ResponseObject struct {
	TotalSick            int `json:"total_sick"`
	TotalPermittedLeaves int `json:"total_permitted_leaves"`
	TotalPresent         int `json:"total_present"`
	Student              Student
}

type ResponseList struct {
	Data []ResponseObject
}
