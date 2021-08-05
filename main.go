package main

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strings"
	"fmt"
	"log"
	"todo/todo"
)

func main() {
	r := mux.NewRouter()

	
	r.HandleFunc("/auth",func(rw http.ResponseWriter,r *http.Request){
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt : time.Now().Add(2*time.Minute).Unix(),
			Issuer: "test",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})
	})

	api := r.NewRoute().Subrouter()
	api.Use(authMiddleware)

	api.HandleFunc("/todos", todo.AddTask).Methods(http.MethodPut)

	api.HandleFunc("/todos/{index}", todo.ChangeDoneTask).Methods(http.MethodPut)
	api.HandleFunc("/todos", todo.GetTask).Methods(http.MethodGet)

	http.ListenAndServe(":9090", r)

	// r.HandleFunc("/todos",func(rw http.ResponseWriter,r *http.Request){
	// 	tokenString := r.Header.Get("Authorization")
	// 	tokenString = strings.ReplaceAll(tokenString,"Bearer ","")
	// 	mySigningKey := []byte("password")
	// 	_,err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
	// 		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, fmt.Errorf("Unexpected signing method: %v",token.Header["alg"])
	// 		}
	// 		return mySigningKey,nil
	// 	})
	// 	if err != nil {
	// 		rw.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	defer r.Body.Close()
	// 	var task NewTaskTodo
	// 	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
	// 		rw.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	New(task.Task)
	// }).Methods(http.MethodPut)

	
	// r.HandleFunc("/todos/{index}",func(rw http.ResponseWriter,r *http.Request){
	// 	vars := mux.Vars(r)
	// 	index := vars["index"]
	// 	i,err := strconv.Atoi(index)
	// 	if err != nil {
	// 		rw.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	tasks[i].Done = true
	// }).Methods(http.MethodPut)


	// r.HandleFunc("/todos",func(rw http.ResponseWriter,r *http.Request){
	// 	// json.NewEncoder(rw).Encode(tasks)
		
	// 	if err := json.NewEncoder(rw).Encode(List()); err != nil {
	// 		rw.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// }).Methods(http.MethodGet)
	// http.ListenAndServe(":9090",r)
	// New("task1")
	// New("task2")
	// New("task3")
	// for k,v := range List() {
	// 	fmt.Println(k,v)
	// }
}

//authMiddleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte("password")
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		})
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println(r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}

//loggingMiddleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}