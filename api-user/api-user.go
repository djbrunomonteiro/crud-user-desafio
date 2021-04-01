package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nome  string             `json:"nome,omitempty" bson:"nome,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Senha string             `json:"senha,omitempty" bson:"senha,omitempty"`
}

var client *mongo.Client

func main() {

	fmt.Println("Iniciando..")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.TODO(), clientOptions)

	fmt.Println("Conectado ao mongo db ")
	Rotas()
}

func Rotas() {
	//Rotas e metodos
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router)) //Ativa o cors no localhost

	router.HandleFunc("/usuarios", CreateUser).Methods("POST", "OPTION")
	router.HandleFunc("/usuarios", GetUser).Methods("GET", "OPTION")
	router.HandleFunc("/usuarios/delete/{id}", GetUserId).Methods("GET", "OPTION")
	router.HandleFunc("/usuarios/update/{id}", UpdateUser).Methods("PUT", "OPTION")
	router.HandleFunc("/usuarios/delete/{id}", DeleteUser).Methods("DELETE", "OPTION")

	http.ListenAndServe(":7000", router)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Methods", "GET")
	response.Header().Add("Access-Control-Allow-Methods", "OPTION")
	response.Header().Add("content-type", "application/json")

	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("users").Collection("user")
	result, _ := collection.InsertOne(context.TODO(), user)
	json.NewEncoder(response).Encode(result)
	println("user criado")

}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Methods", "GET")
	response.Header().Add("Access-Control-Allow-Methods", "OPTION")
	response.Header().Add("content-type", "application/json")

	var results []User
	collection := client.Database("users").Collection("user")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err, "Erro ao listar users")
		return
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user User
		cur.Decode(&user)
		results = append(results, user)
	}
	if err := cur.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(results)
	println("Usuarios Listados")

}

func GetUserId(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Methods", "GET")
	response.Header().Add("Access-Control-Allow-Methods", "OPTION")
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	_id, _ := primitive.ObjectIDFromHex(params["_id"])
	var user User
	collection := client.Database("users").Collection("user")
	err := collection.FindOne(context.TODO(), User{ID: _id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err, "Erro em listar User ID")
		return
	}
	json.NewEncoder(response).Encode(user)
	println("Usuario Listado")
}
func UpdateUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Methods", "PUT")
	response.Header().Add("Access-Control-Allow-Methods", "OPTION")
	response.Header().Add("content-type", "application/json")

	collection := client.Database("users").Collection("user")

	// Pega o id dos parâmetros
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user User

	// Criar filtro
	filter := bson.M{
		"_id": id,
	}

	//Leia o modelo de atualização da solicitação do corpo
	_ = json.NewDecoder(request.Body).Decode(&user)

	// prepara o modelo para atualizar.
	update := bson.M{
		"$set": bson.M{
			"nome":  user.Nome,
			"email": user.Email,
			"senha": user.Senha,
		},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(user)
	println("Usuarios atualizado")

}
func DeleteUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Methods", "PUT")
	response.Header().Add("Access-Control-Allow-Methods", "OPTION")
	response.Header().Add("content-type", "application/json")

	collection := client.Database("users").Collection("user")

	params := mux.Vars(request)
	_id, _ := primitive.ObjectIDFromHex(params["_id"])

	var user User

	filter := bson.M{
		"_id": _id,
	}

	_ = json.NewDecoder(request.Body).Decode(&user)

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(user)
	println("Usuario Deletado!")

}
