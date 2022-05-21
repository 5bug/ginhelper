package rest

import (
	"context"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ginContext = *gin.Context

// Context context define
type Context struct {
	context.Context
	ginContext
	Operator *Operator
}

// NewContext new context
func NewContext(c *gin.Context) *Context {
	ctx := &Context{}
	ctx.Context = context.Background()
	ctx.ginContext = c
	if v, ok := c.Get("operator"); ok {
		ctx.Operator = v.(*Operator)
	}
	return ctx
}

// ResultReply ...
type ResultReply struct {
	Result bool        `json:"result"`
	Msg    interface{} `json:"msg"`
}

// SetResult ...
func (r *ResultReply) SetResult(err error) {
	if err != nil {
		r.Result = false
		r.Msg = fmt.Errorf("失败，错误：%s", err.Error())
	} else {
		r.Result = true
		r.Msg = "成功"
	}
}

// HTTPResponse 统一返回格式
func HTTPResponse(c *Context, httpCode *int, err *error, data *interface{}) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error(r, string(debug.Stack()))
			return
		}
	}()
	if (*err) != nil {
		reply := &ResultReply{Result: false, Msg: (*err).Error()}
		c.JSON(*httpCode, reply)
		logrus.Error(c.Request.URL.Path+" error:", reply.Msg)
		return
	}
	if !IsNil(*data) {
		c.JSON(*httpCode, data)
	} else {
		reply := &ResultReply{Result: true}
		reply.SetResult(nil)
		c.JSON(*httpCode, reply)
	}
}

// Bind check bind uri
func Bind(c *gin.Context, req interface{}) error {
	if strings.Contains(c.FullPath(), ":") {
		c.ShouldBindUri(req) // nolint: errcheck
	}
	if err := c.Bind(req); err != nil {
		return err
	}
	return nil
}

// IsNil check obj of interface{} is nil
func IsNil(obj interface{}) bool {
	vi := reflect.ValueOf(obj)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
