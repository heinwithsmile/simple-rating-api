package models

import (
	"database/sql"
	"entities"
)

//RatingModel Struct
type RatingModel struct {
	Db *sql.DB
}

//GetAllData Method
func (ratingModel RatingModel) GetAllData() (rating []entities.Rating, err error) {
	rows, err := ratingModel.Db.Query("SELECT AVG(rating) avg_admin FROM ratings WHERE ratings.departmentId=1")
	if err != nil {
		return nil, err
	} else {
		var ratings []entities.Rating
		for rows.Next() {
			// var rating int64
			var avg float64
			// var createddate string
			err2 := rows.Scan( &avg)
			if err2 != nil {
				return nil, err2
			} else {
				rating := entities.Rating{
					// Rating:         rating,
					Avg :   avg,
				}
				ratings = append(ratings, rating)
			}
		}
		return ratings, nil
	}
}
//LoadRating Method
func (ratingModel RatingModel) LoadRating() (rating []entities.Load, err error) {
	rows, err := ratingModel.Db.Query("SELECT department.id ,ratings.rating,ratings.departmentId,department.name_mm FROM ratings INNER JOIN department ON ratings.departmentId=department.id")
	if err != nil {
		return nil, err
	} else {
		var loads []entities.Load
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
				load := entities.Load{
					ID:             id,
					Rating:         rating,
					DepartmentID:   departmentID,
					DepartmentName: departmentname,
					// CreatedDate: createddate,
				}
				loads = append(loads, load)
			}
		}
		return loads , nil
	}
}