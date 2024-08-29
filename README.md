# PassGen - CLI Password Manager

PassGen is a secure command-line password manager built in Go that allows users to store and retrieve passwords with encryption.

## Features

- User authentication (signup/login)
- Secure password storage with AES encryption
- Password retrieval using keys
- List all stored password keys
- JWT-based session management
- PostgreSQL database backend

## Prerequisites

- Go 1.21.7 or higher
- PostgreSQL database
- Environment variables configured

## Installation

1. Clone the repository:
```bash
git clone https://github.com/abdullahnettoor/pass-gen.git
cd pass-gen
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `dev.env` file in the root directory with the following variables:
```env
DB_CONNECTION_URI=your_postgres_connection_string
JWT_SECRET=your_jwt_secret_key
CIPHER_SECRET=your_32_byte_cipher_key
CONFIG_PATH=.pass-gen
CONFIG_FILE_PATH=config.json
```

4. Build the application:
```bash
go build
```

## Usage

### User Management

1. Create a new account:
```bash
./pass-gen signup
```

2. Login to your account:
```bash
./pass-gen login
```

### Password Management

1. Save a new password:
```bash
./pass-gen save
```

2. Retrieve a password using its key:
```bash
./pass-gen fetch --key YOUR_KEY
```

3. List all stored password keys:
```bash
./pass-gen keys
```

## Security Features

- Password hashing using bcrypt
- AES encryption for stored passwords
- JWT-based authentication
- Secure password input (hidden input)
- Database-level unique constraints

## Project Structure

```
.
├── app/
│   ├── config/     # Configuration management
│   ├── db/         # Database models and connection
│   ├── models/     # Request/Response models
│   ├── pkg/        # Encryption utilities
│   ├── repo/       # Database operations
│   ├── usecase/    # Business logic
│   └── utils/      # Helper functions
├── cmd/            # CLI commands
└── main.go         # Application entry point
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Author

Abdullah Nettoor - [abdullahnettoor@gmail.com](mailto:abdullahnettoor@gmail.com)
