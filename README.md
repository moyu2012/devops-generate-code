#just for devops config agent

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o devops-generate-code main.go