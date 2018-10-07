package langmap

type Filter struct {
	Query  string
	Values []interface{}
}

func NewFilter(query string, values ...interface{}) Filter {
	return Filter{
		Query:  query,
		Values: values,
	}
}
