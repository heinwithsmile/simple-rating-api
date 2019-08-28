package models

import (
	"database/sql"
	"entities"
)

//DepartmentModel Struct
type DepartmentModel struct {
	Db *sql.DB
}

//FindAllDepartment Function
func (departmentModel DepartmentModel) FindAllDepartment() (department []entities.Department, err error) {
	rows, err := departmentModel.Db.Query("select * from department")
	if err != nil {
		return nil, err
	} else {
		var departments []entities.Department
		for rows.Next() {
			var id int64
			var name string
			var namemm string
			err2 := rows.Scan(&id, &name,&namemm)
			if err2 != nil {
				return nil, err2
			} else {
				department := entities.Department{
					ID:     id,
					Name : name,
					NameMM : namemm,
				}
				departments = append(departments, department)
			}
		}
		return departments, nil
	}
}