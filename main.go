package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID       int        `json:"id"`
	Name     string     `json:"name`
	Teacher  *Teacher   `json:"teacher"`
	Students []*Student `json:"students"`
	Rating   float64    `json:"rating"`
	ETCS     float64    `json:"etcs`
}

type Teacher struct {
	Name  string
	ID    int
	Score float64
}

type Student struct {
	Name       string
	ID         int
	Age        int
	Enrollment bool
}

var Courses = []Course{
	{
		ID: 0, Name: "DISYS", Teacher: &Teacher{Name: "Bob Bowman", ID: 0, Score: 9.9}, Students: []*Student{&Student{Name: "Aaron", ID: 0, Age: 21, Enrollment: true}, &Student{Name: "Biggie", ID: 1, Age: 22, Enrollment: true}}, Rating: 3.0, ETCS: 7.5,
	},
	{
		ID: 1, Name: "BDSA", Teacher: &Teacher{Name: "Steve Loose", ID: 1, Score: 2.1}, Students: []*Student{&Student{Name: "Carl", ID: 2, Age: 19, Enrollment: true}, &Student{Name: "Louise", ID: 3, Age: 21, Enrollment: true}}, Rating: 7.0, ETCS: 15.0,
	},
	{
		ID: 2, Name: "IDBS", Teacher: &Teacher{Name: "Petra Ponytail", ID: 2, Score: 5.0}, Students: []*Student{&Student{Name: "Niklas", ID: 4, Age: 20, Enrollment: true}, &Student{Name: "Josephine", ID: 5, Age: 21, Enrollment: true}}, Rating: 4.1, ETCS: 7.5,
	},
}

func main() {
	router := gin.Default()
	router.GET("/Courses", getCoursesByID)

	router.Run("localhost:8080")
}

func getCoursesByID(g *gin.Context) {
	fmt.Println("Enter course ID: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	temp := strings.TrimSpace(line)
	fmt.Println(temp)
	number, _ := strconv.Atoi(temp)
	fmt.Println(number)

	for _, Course := range Courses {
		fmt.Println(Course.ID, number)
		if Course.ID == number {
			g.IndentedJSON(http.StatusOK, Course)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}
