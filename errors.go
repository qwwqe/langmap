package langmap

const (
	ErrDatabaseFailure   = "database: non-specific failure"
	ErrDatabaseNotFound  = "database: record not found"
	ErrInvalidResourceId = "invalid resource id requested"
	ErrJsonFailed        = "json: failed to bind data"
)

type ErrorsJSON struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func NewErrorsJSON(errors []error) []ErrorsJSON {
	d := make([]ErrorsJSON, len(errors))
	for i, e := range errors {
		d[i].Error = e.Error()
		d[i].Data = e
	}
	return d
}
