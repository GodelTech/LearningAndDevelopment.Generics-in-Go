# About 

In the repository we have a few examples of *generics* usage in Go language. 
All those examples had been written and verified using Go1.18 beta1.

Below, under the "Project Structure" section you could read details about each example. 

More information regarding Generics in Golang you can find in the Godel blog. 

# Project Structure

- cmd/comparable 

shows how to use *comparable* constraint in generic functions;

- cmd/concat

shows how to use an interface as a constraint in generic functions;

- cmd/filter

shows how to use *any* in generic functions and how to use generics in your code at all;

- cmd/sum

shows how to use *union types* in your custom interfaces;

- cmd/zeros

an example of what's happening with *empty values* when you use generic functions.

# How to Run

Just like any other CLI Go application: 

```shell
go run cmd/{APP_NAME_HERE}/main.go
```
for example: 
```shell
go run cmd/comparable/main.go
```
