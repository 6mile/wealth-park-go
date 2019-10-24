# Design of the Wealth Park API

## Terminology

- `Resource` (or Entity): An object representing a resource as defined in our domain model, e.g. `Product`, `Purchaser`, etc.

- `Controller`: A set of `Endpoints` (or routes or route handlers) that deal mainly with receiving and sending payloads in various formats, e.g. JSON. `Controllers` represent our _endpoint layer_.

- `Service`: A set of related functions that implement our business logic. `Services` represent our _service layer_.

- `Model` (or Store): A set of related functions that deal with creating, reading, updating and deleting one `Resource`. `Models` represent our _data layer_.


## Folder Structure

```bash
/cmd     # Contains bootstraps for all our executables, i.e. the things we build and put in docker containers to run.
/wpark # Main folder that contains all our Go code.
```


## Go Packages

```bash
/wpark/apiserver  # The web layer. This server is used to expose `Endpoints` over HTTPS / TLS.
/wpark/config     # The Core API config.
/wpark/controller # Contains all our `Controllers`, e.g. `ProductController`, `PurchaserController`, etc. This constitutes our endpoint layer.
/wpark/core       # Core API root package and domain model. Contains all interface definitions (services, models, payloads, etc.) and constructor functions for all resources.
/wpark/mysql      # MySQL implementation of our data layer (implementation of all model interfaces defined in /wpark/core).
/wpark/e2e        # All our end-to-end tests.
/wpark/mock       # Mock implementations of all services and models defined in /wpark/core.
/wpark/pkg/*      # Various standalone packages, e.g. /wpark/pkg/logger, which handles logging.
/wpark/service    # All service implementations, e.g `ProductService`, `PurchaserService`, etc. This constitutes our service layer.
```

## Server Dependencies

1. MySQL >= v5.7.23 - data store.

## Rules & Guidelines

This section outlines the rules that govern the design of the system.

- _Dependency Injection_ (DI): To keep things modular, easy to test, and to separate concerns, we want to be coding to interfaces, and injecting implementations of those interfaces into our `Controllers` and `Services` as part of our bootstrap. Using DI we're able to mock all our dependencies and unit test each layer in isolation. It's also trivial to swap out one implementation for another whenever we want.

- _Endpoint Layer_: `Controllers` make up our _endpoint layer_. A `Controller` is basically just a set of handlers (e.g route handlers). `Controllers` normally depend on one or more `Services` to do their job.

  A typical handler (`Controller` method) should ideally do _only_ the following:

  1. Receive a payload (e.g. JSON or URL parameters) over HTTPS and validate the payload against a validation schema, e.g. parse an incoming JSON payload into a valid `CreateProductRequestV1` struct.
  2. Pass validated data to a `Service` and wait for a response, e.g. by calling `ProductService.CreateProduct(...)`.
  3. Create a _versioned_ endpoint response, based on the `Service` response, and send it back to the caller, e.g create and serialize a `CreateProductResponseV1` struct.

  _AVOID doing the following_:

  - Inject `Models` into a `Controller` and use those `Models` directly. For example, calling `ProductModel.Update(...)` directly from within a `Controller`, instead of going through `ProductService`, is considered wrong. It's important to note that a `Model` is _owned by one and only one_ `Service`, and only that `Service` ever gets to call it's `Model`.

- _Service Layer_: `Services` contain all our business logic. `Services` own one or more `Models`, and _may_ depend on one or more other `Services` to do their job.
  
- _Data Layer_: `Models` make up our _data layer_. A `Model` implements the CRUD operations for a `Resource` and _should not_ contain any business logic.


## API Design Guide

We draw inspiration from the following APIs and documents:

- Go Project layout(by Ben Johnson): https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- Google Cloud API: https://cloud.google.com/apis/design/

## Tests

This section outlines how to run tests.  
We use `go modules` for dependency management. [Using Go Modules](https://blog.golang.org/using-go-modules)

> As of Go 1.11, the go command enables the use of modules when the current directory or any parent directory has a go.mod, provided the directory is outside $GOPATH/src. (Inside $GOPATH/src, for compatibility, the go command still runs in the old GOPATH mode, even if a go.mod is found. See the go command documentation for details.) Starting in Go 1.13, module mode will be the default for all development.

As noted above, please make sure that you checkout this repository outside of `$GOPATH`. Since we are using `go modules` we do not need to explicitly download the dependencies, `go build` or `go test` would take care of it.  
We use `go test` to run our tests and show us our test coverage in the output.


- Run unit tests.

  ```bash
  bash test.sh
  ```

  This will run all the unit tests in isolation with mocked dependencies.  
  Pasting output here for convenience:

  ```bash
  Running unit tests ..
  ok  	github.com/.../wpark/core	0.002s	coverage: 89.5% of statements
  ok  	github.com/.../wpark/service	0.003s	coverage: 90.3% of statements
  ok  	github.com/.../wpark/apiserver	0.006s	coverage: 56.8% of  statements
  ok  	github.com/.../wpark/controller	0.007s	coverage: 92.3% of  statements
  Done.
  ```

- Run end-to-end tests.

  ```bash
  bash test.sh -e2e
  ```

  This will run all the unit tests along with tests that require real external dependencies like the data layer. Before running end-to-end tests, make sure that our data layer provider, mysql in this case, is up and running.

  ```bash
  docker-compose up
  ```

  Pasting output here for convenience:

  ```bash
  Running unit tests ..
  ok  	github.com/.../wpark/core	0.003s	coverage: 89.5% of statements
  ok  	github.com/.../wpark/service	0.003s	coverage: 90.3% of statements
  ok  	github.com/.../wpark/apiserver	0.006s	coverage: 56.8% of   statements
  ok  	github.com/.../wpark/controller	0.007s	coverage: 92.3% of   statements
  Running e2e tests ..
  ok  	github.com/.../wpark/mysql	0.341s	coverage: 72.0% of statements
  ok  	github.com/.../wpark/e2e	0.111s	coverage: 77.5% of statements
  Done.
  ```


## Run the API Server

To run the api server and listen on our endpoints, run the following command:

- Make sure that our data layer provider, mysql in this case, is up and running.

  ```bash
  docker-compose up
  ```

- Initialize and create our tables in mysql the first time.

  ```
  go run cmd/wpark-api/main.go --initdb
  ```

- Run our api server.  
  Open a new terminal window, and run the following command.

  ```
  go run cmd/wpark-api/main.go 
  ```

  This will start the api server and log all request information in that terminal window.

  By default it will run on http://localhost:11111  
  You can override the default variables by looking at various env variables defined in `config-example.sh` file.  
