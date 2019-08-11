package entities

import (
	"fmt"
)

//Rating Struct
type Rating struct {
	ID             int64  `json:"id"`
	Rating         int64  `json:"rating"`
	DepartmentID   int64  `json:"departmentID"`
	DepartmentName string `json:"departmentname"`
	// CreatedDate string `json:"createdate"`
}

//ToString Function
func (rating Rating) ToString() string {
	return fmt.Sprintf("id: %d\nrating:%d\ndepartmentID:%d\ndepartmentname:%s\n", rating.ID, rating.Rating, rating.DepartmentID, rating.DepartmentName)
}
