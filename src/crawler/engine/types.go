package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
