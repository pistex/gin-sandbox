package interfaces

type IHTTPError interface {
	GetStatus() int
	GetJSON() interface{}
}
