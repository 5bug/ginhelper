package templates

//APITpl TODO 反引号嵌套反引号的处理方法
const APITpl = `
package api

// HelloRequest hello request
// @Router /hello [get]
// @Auth false
type HelloRequest struct {
	Name   string [*]form:"name" binding:"required"[*]
}

// HelloReply hello reply
type HelloReply struct {
	Msg   string [*]json:"msg"[*]
}
`
