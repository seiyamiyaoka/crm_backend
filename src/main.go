package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seiyamiyaoka/crm_backend/configs"
	"github.com/seiyamiyaoka/crm_backend/models"
	"github.com/seiyamiyaoka/crm_backend/routes"
)

func main() {
	// db setup
	db := configs.ConnectDB()
	// db migration
	configs.Migrate(db)

	var customers []models.Customer
	result := db.Find(&customers)
	if count := result.RowsAffected; count > 0 {
		fmt.Println("Customers already exist")
	} else {
		fmt.Println("Customers do not exist, I'll create mock data")
		configs.SetUpSeedData()
	}
	router := mux.NewRouter()
	// set routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths := []string{}
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, _ := route.GetPathTemplate()
			paths = append(paths, t)
			return nil
		})
		html, err := template.ParseFiles("static/index.html")
		fmt.Println(router.Schemes())
		w.Header().Set("Content-Type", "text/html")
		if err != nil {
			log.Fatal(err)
		}
		if err := html.Execute(w, paths); err != nil {
			log.Fatal(err)
		}
	})

	routes.CustomerRoute(router)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", router)
}
