# Go-Home-Assignment

A Golang home assignment I received as part of an interview process

## How to build and execute

    - Download and install go version 1.15 from https://golang.org/
    - Clone this repo.
    - Go to root project folder.
    - run command: "go build ./cmd/validator".
    - On the root project folder, an executable file called "validator.exe" should be created. Run it to execute the program.

## Sending requests to the endpoint

The server listens to requests on port 8080 and we currently support the endpoint "/validate-url" for POST method.
So when running the project locally, POST requests can be sent to http://localhost:8080/validate-url

The body of the request should be json of this form:

    {
        "domain": "yourDomain",
        "path": "yourPath"
    }

If an item with the specified domain and path exists in the data store, the response would be of this form:

    {
        "location": "https://<yourDomain>/<yourPath>"
    }

otherwise, the location field would be an empty string

URLs currently stored in our inmemory data store are:

    {Domain: "ynte.co.il", Path: "/home/0,7340,L-8,00.html"},
    {Domain: "walla.co.il", Path: "/index.html"},
    {Domain: "google.com", Path: "/index.html"}
