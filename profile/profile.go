package profile

type Profile struct {
	FullName     string
	Age          string
	Corporation  string
	Experience   string
	LinkedIn     string
	Twitter      string
	Facebook     string
	Instagram    string
	Autorization string
}

func NewProfile(fn, age, corp, exp, lkin, tw, fb, insta, aut string) Profile {
	return Profile{fn, age, corp, exp, lkin, tw, fb, insta, aut}
}

func(p *Profile) EditInfo(field string, value string) {
	switch field {
	case "fullname":
		p.FullName = value
	case "age":
		p.Age = value
	case "corp":
		p.Corporation = value
	case "exp":
		p.Experience = value
	case "linkedin":
		p.LinkedIn = value
	case "twitter":
		p.Twitter = value
	case "facebook":
		p.Facebook = value
	case "instagram":
		p.Instagram = value
	case "aut":
		p.Autorization = value
	}
}