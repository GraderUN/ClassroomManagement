package models

// Courses ...

type Courses struct {
	Grado         int8   `json:"grado"`
	Letra         string `json:"letra"`
	Id_estudiante []int  `json:"id_estudiante"`
}
