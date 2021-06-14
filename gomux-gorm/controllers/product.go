package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pudjamansyurin/gomux-gorm/models"
)

const (
	ERR_INVALID_PAYLOAD = "Invalid request payload"
	ERR_INVALID_ID      = "Invalid item ID"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("page_size"))

	model := &models.Product{}
	products, err := model.All(page, pageSize)
	if err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}
	count := strconv.Itoa(model.Count())
	w.Header().Set("X-Total-Count", count)
	respJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	model := &models.Product{}
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		respError(w, http.StatusBadRequest, ERR_INVALID_PAYLOAD)
		return
	}
	defer r.Body.Close()

	if err := model.Create(); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respJSON(w, http.StatusCreated, model)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, ERR_INVALID_ID)
		return
	}

	model := &models.Product{}
	if err := model.Read(id); err != nil {
		respError(w, http.StatusNotFound, err.Error())
		return
	}
	respJSON(w, http.StatusOK, model)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, ERR_INVALID_ID)
		return
	}

	model := &models.Product{}
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		respError(w, http.StatusBadRequest, ERR_INVALID_PAYLOAD)
		return
	}
	defer r.Body.Close()

	if err := model.Update(id); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respJSON(w, http.StatusOK, model)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respError(w, http.StatusBadRequest, ERR_INVALID_ID)
		return
	}

	model := &models.Product{}
	if err := model.Delete(id); err != nil {
		respError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
