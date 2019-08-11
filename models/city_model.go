package models

import (
	"database/sql"
	"entities"
)
//CityModel Struct
type CityModel struct {
	Db *sql.DB
}
//FindAll Function
func (cityModel CityModel) FindAll() (city []entities.City,err error){
	rows, err := cityModel.Db.Query("select * from city")
	if err != nil {
		return nil, err
	}else{
		var cities []entities.City
		for rows.Next() {
			var id int64
			var name string
			err2 := rows.Scan(&id,&name)
			if err2 != nil{
				return nil , err2
			}else{
				city := entities.City {
					ID : id,
					Name : name,
				}
				cities = append(cities,city)
			}
		}
		return cities , nil
	}
}

//Search Method
func (cityModel CityModel) Search(keyword string) (product []entities.City,err error){
	rows, err := cityModel.Db.Query("select * from city where name like ?","%"+keyword+"%")
	if err != nil {
		return nil, err
	}else{
		var cities []entities.City
		for rows.Next() {
			var id int64
			var name string
			err2 := rows.Scan(&id,&name)
			if err2 != nil{
				return nil , err2
			}else{
				city := entities.City{
					ID : id,
					Name : name,
				}
				cities = append(cities,city)
			}
		}
		return cities , nil
	}
}

//Search Prices Function
// func (productModel ProductModel) SearchPrices(min,max float64) (product []entities.Product,err error){
// 	rows, err := productModel.Db.Query("select * from product where price >= ? and price <= ? ", min , max)
// 	if err != nil {
// 		return nil, err
// 	}else{
// 		var products []entities.Product
// 		for rows.Next() {
// 			var id int64
// 			var name string
// 			var price float64
// 			var quantity int64
// 			err2 := rows.Scan(&id,&name,&price,&quantity)
// 			if err2 != nil{
// 				return nil , err2
// 			}else{
// 				product := entities.Product{
// 					ID : id,
// 					Name : name,
// 				}
// 				products = append(products,product)
// 			}
// 		}
// 		return products , nil
// 	}
// }

//Create Method
func (cityModel CityModel) Create(city *entities.City) (err error) {
	result, err := cityModel.Db.Exec("INSERT INTO city (name) VALUES (?)", city.Name)
	if err != nil {
		return err
	} else {
		city.ID, _ = result.LastInsertId()
		return nil
	}
}
// // //Update Method
// // func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
// // 	result, err := productModel.Db.Exec("UPDATE product SET name=?, price=?, quantity=? WHERE id=?", product.Name,product.Price,product.Quantity,product.ID)
// // 	if err != nil {
// // 		return 0 , err
// // 	} else {
// // 		return result.RowsAffected()
// // 	}
// // }

//Delete Method
func (cityModel CityModel) Delete(id int64) (int64 , error) {
	result, err := cityModel.Db.Exec("DELETE FROM city WHERE id=?",id)
	if err != nil {
		return 0 , err
	} else {
		return result.RowsAffected()
	}
}

