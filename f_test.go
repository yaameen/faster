package faster_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/yaameen/faster"
)

func BenchmarkGet(b *testing.B) {
	app := faster.New()
	app.Get("/", func(c *faster.Ctx) error {
		return c.SendString("Hello, World!")
	})

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		app.Test(req)
	}
}

func BenchmarkJsonGet(b *testing.B) {
	app := faster.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello, World!")
	})
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		app.Test(req)
	}
}

func BenchmarkFiberGet(b *testing.B) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		app.Test(req)
	}
}

func BenchmarkFiberJsonGet(b *testing.B) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello, World!")
	})
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		app.Test(req)
	}
}

func CheckResponse(app faster.FastApp, method, route, cmp string, t *testing.T) {
	req := httptest.NewRequest(method, route, nil)
	response, err := app.Test(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Status code should be 200, but got %d", response.StatusCode)
	}
	resp, _ := ioutil.ReadAll(response.Body)
	if string(resp) != cmp {
		t.Errorf("Response should be '%s', but got '%s'", cmp, string(resp))
	}
}

// Test Get
func TestHttp(t *testing.T) {

	suite := []struct {
		name     string
		method   string
		reqRoute string
		route    string
		response string
		expected string
		setup    func(faster.FastRouter) faster.FastRouter
	}{
		{
			name:     "Test Basic Get",
			method:   "GET",
			reqRoute: "/",
			route:    "/",
			response: "Hello, World!",
			expected: "Hello, World!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app
			},
		},
		{
			name:     "Test Basic Post",
			method:   "POST",
			reqRoute: "/",
			route:    "/",
			response: "Hello, World!",
			expected: "Hello, World!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app
			},
		},
		{
			name:     "Test Basic Patch",
			method:   "PATCH",
			reqRoute: "/",
			route:    "/",
			response: "Hello, World!",
			expected: "Hello, World!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app
			},
		},
		{
			name:     "Test With A Prefix",
			method:   "GET",
			reqRoute: "/v1/hello",
			route:    "/hello",
			response: "Hello, World!",
			expected: "Hello, World!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Prefix("/v1")
			},
		},
		{
			name:     "Test With A Prefixless Route",
			method:   "PATCH",
			reqRoute: "/hello",
			route:    "/hello",
			response: "Hello, World!",
			expected: "Hello, World!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Group()
			},
		},
		{
			name:     "Test With A Prefixless Route and middleware",
			method:   "PATCH",
			reqRoute: "/hello",
			route:    "/hello",
			response: "Hello, World!",
			expected: "From Middleware!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Group(func(c *faster.Ctx) error {
					return c.SendString("From Middleware!")
				})
			},
		},
		{
			name:     "Test With A Prefix and middleware",
			method:   "PATCH",
			reqRoute: "/v1/hello",
			route:    "/hello",
			response: "Hello, World!",
			expected: "From Middleware!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Prefix("v1").Group(func(c *faster.Ctx) error {
					return c.SendString("From Middleware!")
				})
			},
		},
		{
			name:     "Test With A Prefix And Middleware But Different Order",
			method:   "PATCH",
			reqRoute: "/v1/hello",
			route:    "/hello",
			response: "Hello, World!",
			expected: "Should carry forward middleware!",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Group(func(c *faster.Ctx) error {
					return c.SendString("Should carry forward middleware!")
				}).Prefix("v1")
			},
		},
		{
			name:     "Test With A Static",
			method:   "STATIC",
			reqRoute: "/storage/test.txt",
			route:    "/",
			response: "",
			expected: "Hello From File",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Prefix("storage")
			},
		},
		{
			name:     "Test With A Static Root With Double Prefix",
			method:   "STATIC",
			reqRoute: "/storage/v1",
			route:    "/",
			response: "",
			expected: "Hello from index.html",
			setup: func(app faster.FastRouter) faster.FastRouter {
				return app.Prefix("storage").Prefix("v1")
			},
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			app := faster.New()
			if test.method == "STATIC" {
				test.setup(app).Static(test.route, "./test")
				// reset method to GET, as STATIC is not a valid http method
				test.method = "GET"
			} else {
				test.setup(app).Add(test.method, test.route, func(c *faster.Ctx) error {
					return c.SendString(test.response)
				})
			}
			CheckResponse(*app, test.method, test.reqRoute, test.expected, t)

		})
	}

}
