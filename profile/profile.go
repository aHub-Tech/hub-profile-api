package profile

import (
//"github.com/ahub-tech/hub-profile-api/db"
)

type Profile struct {
	FullName    string
	Age         string
	Corporation string
	Experience  string
	Languages   string
	LinkedIn    string
	Twitter     string
	Facebook    string
	Instagram   string
	Autorization string
}
func NewProfile(fn, age, corp, exp, langs, lkin, tw, fb, ig, aut string) Profile {
	return Profile{fn, age, corp, exp, langs, lkin, tw, fb, ig, aut}
}
