# Weather Forecast App

A simple Go application that provides weather forecasts for cities using OpenStreetMap and a Redis cache for caching weather data.

## Table of Contents

- [Weather Forecast App](#weather-forecast-app)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Commands](#commands)
- [Update based on 2h+ progress](#update-based-on-2h-progress)

## About

The Weather Forecast App is a Go application that allows users to retrieve weather forecasts for cities using the OpenStreetMap geocoding service. It leverages in-memory and file based as a caching mechanism to store previously fetched weather data, intented to improve response times and reducing the load on external weather services.

## Getting Started

To get started with the Weather Forecast App, follow the instructions below.

### Prerequisites

Before you begin, ensure you have the following prerequisites installed:

- Go 1.19

### Installation

Locally 
```
cd weather-forecast-app/cmd/weather_app
go build
./main
```

Start Image in Docker (Container)
```
docker build -t weather-forecast-app .
docker run -d -p 8080:8080 --name weather-forecast-container weather-forecast-app
// Watch the logs ->
docker logs -f weather-forecast-container
```

Access Your Application: Your Weather Forecast App should now be accessible at `http://localhost:8080` in your web browser or through HTTP requests.

Stop+Remove Image in Docker (Container)
```
docker stop weather-forecast-container
docker rm weather-forecast-container
```

## Commands

If you are running this locally uoi can do the following: 
Assuming you have installed curl - 
```
curl localhost:<PORT>/weather?city=los%20angels,%20new%20york,chicago
```

Alternatively you can make a get request using postman - `localhost:<PORT>/weather?city=los%20angels,%20new%20york,chicago`


# Update based on 2h+ progress
- I had to understand the task (5-10mins) and manually navitage and understand the json keys (that didnt take long ~ max 15mins)
- I started to create a ideal go folder structure ~ 10mins
- Created a server and main ~ 15mins
- Handler too much longer this came up to an 1h
- created cache 20mins
- Added tests for config BDD - 20mins
- Added TDD tests with Mocks - upto 1h
- Dockerfile - 5mins
- Readme - 15mins

Overall the to complete the task it was possible within 2h, however, to refactor the code, apply some refined tests, debug, manually tests and so on... This exceeded 2h came roughly upto ~4.5h. It was fun thats why continued past the 2h.