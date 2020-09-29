package models

// AssignedClassroom ..

type AssignedClassroom struct {
	Curso    string `json:"curso"`
	Salon    string `json:"salon"`
	Profesor string `json:"profesor"`
	Horario  string `json:"horario"`
}
