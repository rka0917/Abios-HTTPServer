# Abios-HTTPServer

Sample server middleware application that fetches data from Abios API.  

## Setting up environment

When deploying the docker image from local machine or running the project locally, create a .env file with the following values:                                                                                                                                        
```
ABIOSURL=
ABIOSAPIKEY=
```     

## Launching the application

A dockerfile is included in the project. Build the image with: 

```
docker build -t local-abios-server .
```

Run the container with command: 
```
docker run -e ABIOSURL=<Insert API_URL> -e ABIOSAPIKEY=<Inser API Key> -p 8008:8080 local-abios-server
```
