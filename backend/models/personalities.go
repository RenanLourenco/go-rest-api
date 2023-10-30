package models

import "github.com/RenanLourenco/go-rest-api/database"

type Personality struct {
	Id int `json:"id"`
	Name string `json:"name"`
	History string `json:"history"`
}

var Personalities []Personality


func Find() []Personality {
	var p  []Personality
	database.DB.Find(&p)
	return p
}

func FindOne(id string) Personality {
	var p Personality
	err := database.DB.First(&p, id)

	if err.Error != nil {
		panic(err.Error)
	}
	

	return p
}

func Create(p Personality) Personality {
	createdPersonality := p
	database.DB.Create(&createdPersonality)

	return createdPersonality
}

func Delete(id string) Personality {
	var deletedPersonality Personality
	database.DB.Delete(&deletedPersonality,id)

	return deletedPersonality
}


func Edit(id string ,newInfo Personality) Personality {
	var find Personality
	database.DB.First(&find,id)

	find.Name = newInfo.Name
	find.History = newInfo.History

	database.DB.Save(&find)

	return find
}	



