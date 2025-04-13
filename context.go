package bottle

import (
	"context"
	"crypto/tls"
	"io"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type HandlerCtx struct {
	ctx *fiber.Ctx
}

type Cookie struct {
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	Path        string    `json:"path"`
	Domain      string    `json:"domain"`
	MaxAge      int       `json:"max_age"`
	Expires     time.Time `json:"expires"`
	Secure      bool      `json:"secure"`
	HTTPOnly    bool      `json:"http_only"`
	SameSite    string    `json:"same_site"`
	SessionOnly bool      `json:"session_only"`
}

func WrapEngineContext(c *fiber.Ctx) *HandlerCtx {
	return &HandlerCtx{ctx: c}
}

func (c *HandlerCtx) Accepts(offers ...string) string {
	return c.ctx.Accepts(offers...)
}

func (c *HandlerCtx) AcceptsCharsets(offers ...string) string {
	return c.ctx.AcceptsCharsets(offers...)
}

func (c *HandlerCtx) AcceptsEncodings(offers ...string) string {
	return c.ctx.AcceptsEncodings(offers...)
}

func (c *HandlerCtx) AcceptsLanguages(offers ...string) string {
	return c.ctx.AcceptsLanguages(offers...)
}

func (c *HandlerCtx) Append(field string, values ...string) {
	c.ctx.Append(field, values...)
}

func (c *HandlerCtx) Attachment(filename ...string) {
	c.ctx.Attachment(filename...)
}

func (c *HandlerCtx) BaseURL() string {
	return c.ctx.BaseURL()
}

func (c *HandlerCtx) BodyRaw() []byte {
	return c.ctx.BodyRaw()
}

func (c *HandlerCtx) Body() []byte {
	return c.ctx.Body()
}

func (c *HandlerCtx) BodyParser(out interface{}) error {
	return c.ctx.BodyParser(out)
}

func (c *HandlerCtx) ClearCookie(key ...string) {
	c.ctx.ClearCookie(key...)
}

func (c *HandlerCtx) Context() *fasthttp.RequestCtx {
	return c.ctx.Context()
}

func (c *HandlerCtx) UserContext() context.Context {
	return c.ctx.UserContext()
}

func (c *HandlerCtx) SetUserContext(ctx context.Context) {
	c.ctx.SetUserContext(ctx)
}

func (c *HandlerCtx) Cookie(cookie Cookie) {
	c.ctx.Cookie(&fiber.Cookie{
		Name:        cookie.Name,
		Value:       cookie.Value,
		Path:        cookie.Path,
		Domain:      cookie.Domain,
		MaxAge:      cookie.MaxAge,
		Expires:     cookie.Expires,
		Secure:      cookie.Secure,
		HTTPOnly:    cookie.HTTPOnly,
		SameSite:    cookie.SameSite,
		SessionOnly: cookie.SessionOnly,
	})
}

func (c *HandlerCtx) Cookies(key string, defaultValue ...string) string {
	return c.ctx.Cookies(key, defaultValue...)
}

func (c *HandlerCtx) CookieParser(out interface{}) error {
	return c.ctx.CookieParser(out)
}

func (c *HandlerCtx) Download(file string, filename ...string) error {
	return c.ctx.Download(file, filename...)
}

func (c *HandlerCtx) Request() *fasthttp.Request {
	return c.ctx.Request()
}

func (c *HandlerCtx) Response() *fasthttp.Response {
	return c.ctx.Response()
}

func (c *HandlerCtx) Format(body interface{}) error {
	return c.ctx.Format(body)
}

func (c *HandlerCtx) FormFile(key string) (*multipart.FileHeader, error) {
	return c.ctx.FormFile(key)
}

func (c *HandlerCtx) FormValue(key string, defaultValue ...string) string {
	return c.ctx.FormValue(key, defaultValue...)
}

func (c *HandlerCtx) Fresh() bool {
	return c.ctx.Fresh()
}

func (c *HandlerCtx) Get(key string, defaultValue ...string) string {
	return c.ctx.Get(key, defaultValue...)
}

func (c *HandlerCtx) GetRespHeader(key string, defaultValue ...string) string {
	return c.ctx.GetRespHeader(key, defaultValue...)
}

func (c *HandlerCtx) GetReqHeaders() map[string][]string {
	return c.ctx.GetReqHeaders()
}

func (c *HandlerCtx) GetRespHeaders() map[string][]string {
	return c.ctx.GetRespHeaders()
}

func (c *HandlerCtx) Hostname() string {
	return c.ctx.Hostname()
}

func (c *HandlerCtx) Port() string {
	return c.ctx.Port()
}

func (c *HandlerCtx) IP() string {
	return c.ctx.IP()
}

func (c *HandlerCtx) IPs() []string {
	return c.ctx.IPs()
}

func (c *HandlerCtx) Is(extension string) bool {
	return c.ctx.Is(extension)
}

func (c *HandlerCtx) JSON(data interface{}, ctype ...string) error {
	return c.ctx.JSON(data, ctype...)
}

func (c *HandlerCtx) JSONP(data interface{}, callback ...string) error {
	return c.ctx.JSONP(data, callback...)
}

func (c *HandlerCtx) XML(data interface{}) error {
	return c.ctx.XML(data)
}

func (c *HandlerCtx) Links(link ...string) {
	c.ctx.Links(link...)
}

func (c *HandlerCtx) Locals(key interface{}, value ...interface{}) interface{} {
	return c.ctx.Locals(key, value...)
}

func (c *HandlerCtx) Location(path string) {
	c.ctx.Location(path)
}

func (c *HandlerCtx) Method(override ...string) string {
	return c.ctx.Method(override...)
}

func (c *HandlerCtx) MultipartForm() (*multipart.Form, error) {
	return c.ctx.MultipartForm()
}

func (c *HandlerCtx) ClientHelloInfo() *tls.ClientHelloInfo {
	return c.ctx.ClientHelloInfo()
}

func (c *HandlerCtx) Next() error {
	return c.ctx.Next()
}

func (c *HandlerCtx) RestartRouting() error {
	return c.ctx.RestartRouting()
}

func (c *HandlerCtx) OriginalURL() string {
	return c.ctx.OriginalURL()
}

func (c *HandlerCtx) Params(key string, defaultValue ...string) string {
	return c.ctx.Params(key, defaultValue...)
}

func (c *HandlerCtx) AllParams() map[string]string {
	return c.ctx.AllParams()
}

func (c *HandlerCtx) ParamsParser(out interface{}) error {
	return c.ctx.ParamsParser(out)
}

func (c *HandlerCtx) ParamsInt(key string, defaultValue ...int) (int, error) {
	return c.ctx.ParamsInt(key, defaultValue...)
}

func (c *HandlerCtx) Path(override ...string) string {
	return c.ctx.Path(override...)
}

func (c *HandlerCtx) Protocol() string {
	return c.ctx.Protocol()
}

func (c *HandlerCtx) Query(key string, defaultValue ...string) string {
	return c.ctx.Query(key, defaultValue...)
}

func (c *HandlerCtx) Queries() map[string]string {
	return c.ctx.Queries()
}

func (c *HandlerCtx) QueryInt(key string, defaultValue ...int) int {
	return c.ctx.QueryInt(key, defaultValue...)
}

func (c *HandlerCtx) QueryBool(key string, defaultValue ...bool) bool {
	return c.ctx.QueryBool(key, defaultValue...)
}

func (c *HandlerCtx) QueryFloat(key string, defaultValue ...float64) float64 {
	return c.ctx.QueryFloat(key, defaultValue...)
}

func (c *HandlerCtx) QueryParser(out interface{}) error {
	return c.ctx.QueryParser(out)
}

func (c *HandlerCtx) ReqHeaderParser(out interface{}) error {
	return c.ctx.ReqHeaderParser(out)
}

func (c *HandlerCtx) Redirect(location string, status ...int) error {
	return c.ctx.Redirect(location, status...)
}

func (c *HandlerCtx) RedirectBack(fallback string, status ...int) error {
	return c.ctx.RedirectBack(fallback, status...)
}

func (c *HandlerCtx) Render(name string, bind interface{}, layouts ...string) error {
	return c.ctx.Render(name, bind, layouts...)
}

func (c *HandlerCtx) Secure() bool {
	return c.ctx.Secure()
}

func (c *HandlerCtx) Send(body []byte) error {
	return c.ctx.Send(body)
}

func (c *HandlerCtx) SendFile(file string, compress ...bool) error {
	return c.ctx.SendFile(file, compress...)
}

func (c *HandlerCtx) SendStatus(status int) error {
	return c.ctx.SendStatus(status)
}

func (c *HandlerCtx) SendString(body string) error {
	return c.ctx.SendString(body)
}

func (c *HandlerCtx) SendStream(stream io.Reader, size ...int) error {
	return c.ctx.SendStream(stream, size...)
}

func (c *HandlerCtx) Set(key, val string) {
	c.ctx.Set(key, val)
}

func (c *HandlerCtx) Subdomains(offset ...int) []string {
	return c.ctx.Subdomains(offset...)
}

func (c *HandlerCtx) Stale() bool {
	return c.ctx.Stale()
}

func (c *HandlerCtx) Status(status int) *HandlerCtx {
	c.ctx.Status(status)
	return c
}

func (c *HandlerCtx) String() string {
	return c.ctx.String()
}

func (c *HandlerCtx) Type(extension string, charset ...string) *HandlerCtx {
	c.ctx.Type(extension, charset...)
	return c
}

func (c *HandlerCtx) Vary(fields ...string) {
	c.ctx.Vary(fields...)
}

func (c *HandlerCtx) Write(p []byte) (int, error) {
	return c.ctx.Write(p)
}

func (c *HandlerCtx) Writef(f string, a ...interface{}) (int, error) {
	return c.ctx.Writef(f, a...)
}

func (c *HandlerCtx) WriteString(s string) (int, error) {
	return c.ctx.WriteString(s)
}

func (c *HandlerCtx) XHR() bool {
	return c.ctx.XHR()
}

func (c *HandlerCtx) IsProxyTrusted() bool {
	return c.ctx.IsProxyTrusted()
}

func (c *HandlerCtx) IsFromLocal() bool {
	return c.ctx.IsFromLocal()
}

func (c *HandlerCtx) EngineCtx() *fiber.Ctx {
	return c.ctx
}
