# Create a RESTful server in Go

Source: Linux Format November 2021 Coding Academy page 88-91

## Information

Cover how to develop RESTful API in go and command line clients to access the
RESTful server.

REST(REpresentation State Transfer) is an architecture for designing web
services. REST isn't tied to any OS or system and is not tied to a protocol;
however to implement a RESTful service you need to use a protocol such as HTTP.
REST can work with any data format, but usually means JSON or plain-text over
HTTP. When using JSON  the appropriate Go structures need to be created.

Go is ideal for developing servers and clients for REST APIs because it works
well with JSON data and supports concurrency by design. As a result, all
RESTful servers are concurrent without needing any extra code because the
packages used for serving client requests operate concurrently.
Principles of concurrent data sharing apply and data variables shouldn't share
data carelessly.

2 different builds, first will use standard Go lib and will implement the
following endpoints:

1. `/time`: returns current date and time; mainly used to test that the server
and client can communicate without issues
2. `/insert`: insert new record into the server (requires JSON input)
3. `/list`: retrieve contents of a map with the data to client
4. Fall-through: default handler for all routes that don't match the ones named
above. In the default Go router is associated with the `/` endpoint.

RESTful server will accept all HTTP methods for all endpoints apart from `/list`
which only works with GET. This is bad practice that we will correct with
gorilla/mux. The default Go router requires manual use of code.

## The Go Router

At this point in article, v1/main.go doesn't look too promising, but we will see
what happens when we get to the end of it.

## The Marshall plan

JSON marshalling and un-marshalling, which is required for sending and
receiving JSON. Marshalling is the process of converting a Go structure into a
JSON record. Un-marshalling is the process of converting a JSON record given
as byte slice into a Go structure (usually desired when receiving JSON data via
networks or when loading JSON data from disk files).

`json.Marshal()` and `json.Unmarshal()` are the related methods.

## HTTP Status Codes

According to the HTTP protocol, HTTP method of a request can be defined.
POST is used for creating new resources, GET is used for reading(getting)
existing resources an PUT is used for updating resources.
PUT requests should contain the full and updated version of an existing
resource.
DELETE is used for deleting resources.
PATCH is used for updating existing resources. A PATCH request only contains
the modifications to an existing resource.

HTTP status code 200 means everything went well and the specified action was
executed successfully.
201 means the desired resource was created;
202 means that the request was accepted and is currently being processed
(usually used when an action takes too much time to complete);
301 indicates that the requested resources has been permanently moved - the
new URI should be a part of the response. This is rarely used in RESTful
services because usually you use API versioning.
400 indicates a bad request and that the initial request should be changed
before sending it again;
401 means that the client attempted to access a protected request without
authorization;
403 means that the client doesn't have required permission for accessing a
resource even though the client is properly authorized. In unix terminology it
means that the use doesn't have required privileges to perform an action;
404 means that resource wasn't found;
405 indicates client method isn't permitted by the type of resource;
500 means internal server error

## Other

Go follows a simple rule that states functions, variables, data types, struct
fields, etc. that begin with an uppercase letter are public, whereas if they
begin with a lowercase they are private.

Author (Mihalis T.) did not specify what is considered public and what is
considered private. From what I remember from go play ground "public" means it
can be imported from other files or packages (I don't remember which)
this: <https://tour.golang.org/basics/3>, a name is exported if it begins with a
capital letter. When importing a package, only its exported names can be
referred to. Un-exported names are not accessible from outside the package.
