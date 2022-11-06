package routes

import (
	"github.com/gorilla/mux"
	"github.com/seiyamiyaoka/crm_backend/controllers"
	"github.com/seiyamiyaoka/crm_backend/routes/middleware"
)

func CustomerRoute(router *mux.Router) {
	router.HandleFunc("/customers", controllers.GetCustomers()).Methods("GET")
	router.HandleFunc("/customers", controllers.AddCustomer()).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.GetCustomer()).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomer()).Methods("PUT")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer()).Methods("DELETE")
	router.HandleFunc("/customers_mocks", controllers.MockCustomers()).Methods("GET")
	router.Use(middleware.HeaderMiddleware)
}
