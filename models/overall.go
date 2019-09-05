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
	rows, err := allRatingModel.Db.Query("SELECT department.id, department.name_mm, department.name_en, AVG(ratings.rating) average_rating, COUNT(*) 'sum', COUNT( IF(ratings.rating = 0, 1, NULL) ) 'verypoor', COUNT( IF(ratings.rating = 1, 1, NULL) ) 'poor', COUNT( IF(ratings.rating = 2, 1, NULL) ) 'average', COUNT( IF(ratings.rating = 3, 1, NULL) ) 'good', COUNT( IF(ratings.rating = 4, 1, NULL) ) 'verygood',CONCAT(ROUND((COUNT(IF(ratings.rating = 0, 1, NULL)) / COUNT(*) * 100),2),'%') AS vppercent,CONCAT(ROUND((COUNT(IF(ratings.rating = 1, 1, NULL)) / COUNT(*) * 100),2),'%') AS ppercent,CONCAT(ROUND((COUNT(IF(ratings.rating = 2, 1, NULL)) / COUNT(*) * 100),2),'%') AS avgpercent,CONCAT(ROUND((COUNT(IF(ratings.rating = 3, 1, NULL)) / COUNT(*) * 100),2),'%') AS gdpercent,CONCAT(ROUND((COUNT(IF(ratings.rating = 4, 1, NULL)) / COUNT(*) * 100),2),'%') AS vgpercent FROM ratings INNER JOIN department ON ratings.departmentId = department.id GROUP BY `departmentId` ORDER BY id")
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
			var vppercent string
			var ppercent string
			var avgpercent string
			var gpercent string
			var vgpercent string
			err2 := rows.Scan(&id, &namemm, &nameen, &avgrating, &sum, &verypoor, &poor, &average, &good, &verygood, &vppercent, &ppercent, &avgpercent, &gpercent, &vgpercent)
			if err2 != nil {
				return nil, err2
			} else {
				data := entities.OverALLData{
					ID:         id,
					NameMM:     namemm,
					NameEN:     nameen,
					AvgRating:  avgrating,
					Sum:        sum,
					VeryPoor:   verypoor,
					Poor:       poor,
					Average:    average,
					Good:       good,
					VeryGood:   verygood,
					Vppercent:  vppercent,
					Ppercent:   ppercent,
					Avgpercent: avgpercent,
					Gpercent:   gpercent,
					Vgpercent:  vgpercent,
				}
				alldata = append(alldata, data)
			}
		}
		return alldata, nil
	}
}

//WeeklyData Method
func (allRatingModel OverallDataModel) WeeklyData() (alldata []entities.OverALLData, err error) {
	rows, err := allRatingModel.Db.Query("SELECT department.id, department.name_mm, department.name_en, AVG(ratings.rating) average_rating, COUNT(*) 'sum', COUNT( IF(ratings.rating = 0, 1, NULL) ) 'verypoor', COUNT( IF(ratings.rating = 1, 1, NULL) ) 'poor', COUNT( IF(ratings.rating = 2, 1, NULL) ) 'average', COUNT( IF(ratings.rating = 3, 1, NULL) ) 'good', COUNT( IF(ratings.rating = 4, 1, NULL) ) 'verygood' FROM ratings INNER JOIN department ON ratings.departmentId = department.id WHERE `createdOn`>= DATE_SUB(CURDATE(), INTERVAL 10 DAY) GROUP BY `departmentId` ORDER BY id")
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
