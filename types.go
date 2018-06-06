package main

// API Entity
type StationView struct {
	Id		int64		`json:"id"`
	Name 		[]string	`json:"name"`
	Color 		[]int		`json:"color"`
	X 		float32		`json:"x"`
	Y 		float32		`json:"y"`
	Active 		[]bool		`json:"active"`
}

type EventView Event
type EventViews []EventView

type ColorView Color
type ColorViews []ColorView

type Configuration struct {
	Url 		string
	User 		string
	Password 	string
}

// ORM
type Station struct {
	Id 		int64		`gorm:"primary_key;auto_increment:false"`
	Part 		int64		`gorm:"primary_key;auto_increment:false"`
	Name 		string
	Color 		int
	X 		float32
	Y 		float32
	Active 		bool
	Description 	string
	PlanTime	int64		`gorm:"type:time"`
	OpenTime	int64		`gorm:"type:time"`
	CloseTime	int64		`gorm:"type:time"`
}

type Color struct {
	Id		int		`gorm:"primary_key" json:"id"`
	Color		int		`json:"color"`
}

type Event struct {
	Id 		int		`gorm:"primary_key" json:"id"`
	Name		string		`json:"name"`
	EventTime	string		`gorm:"type:datetime" json:"time"`
	Description	string		`json:"description"`
}
