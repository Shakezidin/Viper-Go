# Viper-Go Sample Project

This is a sample project demonstrating how to use the Viper configuration library with Go. It includes a basic web application with user signup and login functionality.

## Getting Started

To run this project, you'll need Go and a PostgreSQL database. Here's how to set it up:

1. Clone this repository.
2. Update the database connection details in the `config/test.json` file.
3. Run the project with `go run main.go`.

## Endpoints

- POST `/user/signup`: Create a new user.
- POST `/user/login`: User login.

## Dependencies

- Viper for configuration management.
- Gin for web server.
- GORM for database interaction.

## License

This project is open-source and available under the MIT License.

Feel free to explore and learn from this project. If you have any questions or need further assistance, please let me know.
