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

## Build and Run application
Create a pod on your local machine:
```
sudo podman pod create
```

Build and run your container:
```
podman build -t pokedex-web-app:latest .
podman run -d -p 8080:8080 pokdex-web-app:latest
```
Navigate to [localhost:8080/pokemon](http://localhost:8080/pokemon).

## Resource List
| Resource             | Description                                     |
| :---                 | :---                                            |
| GET /v1/pokemon/     | Retrieve a list of all pokemon                  |
| GET /v1/pokemon/:id  | Retrieve the pokemon corresponding to given id  |

## Screenshots of Application Running
<p align="center">
  <img src="/img/pokemon-documentation-image-1.png" />
</p>

<p align="center">
  <img src="/img/pokemon-documentation-image-2.png" />
</p>

## Building/Running Container Locally
```
podman build -t pokedex-web-app:latest .
podman run -p 8080:8080 -it --network host <container name>
```

## Running Docker Compose
```
podman-compose up
```
