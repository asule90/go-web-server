package service

import (
	"encoding/json"
	"go-web-server/domain"
	"go-web-server/repository"
	"log"
)

type AnimalService struct {
	Repo *repository.AnimalRepository
}

func NewAnimalService() *AnimalService {
	return &AnimalService{repository.NewAnimalRepository()}
}

func (s *AnimalService) Get() []domain.Animal {

	animals := s.Repo.Get()

	return animals
}

func (s *AnimalService) GetOne(id int) domain.Animal {

	animal := s.Repo.GetById(id)

	return animal
}

func (s *AnimalService) Create(a domain.Animal) (domain.Animal, error) {

	animal, err := s.Repo.Insert(a)
	if err != nil {
		panic(err.Error())
	}

	res2B, _ := json.Marshal(animal)
	log.Printf("%s", res2B)

	return animal, nil
}

func (s *AnimalService) Update(id int, newA domain.Animal) (domain.Animal, error) {

	animal := s.Repo.GetById(id)
	animal.Name = newA.Name
	animal.Color = newA.Color
	animal.LegsCount = newA.LegsCount

	ani, err := s.Repo.Update(animal)
	if err != nil {
		panic(err.Error())
	}

	return ani, nil
}

func (s *AnimalService) Delete(id int) error {

	animal := s.Repo.GetById(id)

	err := s.Repo.Delete(animal)
	if err != nil {
		panic(err.Error())
	}

	return nil
}
