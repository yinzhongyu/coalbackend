package http

type Client interface {
	Post(url string, body ...string) (int, []byte, error)
	Get(url string) (int, []byte, error)
}
