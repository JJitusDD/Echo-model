# Model-Service-Echo

## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!


### Structure design

```
|-- cmd
    |-- server
        |-- main.go (entry point for the application)
|-- internal
    |-- app
        |-- app.go (Echo instance and setup)
        |-- middleware (custom middleware for the Echo instaenc)
        |-- routes (definition of all application routes)
    |-- domain
        |-- model (structs for domain objects)
        |-- repository (interfaces for data access)
        |-- service (interfaces for business logic)
        |-- usecase
    |-- infrastructure
        |-- persistence (implementations of repositories)
        |-- utils (utility functions)
|-- pkg (third-party packages)
    |-- logger (implementations of logging lib)
        |-- logger.go
    |-- error (definition errors for the project)
        |-- error.go
```

In this structure **`internal`** directory contains all the code specific to the
application, including the **`app`**, **`domain`**, **`infrastructure`**
packages. The **`cmd`** directory contains the **`main.go`** file, which is the
entry point of the application. The **`pkg`** directory contains third-party
packages used in this application.