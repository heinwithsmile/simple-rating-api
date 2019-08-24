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
	rows, err := ratingModel.Db.Query("SELECT departments.id,departments.name,ratings.rating,ratings.departmentId FROM ratings INNER JOIN departments ON ratings.departmentId=departments.id")
	if err != nil {
		return nil, err
	} else {
		var ratings []entities.Rating
		for rows.Next() {
			var id int64
			var departmentname string
			var rating int64
			var departmentID int64
			// var createddate string
			err2 := rows.Scan(&id, &departmentname, &rating, &departmentID)
			if err2 != nil {
				return nil, err2
			} else {
				rating := entities.Rating{
					ID:             id,
					DepartmentName: departmentname,
					Rating:         rating,
					DepartmentID:   departmentID,
					// CreatedDate: createddate,
				}
				ratings = append(ratings, rating)
			}
		}
		return ratings, nil
	}
}