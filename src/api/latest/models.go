package latest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Orders struct {
	gorm.Model
	Project     string `gorm:"not null"`
	Environment string `gorm:"not null"` // poc,dev,staging,prod
	CostCentre  int    `gorm:"not null"`
	Owner       string `gorm:"not null"`
	Service     string `gorm:"not null"` // aws_account,aws_ami, aws_j5
	Tier        string // Tier: 1,2,3
	Status      string `sql:"DEFAULT:'submitted'"`
	Tracking    string
	Account     Accounts
	Image       Images
}

type Accounts struct {
	gorm.Model
	Environment string `gorm:"not null"`
	Version     int    `gorm:"not null"`
	AccountId   int    `gorm:"not null;unique"`
	Status      string
	Email       string `gorm:"not null;unique"`
}

type Images struct {
	gorm.Model
	Reference   string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
}

// ** ORDERS **
func (app *App) get_orders(w http.ResponseWriter, r *http.Request) {

}

func (app *App) get_order(w http.ResponseWriter, r *http.Request) {

}
func (app *App) create_order(w http.ResponseWriter, r *http.Request) {

}
func (app *App) update_order(w http.ResponseWriter, r *http.Request) {

}
func (app *App) delete_order(w http.ResponseWriter, r *http.Request) {

}

// ** ACCOUNTS ** WEB HANDLERS

func (app *App) get_accounts(w http.ResponseWriter, r *http.Request) {

	accounts, rowCount := get_accounts(app.Db)

	if rowCount == 0 {
		respondWithError(w, 404, "record not found")
		return
	}

	respondWithJSON(w, http.StatusOK, accounts)

}

func (app *App) get_account(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	account := Accounts{Model: gorm.Model{ID: uint(id)}}

	if RecordNotFound := account.get(app.Db); RecordNotFound {
		respondWithError(w, 404, "record not found")
		return
	}

	respondWithJSON(w, http.StatusOK, account)

}

func (app *App) create_account(w http.ResponseWriter, r *http.Request) {

	var account Accounts
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := account.create(app.Db); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, account)

}

// ** EMAIL ** WEB HANDLERS

func (app *App) get_next_email(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Id    int
		Email string
	}

	var result Result
	app.Db.Table("accounts").Where("status = ?", "").Select("id, email").Order("email desc").Limit(1).Scan(&result)

	respondWithJSON(w, http.StatusOK, &result)

}

// ** ACCOUNTS ** DATABASE HANDLERS

func (account *Accounts) create(db *gorm.DB) (err error) {
	return db.Create(&account).Error
}

func get_accounts(db *gorm.DB) (account []Accounts, rowCount int64) {

	var accounts []Accounts
	rowCount = db.Find(&accounts).RowsAffected

	return accounts, rowCount
}

func (account *Accounts) get(db *gorm.DB) (RecordNotFound bool) {
	return db.Unscoped().First(&account, &account.ID).RecordNotFound()
}

//func (a *Accounts) update(db *gorm.DB) {
//	db.Save(&a).Where("Id", &a.Id)
//}
//
//func (a *Accounts) del(db *gorm.DB) {
//	db.Delete(a).Where("Id", &a.Id)
//}
