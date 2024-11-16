package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to servies")
    PerfromFormRequest() 
}
func PerformGetRequest()  {
    const myurl = "http://localhost:8000/get"
    response,err :=  http.Get(myurl)
    if err!= nil{
        panic(err)
    }
    defer response.Body.Close() 
    fmt.Println("Status Code" , response.StatusCode)
    fmt.Println("Content length" ,  response.ContentLength) 

    var responseString strings.Builder
    content,_ := ioutil.ReadAll(response.Body) 
    bytecount,_ := responseString.Write(content)
    fmt.Println("ByteCount" , bytecount) 
    fmt.Println("string is" , responseString.String()) 

    fmt.Println(string(content))
}
func PerfromPostRequest(){
    const myurl = "http://localhost:8000/post"
    requestBody :=  strings.NewReader(`
        {
          "coursename":"python",
          "price":"0",
          "platform":"nitt.edu"  
        }
    `)
    response,err :=http.Post(myurl, "application/json",requestBody )
    if err!= nil{
        panic(err) 
    }

   defer response.Body.Close()
    content,_:= ioutil.ReadAll(response.Body)
    fmt.Println(string(content))
}
func PerfromFormRequest(){
    const myurl = "http://localhost:8000/postform"

    // form data
    data :=  url.Values{}
    data.Add("firstname" , "tanmay") 
    data.Add("lastname","mandal")
    response,err := http.PostForm(myurl,  data)
    if err!=nil {
        panic(err)
    }
    defer response.Body.Close()

    content,_ := ioutil.ReadAll(response.Body)

    fmt.Println(string(content)) 



}