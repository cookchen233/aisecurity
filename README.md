# AI Security

AI Security is a project developed in Go, using the Gin framework for handling HTTP requests and the Ent framework for database operations. It also uses the Qodana linter for code analysis.

## Project Structure

The project is organized into several packages:

- `middlewares`: Contains middleware functions for request logging and error recovery.
- `routes`: Defines the HTTP routes for the application.
- `utils`: Contains utility functions and configurations for logging and database operations.

## Setup

To run the project, follow these steps:

1. Load the environment variables from the `.env` file.
2. Initialize the logger.
3. Set up the session store.
4. Register the middlewares.
5. Open the database connection and run migrations.
6. Register the routes.
7. Start the server.

## Database

The project uses PostgreSQL as its database. Data is imported using the `COPY` command.

## Code Analysis

The project uses the Qodana linter for code analysis. The configuration for Qodana is defined in the `qodana.yaml` file.

## Contributing

Contributions are welcome. Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
