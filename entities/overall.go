package entities

import (
	"fmt"
)

//OverALLData Struct
// type OverALLData struct {
// 	ID             int64   `json:"id"`
// 	DepartmentName string  `json:"departmentname"`
// 	Average        float64 `json:"average"`

// 	// CreatedDate string `json:"createdate"`
// }

// //ToString Function
// func (data OverALLData) ToString() string {
// 	return fmt.Sprintf("id:%d\naverage:%f\ndepartmentname:%s\n", data.ID, data.Average, data.DepartmentName)
// }

//OverALLData Struct
type OverALLData struct {
	ID        int64   `json:"id"`
	NameMM    string  `json:"namemm"`
	NameEN    string  `json:"nameen"`
	AvgRating float64 `json:"avgrating"`
	Sum       int64   `json:"sum"`
	VeryPoor  int64   `json:"verypoor"`
	Poor      int64   `json:"poor"`
	Average   int64   `json:"average"`
	Good      int64   `json:"good"`
	VeryGood  int64   `json:"verygood"`
	Vppercent string `json:"vppercent"`
	Ppercent string `json:"ppercent"`
	Avgpercent string `json:"avgpercent"`
	Gpercent string `json:"gpercent"`
	Vgpercent string `json:"vgpercent"`
	// CreatedDate string `json:"createdate"`
}

//ToString Function
func (data OverALLData) ToString() string {
	return fmt.Sprintf("id:%d\nnamemm:%s\nnameen:%s\navgrating:%f\nsum:%d\nverypoor:%d\npoor%d\naverage:%d\ngood:%d\nverygood:%d\nvppercent:%f\nppercent:%f\navgpercent:%f\ngdpercent:%f\nvgpercent:%f\n", data.ID, data.NameMM, data.NameEN, data.AvgRating, data.Sum, data.VeryPoor, data.Poor, data.Average, data.Good, data.VeryGood,data.Vppercent,data.Ppercent,data.Avgpercent,data.Gpercent,data.Vgpercent)
}
