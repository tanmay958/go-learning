package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MODEL FOR COURSE -  file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}
type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}
var courses[] Course 

// middleware ,helper  -file 
func (c*Course )IsEmpty()  bool{
	return  c.CourseName=="" 
}
func main()  {
	fmt.Println("API-LearnCode ONLINE")
	r:= mux.NewRouter()
	courses =  append(courses , Course{CourseId : "2" ,  CourseName : "Reactjs" ,CoursePrice: 299 ,Author:  &Author{Fullname: "Tanmay Mandal" , Website:  "tanmay.com"} })
	courses =  append(courses , Course{CourseId : "3" ,  CourseName : "Atomjs" ,CoursePrice: 299 ,Author:  &Author{Fullname: "Tanmay Mandal" , Website:  "tanmay2.com"} })
	// routing  
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET") 
	r.HandleFunc("/course/{id}" ,getOneCourse).Methods("GET") 
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}" ,deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000",r))
}

// controllers - file 

//serve home route
func serveHome(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("<h1>hye Tanmay</h1>"))
}
//get all course 
func getAllCourses(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get all courses") 
	w.Header().Set("Contetn-Type" , "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get one course") 
	w.Header().Set("content-Type" , "application/json") 
	// grab id 
	params := mux.Vars(r) 
	fmt.Println(params) 
	
	// loop through courses 
	for  _,course := range courses {
		if  course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)  
			return 
		}
	}
	json.NewEncoder(w).Encode("No course found with give id" ) 
}

func createOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("create course") 
	w.Header().Set("Content-Type",  "application/json") 

	//what if body is empty 
	if r.Body== nil {
		json.NewEncoder(w).Encode("please send some data") 
	}

	// what about   -{}
	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("please send some data") 
		return  
	}
	//geneate unique id append 
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) 
	courses =  append(courses ,  course) 
	json.NewEncoder(w).Encode(course)

}


func updateOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Update ONE COURSE")
	w.Header().Set("Content-Type" , "application/json")
	//first - grab id from request 
	params := mux.Vars(r) 
	// loop through 
	for index,course :=  range courses {
		if course.CourseId == params["id"] {
			courses =  append(courses[:index] , courses[index+1 :]...)
			var course Course
 			_= json.NewDecoder(r.Body).Decode(&course) 
			course.CourseId = params["id"] 
			courses =  append(courses,  course) 
			json.NewEncoder(w).Encode(course)
			return  
		}
	}
	json.NewEncoder(w).Encode("no such course found")
}
func deleteOneCourse(w http.ResponseWriter ,  r* http.Request)  {
	fmt.Println("delete course") 
	w.Header().Set("Content-Type","application/json" ) 
	params:= mux.Vars(r) 
	for index, course :=  range courses {
		if(course.CourseId ==  params["id"]){
			courses = append(courses[:index], courses[index+1 :]...)  
			json.NewEncoder(w).Encode(course) 
			return  

		}
	}
	json.NewEncoder(w).Encode("not able to find course")



}