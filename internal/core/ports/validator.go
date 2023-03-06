package ports

type Validator interface {
	Struct(s interface{}) error
}
