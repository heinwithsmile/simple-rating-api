package entities

import (
	"fmt"
)

//Rating Struct
type Rating struct {
	// Rating         int64  `json:"rating"`
	Avg float64 `json:"avg"`
	// CreatedDate string `json:"createdate"`
}

//ToString Function
func (rating Rating) ToString() string {
	return fmt.Sprintf("avg:%f\n", rating.Avg)
}
