package models

type Ciudad struct {
	ID     		int
	Nombre 		string
	Pais   		string
	Actividad []Actividad
}
