package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pudjamansyurin/gomux-gorm/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("page_size"))

	p := &models.Product{}
	products, err := p.List(page, pageSize)
	if err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := &models.Product{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		respError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.Create(); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respJSON(w, http.StatusCreated, p)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	p := &models.Product{ID: id}
	if err := p.Read(); err != nil {
		respError(w, http.StatusNotFound, err.Error())
		return
	}
	respJSON(w, http.StatusOK, p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := &models.Product{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		respError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p.ID = id
	if err := p.Update(); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON(w, http.StatusOK, p)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := &models.Product{ID: id}
	if err := p.Delete(); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
