// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &Client{}

// abbreviated

func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}

func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) {
	return DefaultClient.Post(url, bodyType, body)
}

func Head(url string) (resp *Response, err error) {
	return DefaultClient.Head(url)
}
