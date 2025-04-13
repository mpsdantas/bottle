package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mpsdantas/bottle"
)

type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *bottle.HandlerCtx) bool

	// AllowOriginsFunc defines a function that will set the 'Access-Control-Allow-Origin'
	// response header to the 'origin' request header when returned true. This allows for
	// dynamic evaluation of allowed origins. Note if AllowCredentials is true, wildcard origins
	// will be not have the 'Access-Control-Allow-Credentials' header set to 'true'.
	//
	// Optional. Default: nil
	AllowOriginsFunc func(origin string) bool

	// AllowOrigin defines a comma separated list of origins that may access the resource.
	//
	// Optional. Default value "*"
	AllowOrigins string

	// AllowMethods defines a list methods allowed when accessing the resource.
	// This is used in response to a preflight request.
	//
	// Optional. Default value "GET,POST,HEAD,PUT,DELETE,PATCH"
	AllowMethods string

	// AllowHeaders defines a list of request headers that can be used when
	// making the actual request. This is in response to a preflight request.
	//
	// Optional. Default value "".
	AllowHeaders string

	// AllowCredentials indicates whether or not the response to the request
	// can be exposed when the credentials flag is true. When used as part of
	// a response to a preflight request, this indicates whether or not the
	// actual request can be made using credentials. Note: If true, AllowOrigins
	// cannot be set to a wildcard ("*") to prevent security vulnerabilities.
	//
	// Optional. Default value false.
	AllowCredentials bool

	// ExposeHeaders defines a whitelist headers that clients are allowed to
	// access.
	//
	// Optional. Default value "".
	ExposeHeaders string

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached.
	// If you pass MaxAge 0, Access-Control-Max-Age header will not be added and
	// browser will use 5 seconds by default.
	// To disable caching completely, pass MaxAge value negative. It will set the Access-Control-Max-Age header 0.
	//
	// Optional. Default value 0.
	MaxAge int
}

func New(config ...Config) bottle.Handler {
	if len(config) == 0 {
		return toBottleHandler(cors.New())
	}

	cfg := cors.Config{
		Next:             nil,
		AllowOriginsFunc: config[0].AllowOriginsFunc,
		AllowOrigins:     config[0].AllowOrigins,
		AllowMethods:     config[0].AllowMethods,
		AllowHeaders:     config[0].AllowHeaders,
		AllowCredentials: config[0].AllowCredentials,
		ExposeHeaders:    config[0].ExposeHeaders,
		MaxAge:           config[0].MaxAge,
	}

	if config[0].Next != nil {
		cfg.Next = func(c *fiber.Ctx) bool {
			return config[0].Next(bottle.WrapEngineContext(c))
		}
	}

	return toBottleHandler(cors.New(cfg))
}

func toBottleHandler(handler fiber.Handler) bottle.Handler {
	return func(r *bottle.HandlerCtx) error {
		return handler(r.EngineCtx())
	}
}
