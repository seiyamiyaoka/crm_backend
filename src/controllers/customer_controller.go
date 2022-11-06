package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seiyamiyaoka/crm_backend/configs"
	"github.com/seiyamiyaoka/crm_backend/models"
)

func GetCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer
		params := mux.Vars(r)

		if ok := configs.DB.First(&customer, params["id"]); ok.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("{}")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}
}

func GetCustomers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customers []models.Customer
		configs.DB.Find(&customers)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	}
}

func AddCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var customer models.Customer

		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &customer)

		new_customer := models.Customer{
			Name:      customer.Name,
			Role:      customer.Role,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Contacted: customer.Contacted,
		}

		configs.DB.Create(&new_customer)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(new_customer)
	}
}

func UpdateCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer
		params := mux.Vars(r)
		if ok := configs.DB.First(&customer, params["id"]); ok.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("{}")
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		configs.DB.Save(&customer)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}
}

func DeleteCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer
		params := mux.Vars(r)
		if ok := configs.DB.First(&customer, params["id"]); ok.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("{}")
			return
		}

		configs.DB.Delete(&customer)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}
}

func MockCustomers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers := []models.Customer{
			{
				Name:      "[Mock]John Doe",
				Role:      "CEO",
				Email:     "mockuser@example.com",
				Phone:     "090-1234-5678",
				Contacted: false,
			},
			{
				Name:      "[Mock]Jane Doe",
				Role:      "CTO",
				Email:     "mockuser2@example.com",
				Phone:     "090-1234-5678",
				Contacted: false,
			},
		}

		json.NewEncoder(w).Encode(customers)
	}
}
