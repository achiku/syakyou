package echo

import (
	"io"
	"log"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/engine"

	"golang.org/x/net/context"
)

type (
	// Context reporesents the context of the current HTTP request. It holds request and
	// response objects, path, path parameters, data and registered handler.
	Context interface {
		Context() context.Context
		SetContext(context.Context)
		Deadline() (deadline time.Time, ok bool)
		Done() <-chan struct{}
		Err() error
		Value(key interface{}) interface{}
		Request() engine.Request
		Response() engine.Response
		Path() string
		SetPath(string)
		P(int) string
		Param(string) string
		ParamNames() []string
		SetParamNames(...string)
		ParamValues() []string
		SetParamValues(...string)

		QueryParam(string) string
		QueryParams() map[string][]string

		FormValue(string) string
		FormParams() map[string][]string
		FormFile(string) (*multipart.FileHeader, error)
		MultipartForm() (*multipart.Form, error)

		Cookie(string) (engine.Cookie, error)
		SetCookie(engine.Cookie)

		Cookies() []engine.Cookie

		Get(string) interface{}
		Set(string, interface{})

		Bind(interface{}) error

		Render(int, string, interface{}) error
		HTML(int, string) error
		JSON(int, interface{}) error
		JSONBlob(int, []byte) error
		JSONP(int, string, interface{}) error
		XML(int, interface{}) error
		XMLBlob(int, []byte) error
		File(string) error

		Attachment(io.ReadSeeker, string) error
		NoContent(int) error
		Redirect(int, string) error
		Error(err error)

		Handler() HandlerFunc
		SetHandler(HandlerFunc)

		Logger() log.Logger

		Echo() *Echo

		ServeContent(io.ReadSeeker, string, time.Time) error

		Reset(engine.Request, engine.Response)
	}

	echoContext struct {
		context  context.Context
		request  engine.Request
		response engine.Response
		path     string
		pnames   []string
		pvalues  []string
		handler  HandlerFunc
		echo     *Echo
	}
)

const (
	indexPage = "index.html"
)
