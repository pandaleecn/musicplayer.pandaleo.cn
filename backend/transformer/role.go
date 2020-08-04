package transformer

type Role struct {
	Id          uint
	Name        string
	DisplayName string
	Description string
	Perms       []*Permission
	CreatedAt   string
}
