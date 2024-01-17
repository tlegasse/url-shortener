# URL Shortener - A Go Learning Project
## Introduction
This project is a URL shortener developed in Go, it's intended as a learning tool ratcher than a production level application. It showcases fundamental Go programming concepts, including HTTP request and response handling, database interactions, unit testing and dynamic configuration options.

## Goals and Learning Outcomes
- *Understang Go's HTTP Package*: Gain hands-on experience with creating web servers and handling requests/responses in Go.
- *Database Interaction*: Learn how to integrate and interact with a SQL database in a Go application.
- *Unit Testing*: Develop skills in writing testable code and understand testing strategies in Go.
- *Code Refactoring*: Continuously improve and refactor the codebase for better design and efficiency.

## Features
- Shorten URLs with simple HTTP requests.
- Redirect shortened URLs to their original destination.
- Basic database integration for storing the URL mappings.

## Setup and Usage
Go > 1.16 is required to run this application.
To start, simply clone the repository and run `go run .` in the root directory.

### Configurations
The application will run on port `8080` by default, but this can be changed by providing your own configuration file.
If you wish to provide your own configurations, you can do so by providing an environment variable `APP_ENV_PATH` with the path to your configuration file.
The configuration file should be in a .env format file, and should contain the following key-value pairs (default values shown):
```
BASE_URL=test_base_url
PORT=test_port
URL_SHORTENER_DB_PATH=test_url_shortener_db_path
```

## Contributing
As a learning project, contributions, suggestions and discussions are very much welcome. Feel free to open issues or pull requests to discuss potential improvements, or ways that I can improve.
