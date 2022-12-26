# My Go Proj
Basic Gin Tonic - Rest Api with Controller test example and Dockerfile

## Execution and Build
### Test
```sh
# run all tests
go test ./...
```
### Run Project locally
```sh
go run
```

### Build for Production
```sh
go download
go mod tidy
go build
```

## Extra
### Init Project
```sh
mkdir <project_name>
cd <project_name>
go mod init
```

### Download packages
```sh
go get <my_package>

# replace Package with local package
go mod edit -replace <my_package>=<local_path_to_package>
```

