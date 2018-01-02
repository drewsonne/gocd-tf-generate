package gocd

import (
	"encoding/json"
	"errors"
	"strconv"
)

// JSONString returns a string of this stage as a JSON object.
func (j *Job) JSONString() (body string, err error) {
	err = j.Validate()
	if err != nil {
		return
	}

	bdy, err := json.MarshalIndent(j, "", "  ")
	body = string(bdy)

	return
}

// Validate a job structure has non-nil values on correct attributes
func (j *Job) Validate() (err error) {
	if j.Name == "" {
		err = errors.New("`gocd.Jobs.Name` is empty")
	}
	return
}

func (tf *TimeoutField) UnmarshalJSON(b []byte) (err error) {
	value := string(b)
	var valInt int

	if value == `"never"` {
		valInt = 0
	} else {
		valInt, err = strconv.Atoi(value)
	}
	valTf := TimeoutField(valInt)
	tf = &valTf

	return
}

func (tf TimeoutField) MarshalJSON() (b []byte, err error) {
	return
}
