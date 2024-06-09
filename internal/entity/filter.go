package entity

type Filter struct {
	Status      string
	Name        string
	PixKeyType  string
	PixKeyValue string
	Limit       int
	Page        int
}

func NewFilter(status, name, pixKeyType, pixKeyValue string, limit, page int) Filter {
	return Filter{
		Status:      status,
		Name:        name,
		PixKeyType:  pixKeyType,
		PixKeyValue: pixKeyValue,
		Limit:       limit,
		Page:        page,
	}
}
