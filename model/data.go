package model

type Data struct {
	Classes         int
	Students        []Student `json:"students"`
	Sicks           []Student `json:"sicks"`
	PermittedLeaves []Student `json:"permitted_leaves"`
}
