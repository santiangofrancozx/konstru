package models

type Proyectos_actividades struct {
	ID_proyecto  string
	ID_actividad string
	Proyectos    Proyectos
	Actividad    Actividad
}
