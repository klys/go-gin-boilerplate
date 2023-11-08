# Go Gin Boilerplate

This boilerplate comes with folder structure of a MVC framework as following:

* accesors
* controllers
* models
* routers

As explained:

## Accesors

Are Data Access Objects, connection to the respective Database, in this case MySQL

## Controllers

Define the behaviours of the routes, allowing the communication between the accessors and the routes

## Models

Are the data definitions, struct mostly but also may contain specific logic for data manipulation

## Routes

The REST API routes to interface to the outside world