package entities

import (
	"fmt"
)
//Department Struct
type Department struct {
	ID int64 `json:"id"`
	NameMM string `json:"name_mm"`
	NameEN string `json:"name_en"`
}
//ToString Function
func (department Department) ToString() string {
	return fmt.Sprintf("id: %d\nname_mm:%s\nname_en:%s\n", department.ID, department.NameMM,department.NameEN)
}