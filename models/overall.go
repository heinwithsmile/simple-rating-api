package models

import (
	"database/sql"
	"entities"
)

//OverallDataModel Struct
type OverallDataModel struct {
	Db *sql.DB
}

//AllData Method
func (allRatingModel OverallDataModel) AllData() (alldata []entities.OverALLData, err error) {
	rows, err := allRatingModel.Db.Query("SELECT department.id,department.name_mm,department.name_en,AVG(ratings.rating) average_rating,COUNT(*) 'sum',COUNT(IF(ratings.rating = 0 ,1, NULL)) 'verypoor',COUNT(IF(ratings.rating = 1 ,1, NULL)) 'poor',COUNT(IF(ratings.rating = 2 ,1, NULL)) 'average',COUNT(IF(ratings.rating = 3 ,1, NULL)) 'good',COUNT(IF(ratings.rating = 4 ,1, NULL)) 'verygood' FROM ratings INNER JOIN department ON ratings.departmentId = department.id GROUP BY departmentId ORDER BY id")
	if err != nil {
		return nil, err
	} else {
		var alldata []entities.OverALLData
		for rows.Next() {
			var id int64
			var namemm string
			var nameen string
			var avgrating float64
			var sum int64
			var verypoor int64
			var poor int64
			var average int64
			var good int64
			var verygood int64
			err2 := rows.Scan(&id, &namemm, &nameen, &avgrating, &sum, &verypoor, &poor, &average, &good, &verygood)
			if err2 != nil {
				return nil, err2
			} else {
				data := entities.OverALLData{
					ID:        id,
					NameMM:    namemm,
					NameEN:    nameen,
					AvgRating: avgrating,
					Sum:       sum,
					VeryPoor:  verypoor,
					Poor:      poor,
					Average:   average,
					Good:      good,
					VeryGood:  verygood,
				}
				alldata = append(alldata, data)
			}
		}
		return alldata, nil
	}
}

//DailyData Method
func (allRatingModel OverallDataModel) DailyData() (alldata []entities.OverALLData, err error) {
	rows, err := allRatingModel.Db.Query("SELECT department.id, department.name_mm, department.name_en, AVG(ratings.rating) average_rating, COUNT(*) 'sum', COUNT( IF(ratings.rating = 0, 1, NULL) ) 'verypoor', COUNT( IF(ratings.rating = 1, 1, NULL) ) 'poor', COUNT( IF(ratings.rating = 2, 1, NULL) ) 'average', COUNT( IF(ratings.rating = 3, 1, NULL) ) 'good', COUNT( IF(ratings.rating = 4, 1, NULL) ) 'verygood' FROM ratings INNER JOIN department ON ratings.departmentId = department.id WHERE `createdOn` > DATE_ADD(DATE(NOW()), INTERVAL - 1 DAY) AND `createdOn` < DATE(NOW()) GROUP BY `departmentId`, DATE(NOW()) ORDER BY id")
	if err != nil {
		return nil, err
	} else {
		var alldata []entities.OverALLData
		for rows.Next() {
			var id int64
			var namemm string
			var nameen string
			var avgrating float64
			var sum int64
			var verypoor int64
			var poor int64
			var average int64
			var good int64
			var verygood int64
			err2 := rows.Scan(&id, &namemm, &nameen, &avgrating, &sum, &verypoor, &poor, &average, &good, &verygood)
			if err2 != nil {
				return nil, err2
			} else {
				data := entities.OverALLData{
					ID:        id,
					NameMM:    namemm,
					NameEN:    nameen,
					AvgRating: avgrating,
					Sum:       sum,
					VeryPoor:  verypoor,
					Poor:      poor,
					Average:   average,
					Good:      good,
					VeryGood:  verygood,
				}
				alldata = append(alldata, data)
			}
		}
		return alldata, nil
	}
}

//MonthData Method
func (allRatingModel OverallDataModel) MonthData() (alldata []entities.OverALLData, err error) {
	rows, err := allRatingModel.Db.Query("SELECT department.id, department.name_mm, department.name_en, AVG(ratings.rating) average_rating, COUNT(*) 'sum', COUNT( IF(ratings.rating = 0, 1, NULL) ) 'verypoor', COUNT( IF(ratings.rating = 1, 1, NULL) ) 'poor', COUNT( IF(ratings.rating = 2, 1, NULL) ) 'average', COUNT( IF(ratings.rating = 3, 1, NULL) ) 'good', COUNT( IF(ratings.rating = 4, 1, NULL) ) 'verygood' FROM ratings INNER JOIN department ON ratings.departmentId = department.id WHERE `createdOn` > DATE_ADD(DATE(NOW()), INTERVAL - 1 MONTH) AND `createdOn` < DATE(NOW()) GROUP BY `departmentId`, DATE(NOW()) ORDER BY id")
	if err != nil {
		return nil, err
	} else {
		var alldata []entities.OverALLData
		for rows.Next() {
			var id int64
			var namemm string
			var nameen string
			var avgrating float64
			var sum int64
			var verypoor int64
			var poor int64
			var average int64
			var good int64
			var verygood int64
			err2 := rows.Scan(&id, &namemm, &nameen, &avgrating, &sum, &verypoor, &poor, &average, &good, &verygood)
			if err2 != nil {
				return nil, err2
			} else {
				data := entities.OverALLData{
					ID:        id,
					NameMM:    namemm,
					NameEN:    nameen,
					AvgRating: avgrating,
					Sum:       sum,
					VeryPoor:  verypoor,
					Poor:      poor,
					Average:   average,
					Good:      good,
					VeryGood:  verygood,
				}
				alldata = append(alldata, data)
			}
		}
		return alldata, nil
	}
}

//YearData Method
func (allRatingModel OverallDataModel) YearData() (alldata []entities.OverALLData, err error) {
	rows, err := allRatingModel.Db.Query("SELECT department.id, department.name_mm, department.name_en, AVG(ratings.rating) average_rating, COUNT(*) 'sum', COUNT( IF(ratings.rating = 0, 1, NULL) ) 'verypoor', COUNT( IF(ratings.rating = 1, 1, NULL) ) 'poor', COUNT( IF(ratings.rating = 2, 1, NULL) ) 'average', COUNT( IF(ratings.rating = 3, 1, NULL) ) 'good', COUNT( IF(ratings.rating = 4, 1, NULL) ) 'verygood' FROM ratings INNER JOIN department ON ratings.departmentId = department.id WHERE `createdOn` > DATE_ADD(DATE(NOW()), INTERVAL - 1 YEAR) AND `createdOn` < DATE(NOW()) GROUP BY `departmentId`, DATE(NOW()) ORDER BY id")
	if err != nil {
		return nil, err
	} else {
		var alldata []entities.OverALLData
		for rows.Next() {
			var id int64
			var namemm string
			var nameen string
			var avgrating float64
			var sum int64
			var verypoor int64
			var poor int64
			var average int64
			var good int64
			var verygood int64
			err2 := rows.Scan(&id, &namemm, &nameen, &avgrating, &sum, &verypoor, &poor, &average, &good, &verygood)
			if err2 != nil {
				return nil, err2
			} else {
				data := entities.OverALLData{
					ID:        id,
					NameMM:    namemm,
					NameEN:    nameen,
					AvgRating: avgrating,
					Sum:       sum,
					VeryPoor:  verypoor,
					Poor:      poor,
					Average:   average,
					Good:      good,
					VeryGood:  verygood,
				}
				alldata = append(alldata, data)
			}
		}
		return alldata, nil
	}
}

