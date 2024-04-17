# numerator-go-sdk sample project
This sample project demonstrates how to integrate the Numerator SDK into a Go application to leverage feature flags. It provides examples of fetching feature flag values based on specific conditions and handling HTTP requests using the Echo framework.

## How to run this project
1. To install the necessary dependencies, run:
```bash 
go mod tidy
```

Or
```bash 
go get github.com/labstack/echo/v4
go get github.com/numerator-io/sdk-go
```

2. You might need to update the apiKey variable in the main function.

3. Run the application:
```bash 
go run main.go
```

4. Access the following endpoints to interact with the application:
- `/hello`: Returns whether a specific feature flag is enabled.
- `/user-country`: Retrieves the country code based on the provided environment.
- `/expiration-token`: Calculates the expiration time of a token based on user ID and email.