package main

import (
	"fmt"
	"github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/XelaMP/inventoryholo-api/helper"
)

func main ()  {
	Demo1()
	

}

func Demo1()  {
	db, err := helper.GetDB()
	if err != nil {
		fmt.Println(err)
		
	} else {
		personc := controllers.PatientController{
			Db: db,

		}
		persons, err2 := personc.FindAll()
		if err2 != nil {
			fmt.Println(err2)
			
		}else{
			for _, person := range persons{
				fmt.Println(person.ToString())
				fmt.Println("----------------")
			}
		}

	}
	
}
