package gocd

// Clean ensures integrity of the schema by making sure
// empty elements are not printed to json.
func (a *Approval) Clean() {
	if a.Type == "success" {
		a.Authorization = nil
	}
}
