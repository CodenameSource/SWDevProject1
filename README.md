# ShoppyList

## Overview

This project is a web application designed for storing Amazon products in a list, with an integrated microservice for scraping their prices. The application provides a set of APIs to add, retrieve, update, and remove items from the stored list. Additionally, it leverages a Kafka-based microservice to periodically refresh and update the prices of the stored Amazon products.

## Specifications by user stories

![image](https://github.com/CodenameSource/SWDevProject1/assets/65411104/11ed3c89-0296-451e-9e3f-e37fa8ced0b1)

## Features

- **Add Item**: Allows users to add new Amazon products to the system by providing the product's URL.

- **Get Items**: Retrieves a list of all stored Amazon products with their details, including URLs and prices.

- **Update Item Price**: Updates the price of a specific Amazon product by providing the product's URL. This triggers a refresh event using the Kafka microservice.

- **Update Prices of All Items**: Initiates a process to update the prices of all stored Amazon products. This involves sending refresh events for each product and listening for updated prices using the Kafka microservice.

- **Remove Item**: Removes a specific Amazon product from the system by providing the product's URL.

## Project Structure

The project consists of the following components:

- **Web Application**: The core functionality is implemented using Go and the Gorilla Mux router. The application interacts with a database to store and retrieve Amazon product information.

- **Price checker Microservice**: A microservice responsible for handling refresh events and updating product prices. It uses the Kafka messaging system for event-driven communication.

- **MYSQL Database**: Used to store the products and their prices

- **Kafka messenger service**: Used for communication between the microservices

## Setup

To run the project you need to have docker and docker compose installed on your system. The project can be started with `docker compose up`

## API Endpoints

1. **Add Item**
   - Endpoint: `/api/addItem`
   - Method: `POST`
   - Parameters: JSON payload containing the new item details.

2. **Get Items**
   - Endpoint: `/api/getItems`
   - Method: `GET`
   - Returns: List of Amazon products with details.

3. **Update Item Price**
   - Endpoint: `/api/updatePrice`
   - Method: `GET`
   - Parameters: URL query parameter (`url`) specifying the product's URL.

4. **Update Prices of All Items**
   - Endpoint: `/api/updatePrices`
   - Method: `GET`
   - Triggers a process to update prices for all stored Amazon products.

5. **Remove Item**
   - Endpoint: `/api/removeItem`
   - Method: `DELETE`
   - Parameters: URL query parameter (`url`) specifying the product's URL.

## Documentation

- API documentation is available using Swagger at `http://localhost:8080/swagger/index.html`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
