package entities

import (
	"fmt"
)
//City Struct
type City struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}
//ToString Function
func (city City) ToString() string {
	return fmt.Sprintf("id: %d\nname:%s\n", city.ID, city.Name)
}