package ports

type Config interface {
	GetPort() (string, error)
}
