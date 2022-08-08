package db

type People struct {
	IdentityNumber int64 `gorm:"primarykey"`
	Name           string
	Age            int64
	Profession     string
	IDFamily       string
	Family         Family `gorm:"references:IDFamily"`
}
type Family struct {
	IDFamily  string `gorm:"primarykey"`
	NumberMem int64
	Address   string
	Town      Town `gorm:"references:IDTown"`
}
type Town struct {
	IDTown   string `gorm:"primarykey"`
	NameTown string
}
