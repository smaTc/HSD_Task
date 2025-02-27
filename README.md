# HSD Task

## Description
The project implements a small REST API that allow managing tasks. Implemented in Go with Gin and GORM.

## Packages
- `main` - Main package that contains the main function and REST API configuration using GIN. Forwards the task operations to the `hsdtask` package after starting the server and handling the incoming requests.
- `hsdtask` - Contains the code that handles the database, JSON operations and API responses for the tasks.

## Compile & Run
1. [Install](https://go.dev/doc/install) Go for your platform
1. Clone the Git repository and navigate inside it
1. Execute `go run main.go`

## Examples & Testing
To test the tool via its REST API you can run the postman collection provided in the file `test-hsd-task-postman.json`. Import the collection into postman to run the API calls.
There are 8 calls doing the following:
1. Create Task
1. Get Task by ID
1. Create Task 2
1. Get all Tasks
1. Update Task
1. Get all Tasks again
1. Delete Task
1. Get all tasks last time

