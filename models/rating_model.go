package models

import (
	"database/sql"
	"entities"
)

//RatingModel Struct
type RatingModel struct {
	Db *sql.DB
}

//FindRating Method
func (ratingModel RatingModel) FindRating() (rating []entities.Rating, err error) {
	rows, err := ratingModel.Db.Query("SELECT department.id,ratings.rating,ratings.departmentId,department.name_mm FROM ratings INNER JOIN department ON ratings.departmentId=department.id")
	if err != nil {
		return nil, err
	} else {
		var ratings []entities.Rating
		for rows.Next() {
			var id int64
			var rating int64
			var departmentID int64
			var departmentname string
			// var createddate string
			err2 := rows.Scan(&id, &rating, &departmentID, &departmentname)
			if err2 != nil {
				return nil, err2
			} else {
				rating := entities.Rating{
					ID:             id,
					Rating:         rating,
					DepartmentID:   departmentID,
					DepartmentName: departmentname,
					// CreatedDate: createddate,
				}
				ratings = append(ratings, rating)
			}
		}
		return ratings, nil
	}
}

// //FilterRating Method
// func (ratingModel RatingModel) FilterRating()(rating){

// }
