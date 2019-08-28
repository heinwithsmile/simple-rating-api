package entities

import (
	"fmt"
)

//Department Struct
type Department struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	NameMM string `json:"name_mm"`
}

//ToString Function
func (department Department) ToString() string {
	return fmt.Sprintf("id: %d\nname:%s\nname_mm:%s\n", department.ID, department.Name,department.NameMM)
}
