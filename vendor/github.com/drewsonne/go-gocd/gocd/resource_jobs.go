package gocd

import (
	"encoding/json"
	"errors"
)

// JSONString returns a string of this stage as a JSON object.
func (j *Job) JSONString() (string, error) {
	err := j.Validate()
	if err != nil {
		return "", err
	}

	bdy, err := json.MarshalIndent(j, "", "  ")
	return string(bdy), err
}

// Validate a job structure has non-nil values on correct attributes
func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("`gocd.Jobs.Name` is empty")
	}
	return nil
}
