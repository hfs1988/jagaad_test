package adapter

type HTTPAdapter interface {
	Get(url string, data interface{}) error
}

type CSVAdapter interface {
	Write(filename string, data [][]string) error
	Read(filename string) ([][]string, error)
}

type LogAdapter interface {
	Error(err error)
	Info(msg string)
}
