package entities

import (
	"fmt"
)

//Department Struct
type Department struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//ToString Function
func (department Department) ToString() string {
	return fmt.Sprintf("id: %d\nname:%s\n", department.ID, department.Name)
}
