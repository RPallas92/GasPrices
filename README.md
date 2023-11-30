# Gas Station Finder

Gas Station Finder is an application I created 2 years ago during the winter holidays with the purpose of learning Go. It is designed to find gas stations along a route or between cities, and order them by price. It only works with Spain gas stations but it can be extended to any country.

![A sample UI](gas_prices_ui.png?raw=true "A sample UI")

## Overview

There are several main components:

- **Routing:** Powered by the OpenRouteService API, the application calculates routes between two geographical points.
- **Gas Station Data:** Retrieves gas station information, including prices, from the goverment of Spain [datasource](https://datos.gob.es/es/catalogo/e05068001-precio-de-carburantes-en-las-gasolineras-espanolas). They are returned in a JSON file [like this one](https://sedeaplicaciones.minetur.gob.es/ServiciosRESTCarburantes/PreciosCarburantes/EstacionesTerrestres/).
- **Geocoding:** Converts city names to coordinates using the OpenRouteService's geocoding service.
- **Web API:** Built with the Gin web framework, the API exposes endpoints for finding gas stations based on coordinates or city names.
- **Scheduled Warm-up:** Periodically warms up the gas station data to ensure quick responses during user requests. This is done because the prices JSON returned includes all Spanish gas stations and it is slow to load.
- **KDTree for Efficient Searches:** it uses a KDTree data structure to efficiently search for gas stations close to a given route, reducing search time complexity.


## Project Structure

The project is organized into several packages:

- **gasstations:** Manages gas station data, including a datasource and repository.
- **positions:** Handles geocoding and provides structures for representing coordinates.
- **routes:** Contains the main logic for finding gas stations along a route or between cities.
- **directions:** Integrates with the OpenRouteService API to obtain directions between two points.

## How to Execute

Follow these steps to run the Gas Station Finder:

1. **Build the executable**
```
go build
```

2. **Run the executable**
```
./prices 
```

3. **Call the API**
A gin server will be started on port 8080.


## API Endpoints

### Find Gas Stations by Coordinates:
```
GET /findGasStations/coordinates/{origin}/{destination}
```

### Find Gas Stations by Cities:
```
GET /findGasStations/cities/{origin}/{destination}
```

## Disclaimer
This project was created two years ago with the primary objective of learning Go. The code might be outdated and not be idiomatic.

## License
This project is licensed under the MIT License.

## Author
Gas Station Finder is a project created by Ricardo Pallás Román in 2022 for learning purposes in Go.