package errors

type BaseErrorInterface interface {
	Error() map[string]interface{}
	HttpStatusCode() int
}
