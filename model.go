package main

type Cv struct {
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	JobTitle     string         `json:"jobTitle"`
	Contact      Contact        `json:"contact"`
	Links        Links          `json:"links"`
	Summary      []string       `json:"summary"`
	Technologies []Technologies `json:"technologies"`
	Employment   []Experience   `json:"employment"`
	Education    []Experience   `json:"education"`
	Interests    []string       `json:"interests"`
}

type Contact struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

type Links struct {
	GitLab  string `json:"gitlab"`
	GitHub  string `json:"github"`
	Website string `json:"website,omitempty"`
}

type Technologies struct {
	Category string   `json:"category"`
	Values   []string `json:"values"`
}

type Experience struct {
	Company       string   `json:"company,omitempty"`
	School        string   `json:"school,omitempty"`
	Location      string   `json:"location"`
	JobTitle      string   `json:"jobTitle,omitempty"`
	Qualification string   `json:"qualification,omitempty"`
	Duration      Duration `json:"duration"`
	Details       []string `json:"details"`
}

type Duration struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
