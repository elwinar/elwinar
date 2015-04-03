package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Validator struct {
	Errs []error
	r *http.Request
}

func NewValidator(r *http.Request) *Validator {
	r.ParseForm()
	return &Validator{
		r: r,
	}
}

func (v *Validator) HasErrors() bool {
	return len(v.Errs) != 0
}

func (v *Validator) Errors() []string {
	var errs []string
	for _, err := range v.Errs {
		errs = append(errs, err.Error())
	}
	return errs
}

func (v *Validator) NotEmpty(field string) {
	if len(v.r.FormValue(field)) == 0 {
		v.Errs = append(v.Errs, fmt.Errorf("%s is empty", field))
	}
}

func (v *Validator) MaxLen(field string, max int) {
	if len(v.r.FormValue(field)) > max {
		v.Errs = append(v.Errs, fmt.Errorf("%s is too long (%d chars max)", field, max))
	}
}

func (v *Validator) DoesntExists(field, table, column string, allowed ...string) {
	for _, a := range allowed {
		if a == v.r.FormValue(field) {
			return
		}
	}
	
	err := db.Get(new(int), "SELECT id FROM `"+table+"` WHERE `"+column+"` = ?", v.r.FormValue(field))
		
	if err == sql.ErrNoRows {
		return
	}
	
	if err != nil {
		panic(err)
	}
	
	v.Errs = append(v.Errs, fmt.Errorf("%s already exists", field))
}

func (v *Validator) Validate(field string, validator func(string) error) {
	if err := validator(v.r.FormValue(field)); err != nil {
		v.Errs = append(v.Errs, fmt.Errorf("%s %s", field, err))
	}
}
