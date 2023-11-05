1. mkdir hello-world-temporal
2. cd hello-world-temporal
3. go mod init hello-world-temporal/app
4. go get go.temporal.io/sdk
5. go.temporal.io/sdk/workflow
6. go get github.com/stretchr/testify/mock
7. go get github.com/stretchr/testify/require
8. go mod tidy
9. mkdir start
10. mkdir worker
11. go run worker/main.go
12. go run start/main.go