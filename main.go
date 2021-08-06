package main

import(
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strings"
	"log"
	"fmt"
	"todo/todo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db , err := gorm.Open(sqlite.Open("gorm.db"),&gorm.Config{})
	if err!=nil {
		log.Panic(err)
	}
	db.AutoMigrate(&Task{})
	
	app := todo.NewApp(db)
	r.GET("/auth",func(c *gin.Context) {
		
		mySigningKey := []byte("password")

		// Create the Claims
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, map[string]string{
			"token": ss,
		})
	})

	api := r.Group("")
	api.Use(authMiddleware)
	// todoApp := todo.NewApp(serialize.NewJSONSerializer())
	r.PUT("/todos",app.AddTask)
	r.PUT("/todos/:id",todo.ChangeDoneTask)
	r.GET("/todos",todo.GetTask)
	r.Run(":9090")

}

//authMiddleware
func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	mySigningKey := []byte("password")

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Next()
}

// //authMiddleware
// func authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		mySigningKey := []byte("password")
// 		tokenString := r.Header.Get("Authorization")
// 		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

// 		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Don't forget to validate the alg is what you expect:
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}

// 			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 			return mySigningKey, nil
// 		})
// 		if err != nil {
// 			rw.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		log.Println(r.RequestURI)
// 		next.ServeHTTP(rw, r)
// 	})
// }

//loggingMiddleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}