package pangolin

import "fmt"

func defaultHandler(ctx *PangolinCtx) {
	fmt.Println(ctx.body)
	ctx.w.Write([]byte("ok"))
}

func defaultControlHandler(ctx *PangolinCtx) {
	if ctx.IsNewConnection(){

	}
}

var _ Header = (*defaultHeader)(nil)

type defaultHeader struct {
	headers map[string]string
}

func NewDefaultHeader() *defaultHeader {
	return &defaultHeader{
		headers: make(map[string]string),
	}
}

func (h *defaultHeader) Get(key string) string {
	return h.headers[key]
}

func (h *defaultHeader) Set(key, value string) {
	h.headers[key] = value
}

func (h *defaultHeader) IsNewConnection() bool {
	return h.headers["type"] == "connection"
}

func (h *defaultHeader)GetMethod()string{
	return h.headers["method"]
}

var _ Body = (*defaultBody)(nil)

type defaultBody struct {
	body map[string]interface{}
}

func NewDefaultBody() *defaultBody {
	return &defaultBody{
		body: make(map[string]interface{}),
	}
}

func (b *defaultBody) Get(key string) interface{} {
	return b.body[key]
}

func (b *defaultBody) Set(key string, value interface{}) {
	b.body[key] = value
}
