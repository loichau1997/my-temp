package ginext

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/jfcore/common/common"
	"net/http"
)

type Request struct {
	GinCtx *gin.Context
	ctx    context.Context
}

type Response struct {
	Header http.Header
	*GeneralBody
}

func GetResponseMessage(code int) string {
	message := make(map[int]string)

	message[common.CODE_SUCCESS] = "success"
	message[common.CODE_CREATE_SUCCESSFULLY] = "Create successfully"
	message[common.CODE_UPDATE_SUCCESSFULLY] = "Update successfully"
	message[common.CODE_DELETE_SUCCESSFULLY] = "Delete successfully"
	message[common.CODE_BAD_REQUEST] = "Bad request"
	message[common.CODE_NOT_FOUND] = "Record not found"
	message[common.CODE_SERVER_ERROR] = "Server error"
	message[common.CODE_UNAUTHORIZED] = "Unauthorized"
	message[common.CODE_FORBIDDEN] = "Forbidden"

	if _, ok := message[code]; !ok {
		return "Unknown"
	}

	return message[code]
}

// NewResponse makes a new response with empty body
func NewResponse(code int) *Response {
	return &Response{
		GeneralBody: &GeneralBody{
			Code:    code,
			Message: GetResponseMessage(code),
		},
	}
}

func NewErrorResponse(code int, err string) *Response {
	return &Response{
		GeneralBody: &GeneralBody{
			Code:    code,
			Error:   err,
			Message: GetResponseMessage(code),
		},
	}
}

// NewResponseData makes a new response with body data
func NewResponseData(code int, data interface{}) *Response {
	return &Response{
		GeneralBody: &GeneralBody{
			Code:    code,
			Message: GetResponseMessage(code),
			Data:    data,
		},
	}
}

// NewResponseWithPager makes a new response with body data & pager
func NewResponseWithPager(code int, data interface{}, pager *Pager) *Response {
	return &Response{
		GeneralBody: NewBodyPaginated(
			code,
			GetResponseMessage(code),
			data,
			pager),
	}
}

type Handler func(r *Request) *Response

// NewRequest creates a new handler request
func NewRequest(c *gin.Context) *Request {
	ctx := FromGinRequestContext(c)
	req := &Request{
		GinCtx: c,
		ctx:    ctx,
	}

	return req
}

func (r *Request) Context() context.Context {
	if r.ctx == nil {
		r.ctx = context.Background()
	}

	return r.ctx
}

func WrapHandler(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			resp *Response
		)

		defer func() {

			if resp == nil {
				return
			}

			for k, v := range resp.Header {
				for _, v_ := range v {
					c.Header(k, v_)
				}
			}
			c.JSON(http.StatusOK, resp.GeneralBody)
		}()

		req := NewRequest(c)
		resp = handler(req)
	}
}

// MustBind does a binding on v with income request data
// it'll panic if any invalid data (and by design, it should be recovered by error handler middleware)
func (r *Request) MustBind(v interface{}) {
	r.MustNoError(r.GinCtx.ShouldBind(v))
}

func (r *Request) MustBindUri(v interface{}) {
	r.MustNoError(r.GinCtx.ShouldBindUri(v))
}

// MustNoError makes a ASSERT on err variable, panic when it's not nil
// then it must be recovered by WrapHandler
func (r *Request) MustNoError(err error) {
	if err != nil {
		//panic(err)
	}
}

func (r *Request) Uint64UserID() uint64 {
	return Uint64HeaderValue(r.GinCtx, common.HeaderUserID)
}

func (r *Request) Uint64TenantID() uint64 {
	return Uint64HeaderValue(r.GinCtx, common.HeaderTenantID)
}

func (r *Request) Param(key string) string {
	return r.GinCtx.Param(key)
}

func (r *Request) Query(key string) string {
	return r.GinCtx.Query(key)
}
