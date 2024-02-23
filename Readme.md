
# Savannah Order Management

Savannah Order Management is a REST API that allows you to manage orders and customers for your business. It is built with Go and PostgreSQL. It also features authentication, authorization, and SMS alerts.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Resources](#resources)
## Installation

To install and run this project, you need to have Go, PostgreSQL, and Docker installed on your machine. Then follow these steps:

- Clone this repository:

```bash
git clone https://github.com/MikeMwita/savannah-ordermanagement.git
cd savannah-ordermanagement
```

- Install the required Go packages:

To install the required packages, you can run the following command:

```bash
go get -u
go mod tidy
```

```bash
- Copy the example .env file and fill in the required variables:

```bash
cp .env-example
```


- Build and run the Docker image:

```bash
docker build -t savannah .
docker run -p 5556:5556 savannah
```

- To Run the application locally ,you can run the following command:

```bash
go run main.go
```

- Open http://localhost:5556 in your browser to access the API.

You can find the full list of endpoints and their parameters in the OpenAPI documentation under  the docs directory of this project

## Testing



- Run the tests with the following command:

```bash
go test ./...
```

- Check the coverage report with the following command:

```bash
go test -cover ./...
```



## Resources

Here are some of the resources and references that we used or consulted for this project:

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [OpenID Connect](https://openid.net/connect/)
- [Africa's Talking SMS](https://africastalking.com/sms)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [GitHub Actions](https://github.com/features/actions)
