# Gin Gonic RESTful API: Security Matrix Calculation with Goroutines, Channels, GORM, and Wire for IoC

The single endpoint of this Gin-Gonic RESTful API automates the recalculation and storage of security access rules between users, enforced by other system components, in a matrix format. By leveraging bitwise operations for rule condensation and employing worker groups with Goroutines and Channels for rule calculation/generation, the endpoint optimizes access control handling, ensuring efficient performance and best practices in data persistence.

## Features
- **Gin-Gonic RESTful API**: Employs Gin-Gonic, a high-performance web framework for Go, to handle HTTP requests robustly and efficiently. Gin-Gonic simplifies routing, middleware management, and request handling, making the `/CalculateGoSecurityMatrix` endpoint scalable and reliable. More information at https://gin-gonic.com.

- **Dependency Injection with Wire**: Implements dependency injection and inversion of control using Wire. Wire automates the setup of dependencies, reducing boilerplate code and ensuring type safety during the initialization of application components. More information at https://pkg.go.dev/github.com/google/wire & https://github.com/google/wire/blob/main/docs/guide.md.

- **Bitwise Operations for Access Control**: Condenses diverse rules for the same relationship into a single access value using bitwise operations for minimizing storage overhead and enhancing runtime efficiency.

- **Concurrency with Worker Groups**: Implements worker groups with Goroutines and Channels for concurrent rule calculation. This approach maximizes CPU utilization, optimizing performance during rule generation. More information at https://go.dev/tour/concurrency/1 & https://go.dev/blog/pipelines

- **Transaction-based Data Persistence**: Ensures data integrity and consistency by performing operations within a single transaction. This includes deleting the existing security matrix, calculating the new one, and persisting it atomically.

- **GORM for Database Interaction**: Employs GORM for CRUD operations, batch creation, and query control, ensuring reliable data management. More information at https://gorm.io.


## Dependency Management
### Installing Dependencies
Dependencies for this project are managed using Go modules. After cloning the repo, run the following command to install all required dependencies and create a `vendor` directory:
```
go mod tidy
go mod vendor
```
- `go mod tidy`: Ensures `go.mod` and `go.sum` files are up to date by adding any missing module requirements and removing any unnecessary ones.
- `go mod vendor`: Copies all dependencies into a `vendor` directory within the project (for version control). This is useful for ensuring that all dependencies are available locally, which can help with reproducible and offline builds.

### Updating Dependencies
To update the dependencies to their latest versions, run:
```
go get -u ./...
go mod tidy
go mod vendor
```
- `go get -u ./...`: Updates all dependencies to their latest minor or patch versions.
- `go mod tidy`: Cleans up the `go.mod` and `go.sum` files.
- `go mod vendor`: Updates the `vendor` directory with the latest versions of the dependencies.