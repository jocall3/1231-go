# 1231 Go SDK

<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go"><img src="https://pkg.go.dev/badge/github.com/stainless-sdks/1231-go.svg" alt="Go Reference"></a>

The 1231 Go SDK provides convenient access to the 1231 REST API from applications written in Go. It is generated with [Stainless](https://www.stainless.com/).

## Key Features

- **Type-Safe**: The SDK provides strongly typed request parameters and response objects to catch errors at compile time.
- **Production Ready**: Built-in support for automatic retries, timeouts, and robust error handling ensures reliability in production environments.
- **Developer Friendly**: Designed with a clean, functional API that integrates seamlessly with standard Go idioms.

## Installation

```go
import (
	"github.com/stainless-sdks/1231-go" // imported as jamesburvelocallaghaniiicitibankdemobusinessinc
)
```

Or to pin the version:

```sh
go get -u 'github.com/stainless-sdks/1231-go@v0.0.1'
```

**Requirements**: This library requires Go 1.22+.

## Quick Start

Here is a simplified example of how to use the SDK:

```go
package main

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/1231-go"
)

func main() {
	// Initialize the client
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient()

	// Make a request to register a user
	user, err := client.Users.Register(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.UserRegisterParams{
		Email:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("alice.w@example.com"),
		Name:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Alice Wonderland"),
		Password: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("SecureP@ssw0rd2024!"),
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("User ID: %+v\n", user.ID)
}
```

## Core Concepts

### Request Fields

All request parameters are wrapped in a generic `Field` type, which is used to distinguish zero values from null or omitted fields.

- **Omitted**: Any field not specified is not sent.
- **Values**: Use helpers like `String()`, `Int()`, or the generic `F[T]()`.
- **Null**: Use `Null[T]()` to explicitly send `null`.

```go
params := FooParams{
	Name: jamesburvelocallaghaniiicitibankdemobusinessinc.F("hello"),

	// Explicitly send `"description": null`
	Description: jamesburvelocallaghaniiicitibankdemobusinessinc.Null[string](),

	Point: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.Point{
		X: jamesburvelocallaghaniiicitibankdemobusinessinc.Int(0),
		Y: jamesburvelocallaghaniiicitibankdemobusinessinc.Int(1),
	}),
}
```

### Response Objects

All fields in response structs are value types. If a field is `null`, not present, or invalid, the corresponding field will simply be its zero value.

Response structs include a special `JSON` field for detailed property information:

```go
if res.Name == "" {
	// Check if the field was explicitly null or missing
	if res.JSON.Name.IsNull() {
		fmt.Println("Name is null")
	}
	if res.JSON.Name.IsMissing() {
		fmt.Println("Name is missing")
	}
}
```

## Advanced Usage

### RequestOptions

This library uses the functional options pattern. Options can be supplied to the client or individual requests.

```go
client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.Users.Register(context.TODO(), ...,
	// Override the header for this specific request
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
)
```

### Pagination

The library provides conveniences for working with paginated list endpoints.

- **Auto-Paging**: Use `.ListAutoPaging()` to iterate through all items.
- **Manual Paging**: Use `.List()` to fetch a single page and use `.GetNextPage()`.

### Errors

When the API returns a non-success status code, an error of type `*jamesburvelocallaghaniiicitibankdemobusinessinc.Error` is returned.

```go
if err != nil {
	var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))
		println(string(apierr.DumpResponse(true)))
	}
	panic(err.Error())
}
```

### Timeouts

Requests do not time out by default. Use context to configure a timeout for the request lifecycle.

```go
// Set a timeout for the request, including retries
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

client.Users.Register(ctx, params,
	// Set a per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

### File Uploads

Request parameters for file uploads are typed as `param.Field[io.Reader]`.

```go
// Use the helper to wrap any io.Reader with a filename and content type
fileParam := jamesburvelocallaghaniiicitibankdemobusinessinc.FileParam(reader, "file.txt", "text/plain")
```

### Retries

Certain errors (connection errors, 408, 409, 429, >=500) are automatically retried 2 times by default with exponential backoff.

```go
// Configure the default for all requests
client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(
	option.WithMaxRetries(0), // Disable retries
)

// Override per-request
client.Users.Register(context.TODO(), params, option.WithMaxRetries(5))
```

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).