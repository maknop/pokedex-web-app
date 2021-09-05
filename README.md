# Pokemon Data API

<p align="center">
  <img src="/img/pokeball.png" width="100" height="100" />
</p>

## Purpose
The purpose of this project is to create a REST API that interacts with  
the [PokéAPI](https://pokeapi.co/) website and retrieves the necessary data which will be  
displayed on the client-side. This project will be built using React.js and  
Node.js.

## High-Level Overview
<p align="center">
  <img src="/img/sequence-diagram.jpg" />
</p>

## Resource List
| Resource             | Description                                     |
| :---                 | :---                                            |
| GET /v1/pokemon/     | Retrieve a list of all pokemon                  |
| GET /v1/pokemon/:id  | Retrieve the pokemon corresponding to given id  |

## Run application
```
go run server.go
```
Navigate to [localhost:8080/pokemon](http://localhost:8080).

## Testing
```
go test
```

