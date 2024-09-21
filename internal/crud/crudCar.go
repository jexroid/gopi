package crud

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/jexroid/gopi/api/models"
	"gorm.io/gorm"
	"net/http"
)

func (db Database) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "application/json")

	var car models.Cars
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&car)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func (db Database) Read(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	var car models.Cars
	result := db.DB.First(&car, "id = ?", uuid)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(car)
}

func (db Database) All(w http.ResponseWriter, r *http.Request) {
	var cars []models.Cars
	result := db.DB.Find(&cars)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func (db Database) Update(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	var updateCar models.Cars
	if err := json.NewDecoder(r.Body).Decode(&updateCar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var car models.Cars
	result := db.DB.First(&car, "id = ?", uuid)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	db.DB.Model(&car).Updates(updateCar)
	json.NewEncoder(w).Encode(car)
}

func (db Database) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	result := db.DB.Delete(&models.Cars{}, "id = ?", uuid)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
