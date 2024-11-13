package ports

type Config interface {
	GetPort() (string, error)
	GetProjectName() (string, error)
}
