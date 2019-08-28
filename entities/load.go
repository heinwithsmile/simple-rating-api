package entities

import (
	"fmt"
)

//Load Struct
type Load struct {
	ID             int64  `json:"id"`
	Rating         int64  `json:"rating"`
	DepartmentID   int64  `json:"departmentID"`
	DepartmentName string `json:"departmentname"`
	// CreatedDate string `json:"createdate"`
}

//ToString Function
func (load Load) ToString() string {
	return fmt.Sprintf("id: %d\nrating:%d\ndepartmentID:%d\ndepartmentname:%s\n", load.ID, load.Rating, load.DepartmentID, load.DepartmentName)
}