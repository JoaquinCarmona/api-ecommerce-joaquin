# API E-Commerce Joaquin

This project is a Golang backend application designed for managing products and shopping cart. It employs SQL queries directly for database management and adheres to a clean architecture approach. UUIDs are utilized as unique identifiers to ensure robust database management practices.

## Features

- **SQL Queries**: Direct usage of SQL queries for database management.
- **Clean Architecture**: Adherence to clean architecture principles for better code organization and maintainability.
- **UUID Usage**: Integration of UUIDs as IDs for improved database management practices.

## Setup

1. Clone the repository.
2. Run the `create_tables.sql` script to initialize the database.

## Usage

To utilize the application effectively, consider the following:

- Ensure proper configuration of environment variables for database connectivity.
- Implement necessary endpoints for product and cart management.
- Please find the `api-ecommerce.postman_collection.json` & `ecommerce-api.postman_environment.json` files in order to get the endpoints documentation on postman

## Areas for Improvement

- **Response Handling**: Improve the implementation of responses throughout all services for better consistency and clarity.
- **Repository Refactoring**: Reduce logic complexity within repositories for enhanced readability and maintainability.
- **Validation Implementation**: Integrate a validator to ensure data integrity and validate fields in CRUD operations.
- **Validation updating Product**: If the update endpoint is called without all the fields, the Product is modified completely to null values
- **Domain Separation**: Separate product and cart domains to enhance modularity and scalability.
- **Total Cost Calculation**: Implement logic to calculate the total cost of the cart in the response.
- **Quantity Tracking**: Incorporate functionality to track the quantity of each product in the cart response.
- **Swagger Documentation**: Correctly document API endpoints and structures using Swagger for enhanced usability and understanding.
- **Product delete validation**: Validate if a product is related to a existent cart, soft delete

