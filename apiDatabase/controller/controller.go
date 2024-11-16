package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanmay958/dbapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionstring = "mongodb+srv://studytime958:Tanmay143;@cluster0.9wpc9.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongodb
func init(){
	// client option 
	clientOption :=  options.Client().ApplyURI(connectionstring)
	//connect to mongodb 
	client , err := mongo.Connect(clientOption) 
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection  = client.Database(dbName).Collection(colName) 
	fmt.Println("collection ref is ready")
}

func insertOneMovie(movie model.Netflix){
 inserted, err := collection.InsertOne(context.Background(), movie) 
 if err!=nil {
	log.Fatal(err) 
 }
 fmt.Println("Inserted 1 movie in db with id" ,  inserted.InsertedID) 
}

func updatedOneMovie(movieId string){
	id , _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update :=  bson.M{"$set" : bson.M{"watched":true}}
	result , err := collection.UpdateOne(context.Background() ,  filter,  update) 
	if err!=nil{
		log.Fatal(err) 

	}
	fmt.Println("modified" , result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id ,_ :=  primitive.ObjectIDFromHex(movieId) 
	filter :=  bson.M{"_id" : id}
	deletecount, err := collection.DeleteOne(context.Background(),filter) 
	if err!=nil {
		log.Fatal(err) 
	}
	fmt.Println("movie got deleted" ,  deletecount.DeletedCount) 



}
//
func deleteAll() int64 {
	
	deletecount,  err :=collection.DeleteMany(context.Background(),bson.D{{}}, nil ) 
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println("Number of movie deleted" ,  deletecount.DeletedCount)
	return deletecount.DeletedCount
}

// get all movie from databases
func getAllMovies() []primitive.M{
	cur ,err := collection.Find(context.Background() ,  bson.D{{}})
	if err!= nil{
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err!=nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) 
	}
	defer cur.Close(context.Background()) 
	return movies 


}


func GetMyAllMovies(w http.ResponseWriter ,  r *http.Request){
   w.Header().Set("Content-Type"  , "application/x-www-form-urlencode")
   allmovies := getAllMovies()
   json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter ,  r *http.Request){
	w.Header().Set("Content-Type"  , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST") 
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarskAsWatched(w http.ResponseWriter ,  r *http.Request){
	w.Header().Set("Content-Type"  , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT") 

	params  :=  mux.Vars(r) 
	updatedOneMovie(params["id"]) 
	json.NewEncoder(w).Encode(params["id"])	
}

func DeleteOneMovie(w http.ResponseWriter ,  r *http.Request){
	w.Header().Set("Content-Type"  , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE") 
	params:=mux.Vars(r) 
	 deleteOneMovie(params["id"])
	
	json.NewEncoder(w).Encode("deleted movie")

}
func DeleteAllMovie(w http.ResponseWriter ,  r *http.Request){
	w.Header().Set("Content-Type"  , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE") 
	
	count := deleteAll()
	
	json.NewEncoder(w).Encode(count)

}
