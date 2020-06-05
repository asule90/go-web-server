package repository

import (
	"database/sql"
	"go-web-server/domain"
	"go-web-server/utils"
)

type AnimalRepository struct {
	Conn *sql.DB
}

func NewAnimalRepository() *AnimalRepository {
	return &AnimalRepository{utils.DBCon}
}

func (a *AnimalRepository) Get() []domain.Animal {
	result, err := a.Conn.Query("SELECT * FROM animals")
	if err != nil {
		panic(err.Error())
	}
	animal := domain.Animal{}
	res := []domain.Animal{}
	for result.Next() {
		var id uint
		var legsCount int8
		var name, color string

		err = result.Scan(&id, &name, &color, &legsCount)

		if err != nil {
			panic(err.Error())
		}

		animal.Id = id
		animal.Name = name
		animal.Color = color
		animal.LegsCount = legsCount

		res = append(res, animal)
	}

	return res
}

func (a *AnimalRepository) GetById(i int) domain.Animal {
	selDB, err := a.Conn.Query("SELECT * FROM animals WHERE id = ? limit 1", i)
	if err != nil {
		panic(err.Error())
	}
	an := domain.Animal{}

	for selDB.Next() {
		var id uint
		var legsCount int8
		var name, color string

		err = selDB.Scan(&id, &name, &color, &legsCount)
		if err != nil {
			panic(err.Error())
		}
		an.Id = id
		an.Name = name
		an.Color = color
		an.LegsCount = legsCount
	}

	if err != nil {
		panic(err.Error())
	}

	return an
}

func (a *AnimalRepository) Insert(an domain.Animal) (domain.Animal, error) {

	inStmt, err := a.Conn.Prepare("INSERT INTO animals (name, color, legs_count) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	res, err := inStmt.Exec(an.Name, an.Color, an.LegsCount)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	an = a.GetById(int(lastId))

	return an, nil
}

func (a *AnimalRepository) Update(an domain.Animal) (domain.Animal, error) {
	stmt, err := a.Conn.Prepare("UPDATE animals set name = ?, color = ?, legs_count = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	ani := domain.Animal{}

	_, err = stmt.Exec(an.Name, an.Color, an.LegsCount, an.Id)
	if err != nil {
		panic(err.Error())
	}

	ani = a.GetById(int(an.Id))

	return ani, nil
}

func (a *AnimalRepository) Delete(an domain.Animal) error {
	stmt, err := a.Conn.Prepare("DELETE FROM animals WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(an.Id)
	if err != nil {
		panic(err.Error())
	}

	return nil
}
