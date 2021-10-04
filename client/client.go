package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var basePath = "http://localhost:8080/"

func main() {
	initClient()
	isExited := false

	for !isExited {
		fmt.Println("Please type in your command: ")
		reader := bufio.NewReader(os.Stdin)
		request, _ := reader.ReadString('\n')
		requestSplit := strings.Split(strings.TrimSpace(request), " ")
		switch requestSplit[0] {
		case "get":
			httpGet()
		case "getID":
			httpGetID(requestSplit)
		case "put":
			httpPut(requestSplit)
		case "post":
			httpPost(requestSplit)
		case "delete":
			httpDelete(requestSplit)
		case "help":
			httpHelp()
		case "q":
			fmt.Println("Client is closed")
			isExited = true
		default:
			fmt.Println()
			fmt.Println("Please use the commands given by the client. If you are uncertain about those, use the help command")
			continue
		}
	}
}

func httpGet() {
	//request data via http request
	response, _ := http.Get(basePath + "Course")

	//close response when method is done
	defer response.Body.Close()

	// read data from request
	data, _ := io.ReadAll(response.Body)

	fmt.Println(string(data))
	fmt.Println()
}

func httpGetID(command []string) {
	response, _ := http.Get(basePath + "Course/" + string(command[1]))

	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	fmt.Println(string(data))
	fmt.Println()

}

func httpPost(command []string) {
	body := &Course{
		ID:      stringToInt(command[1]),
		Name:    command[2],
		Teacher: command[3],
		Rating:  stringToFloat(command[4]),
		ETCS:    stringToFloat(command[5]),
	}

	//marshal json data
	jsondata, _ := json.Marshal(body)

	response, _ := http.Post(basePath+"Course", "application/json", bytes.NewBuffer(jsondata))
	defer response.Body.Close()

	//data, _ := io.ReadAll(response.Body)

	//fmt.Println(string(data))
	httpGet()
	fmt.Println()
}

func httpPut(command []string) {
	// request body
	body := &Course{
		ID:      stringToInt(command[2]),
		Name:    command[3],
		Teacher: command[4],
		Rating:  stringToFloat(command[5]),
		ETCS:    stringToFloat(command[6]),
	}

	//sets up client for http request
	client := &http.Client{}

	//marshall request body
	jsondata, _ := json.Marshal(body)

	//set http method, url and the request body
	request, _ := http.NewRequest(http.MethodPut, basePath+"Course/"+command[1], bytes.NewBuffer(jsondata))
	//set the reques header content - tells it is a json base request body i guess
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	//make the cliet do the request
	response, _ := client.Do(request)
	defer response.Body.Close()
	httpGet()
	fmt.Println()
}

func httpDelete(command []string) {

	client := &http.Client{}

	request, _ := http.NewRequest(http.MethodDelete, basePath+"Course/"+command[1], nil)

	response, _ := client.Do(request)
	defer response.Body.Close()

	httpGet()
	fmt.Println()
}

func httpHelp() {
	fmt.Println()
	fmt.Println("-------------------------------------------------------")
	fmt.Println("The client responds to following commands:")
	fmt.Println("get 									(Returns a list over all courses)")
	fmt.Println("getID <ID> 								(Returns the course with the given id)")
	fmt.Println("put <oldID> <newID> <name> <teacher> <rating> <etcs> 			(Updates the course with the given id)")
	fmt.Println("post <ID> <name> <teacher> <rating> <etcs>				(Creates a course with the given values)")
	fmt.Println("delete <ID> 								(Deletes the course with the given id)")
	fmt.Println("help 									(Shows the list of commands)")
	fmt.Println("q									(Closes the client)")
	fmt.Println("-------------------------------------------------------")
	fmt.Println()
}

func initClient() {
	fmt.Println("---- The ITU API Client is now booted up! Welcome ----")
	fmt.Println("------------------------------------------------------")
	httpHelp()
}

func stringToInt(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

func stringToFloat(s string) float64 {
	number, _ := strconv.ParseFloat(s, 64)
	return number
}

type Course struct {
	ID      int     `json:"id"`
	Name    string  `json:"name`
	Teacher string  `json:"teacher"`
	Rating  float64 `json:"rating"`
	ETCS    float64 `json:"etcs`
}
