package entities

import (
	"fmt"
)

//Rating Struct
type Rating struct {
	// ID             int64  `json:"id"`
	// DepartmentName string `json:"departmentname"`
	Rating         int64  `json:"rating"`
	DepartmentID   int64  `json:"departmentID"`
	// CreatedDate string `json:"createdate"`
}

//ToString Function
func (rating Rating) ToString() string {
	return fmt.Sprintf("rating:%d\ndepartmentID:%d\n",rating.Rating,rating.DepartmentID)
}
