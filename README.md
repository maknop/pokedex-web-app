# Pokédex Web App - Pokémon Go Inspired!

<p align="center">
  <img src="/img/pokeball.png" width="100" height="100" />
</p>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/maknop/pokedex-web-app)
[![Netlify Status](https://api.netlify.com/api/v1/badges/5c782812-04fe-4899-9632-e01d60c92921/deploy-status)](https://app.netlify.com/sites/dainty-griffin-069603/deploys)

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
Navigate to [localhost:8080/pokemon](http://localhost:8080/pokemon).

## Screenshots of Application Running
<p align="center">
  <img src="/img/pokemon-documentation-image-1.png" />
</p>

<p align="center">
  <img src="/img/pokemon-documentation-image-2.png" />
</p>
