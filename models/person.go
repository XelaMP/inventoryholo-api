package models

import "fmt"

type Person struct {
	IdPerson int    `json:"id_person"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Phone    int    `json:"phone"`
	Adress   string `json:"adress"`
	Dni      int    `json:"dni"`
	Mail     string `json:"mail"`
}

func (person Person) ToString() string {
	return fmt.Sprintf("id person: %d\nname: %s\nlastname: %s\nphone: %d\nadress: %d\ndni: %d\nmail: %s",
		person.IdPerson, person.Name, person.LastName, person.Phone, person.Adress, person.Dni, person.Mail)

}
