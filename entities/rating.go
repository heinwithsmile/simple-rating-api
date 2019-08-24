package entities

import (
	"fmt"
)

//Rating Struct
type Rating struct {
	ID             int64  `json:"id"`
	DepartmentName string `json:"departmentname"`
	Rating         int64  `json:"rating"`
	DepartmentID   int64  `json:"departmentID"`
	// CreatedDate string `json:"createdate"`
}

//ToString Function
func (rating Rating) ToString() string {
	return fmt.Sprintf("id: %d\nrating:%d\ndepartmentID:%d\ndepartmentname:%s\n", rating.ID, rating.Rating, rating.DepartmentID, rating.DepartmentName)
}
