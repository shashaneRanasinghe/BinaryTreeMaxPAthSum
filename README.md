# Binary Tree Max Path Sum API

This is a simple  API that that accepts a Binary Tree
and returns its max path sum.

### How to run the application

- Clone the repository
- Navigate to the project directory: `cd BinaryTreeMaxPathSum`

#### Running using Docker

- run `docker build -t {image name} .` to build
- run `docker run -dp 8001:8001 {image name}` to run
- To stop, run `docker stop {container name}` 


#### Running Locally

- run `go build -o {binaryName} .\cmd\app\main.go` to build the binary
- run `.\{binary name}` to run the binary (eg .\binaryTreeMaxPathSum)

### Testing

- run `go test -v ./...` to run the unit tests
- run `go test -v ./... -coverprofile=c.out` to create the cover profile
- run `go tool cover -html c.out` to see the
  coverage of the unit tests

#### Note
if the c.out file gets created without the out extension
run `go tool cover -html c`

## Endpoints

### Calculate the Max Path Sum

This Endpoint Gives the Max Path Sum of the Given Binary
Tree

#### Request

    curl --location 'http://localhost:8001/binaryTree/findMaxPathSum' \
    --header 'Content-Type: application/json' \
    --data '{
      "tree": {
        "nodes": [
          {
            "id": "1",
            "left": "2",
            "right": "3",
            "value": 1
          },
          {
            "id": "3",
            "left": "6",
            "right": "7",
            "value": 3
          },
          {
            "id": "7",
            "left": null,
            "right": null,
            "value": 7
          },
          {
            "id": "6",
            "left": null,
            "right": null,
            "value": 6
          },
          {
            "id": "2",
            "left": "4",
            "right": "5",
            "value": 2
          },
          {
            "id": "5",
            "left": null,
            "right": null,
            "value": 5
          },
          {
            "id": "4",
            "left": null,
            "right": null,
            "value": 4
          }
        ],
        "root": "1"
      }
    }'

#### Request Body

    {
      "tree": {
          "nodes": [
              {
                  "id": "1",
                  "left": "2",
                  "right": "3",
                  "value": 1
              },
              {
                  "id": "3",
                  "left": "6",
                  "right": "7",
                  "value": 3
              },
              {
                  "id": "7",
                  "left": null,
                  "right": null,
                  "value": 7
              },
              {
                  "id": "6",
                  "left": null,
                  "right": null,
                  "value": 6
              },
              {
                  "id": "2",
                  "left": "4",
                  "right": "5",
                  "value": 2
              },
              {
                  "id": "5",
                  "left": null,
                  "right": null,
                  "value": 5
              },
              {
                  "id": "4",
                  "left": null,
                  "right": null,
                  "value": 4
              }
          ],
          "root": "1"
      }
    }

#### Response

    {
      "status": "Success",
      "data": {
          "maxPathSum": 18
      },
      "message": "Max Path Sum Calculated Successfully"
    }


