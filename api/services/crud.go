package services

import "github.com/alfatahh54/create-transaction/db"

func Create(data any) (err error) {
	result := db.Database.DB.Create(data)
	err = result.Error
	if err != nil {
		return
	}
	db.Database.DB.Last(data)
	return
}

func UpdateAll(data any) (err error) {
	result := db.Database.DB.Save(data)
	err = result.Error
	return
}

func GetOne(condition, dataResult any) (err error) {
	db.Database.DB.Model(dataResult).Where(condition).First(dataResult)
	return
}

func UpdateModel(model, data any, condition string, values ...any) (err error) {
	result := db.Database.DB.Model(model).Where(condition, values...).Updates(data)
	err = result.Error
	return
}

func GetAll(query string, data any, values ...any) (err error) {
	result := db.Database.DB.Raw(query, values...).Scan(data)
	err = result.Error
	return
}

func GetAllModel(model any) (err error) {
	result := db.Database.DB.Find(model)
	err = result.Error
	return
}

func GetModelByIds(model any, conds ...any) (err error) {
	result := db.Database.DB.Find(model, conds...)
	err = result.Error
	return
}
