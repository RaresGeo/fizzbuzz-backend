# Installation
clone the repository and run the following command to install the dependencies
```bash
go mod tidy
go mod vendor
```

# Setting up environment variables
Copy the .env.example file to .env and fill in the required values
```bash
cp .env.example .env
```

# Running the application
```bash
go run src/main.go
```