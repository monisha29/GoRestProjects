package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/monisha29/app1/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	//r.GET("/", index)
	r.POST("/user", uc.CreateUser)
	r.GET("/user/:id", uc.GetUser)
	http.ListenAndServe("localhost:8080", r)
	// added route plus parameter

}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

// note: using 'r' instead of 'req'
// func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }

// func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	s := `<!DOCTYPE html>
// <html lang="en">
// <head>
// <meta charset="UTF-8">
// <title>Index</title>
// </head>
// <body>
// <a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
// </body>
// </html>
// 	`
// 	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(s))
// }

// // changed func name
// func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	u := models.User{
// 		Name:   "James Bond",
// 		Gender: "male",
// 		Age:    32,
// 		Id:     p.ByName("id"),
// 	}

// 	fmt.Println("Printing :: ", u.Id)
// 	// Marshal into JSON
// 	uj, err := json.Marshal(u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Write content-type, statuscode, payload
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK) // 200
// 	fmt.Fprintf(w, "%s\n", uj)
// }
