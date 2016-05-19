# Service Discovery
This sample uses a Go library to communicate with Netflix OSS's _Eureka_ service.

To get the code sample to work, you can start a docker image using the following command:

```
docker run -p 8080:8080 netflixoss/eureka:1.3.1
```

Once this image is running, make sure the Go code reflects the right IP address and port and then you can run the
sample with:

`go run fargo.go`

If you don't already have the fargo library installed in your Go workspace:


`go get github.com/hudl/fargo`
