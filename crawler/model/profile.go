package model

import "encoding/json"

type Profile struct {
	Url           string
	Id            string
	Name          string
	Nickname      string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Occupation    string
	House         string
	Constellation string
	Car           string
	Nationality   string
	Smoke         string
	Drink         string
	City          string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
