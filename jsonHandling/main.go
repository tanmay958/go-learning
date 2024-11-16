package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //aliases
	Price    int
	Platform string
	Password string  `json:"-"` //  excluded 
	Tags     []string  `json:"tags,omitempty"`
}

func main() {
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs", 299, "tanmay.com", "1234", []string{"web", "dev"}},
		{"AngulatJs", 499, "tanmay.com", "1234", []string{"web", "dev"}},
		{"Python", 299, "tanmay.com", "12e34", nil},
	}

	// package this data
	finalJson,err  := json.MarshalIndent(courses,"","\t")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n" ,  finalJson) 
}

func DecodeJson()  {
	jsonDataFromWeb := []byte(` {
                "coursename": "AngulatJs",
                "Price": 499,
                "Platform": "tanmay.com",
                "tags": ["web","dev"]
        }`)
	var tanmaycourse course 
	checValid := json.Valid(jsonDataFromWeb) 
	if checValid {
		fmt.Println("JSON IS VALID")
		json.Unmarshal(jsonDataFromWeb,&tanmaycourse)
		fmt.Printf("%#v\n",  tanmaycourse) 
	}else{
		fmt.Printf("Json not valid")
	}
   // some key value pair
   var myOnlineData  map[string]interface{}
   json.Unmarshal(jsonDataFromWeb,  &myOnlineData) 
   for key,value := range myOnlineData {
		fmt.Printf("the key is %v value is %v and type of %T\n" ,  key ,  value,value) 
   }
     
	
}