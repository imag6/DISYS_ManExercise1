package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID      int     `json:"id"`
	Name    string  `json:"name`
	Teacher string  `json:"teacher"`
	Rating  float64 `json:"rating"`
	ETCS    float64 `json:"etcs`
}

var Courses = []Course{
	{
		ID: 1, Name: "DISYS", Teacher: "Bob Bowman", Rating: 3.0, ETCS: 7.5,
	},
	{
		ID: 2, Name: "BDSA", Teacher: "Biggie Big", Rating: 7.0, ETCS: 15.0,
	},
	{
		ID: 3, Name: "IDBS", Teacher: "Petra Ponytail", Rating: 4.1, ETCS: 7.5,
	},
}

func main() {
	router := gin.Default()
	router.GET("/Course", getCourses)
	router.GET("/Course/:id", getCourseByID)
	router.POST("/Course", postCourse)
	router.PUT("/Course/:id", putCourseByID)
	router.DELETE("Course/:id", deleteCourseById)

	router.Run("localhost:8080")
}

func postCourse(g *gin.Context) {
	var newCourse Course

	if err := g.BindJSON(&newCourse); err != nil {
		return
	}

	Courses = append(Courses, newCourse)
	g.IndentedJSON(http.StatusCreated, newCourse)
}

func deleteCourseById(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))

	for i, Course := range Courses {
		if Course.ID == id {
			Courses[i] = Courses[len(Courses)-1]
			Courses = Courses[:len(Courses)-1]
			g.IndentedJSON(http.StatusOK, Courses)
			return
		}
	}
	g.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Course ID not found"})
}

func putCourseByID(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))

	var updatedCourse Course

	if err := g.BindJSON(&updatedCourse); err != nil {
		return
	}

	for i, Course := range Courses {
		if Course.ID == id {
			Courses[i] = updatedCourse
		}
	}
	g.IndentedJSON(http.StatusOK, Courses)
}
func getCourses(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, Courses)
}

func getCourseByID(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))

	for _, Course := range Courses {
		if Course.ID == id {
			g.IndentedJSON(http.StatusOK, Course)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course id not found"})
}
