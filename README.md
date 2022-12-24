Faster (just gofiber)
=========================

Welcome to the Faster (just gofiber) project!

This project extends the [gofiber](https://github.com/gofiber/fiber) framework to allow for prefixless groups. This can be useful for separating different parts of your application and making your routes more organized and easier to read.

To use this extension, you will need to have gofiber installed in your project. Simply import the extension and use the `Group` function as you would normally use it in gofiber, but without specifying a prefix. The routes within the group will automatically be nested under the parent group.

Here is an example of how to use the prefixless group extension:



```go
package main

import (
	"github.com/yaameen/faster"
)

func main() {
	app := faster.New()

    prefixed = app.Prefix("app")
    prefixed.Get("/", info)
	// Create a prefixless group
	v1 := app.Group()
	{
		v1.Get("/", listUsers)
		v1.Get("/:id", getUser)
		v1.Post("/", createUser)
		v1.Put("/:id", updateUser)
		v1.Delete("/:id", deleteUser)
	}

	app.Listen(3000)
}
```

This will create the following routes:

-   `GET /app`
-   `GET /`
-   `GET /:id`
-   `POST /`
-   `PUT /:id`
-   `DELETE /:id`

I hope this extension helps make your gofiber projects more organized and easier to maintain. If you have any questions or issues, feel free to open an issue on the GitHub repository.