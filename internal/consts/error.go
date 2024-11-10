package consts

const (
	ErrUniqueViolation     = Error("unique_violation")
	ErrNullValueNotAllowed = Error("null value not allowed")
	ErrUndefinedTable      = Error("undefined_table")
	ErrNoRowsFound         = Error("no rows found")
	ErrForeignKeyViolation = Error("foreign key violation")
	ErrInternalServerError = Error("internal server error")
)

// Error to detect response status code 400
type Error string

func (e Error) Error() string {
	return string(e)
}
