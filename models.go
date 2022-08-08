package main

type People struct {
	IdentityNumber int64
	Name           string
	Age            int64
	Profession     string
}
type Family struct {
	NumberMem int64
	Address   string
	Members   []People
}
type Town struct {
	NameTown  string
	NumFamily int64
	Familys   []Family
}
