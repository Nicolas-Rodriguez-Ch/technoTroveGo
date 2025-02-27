# Technotrove Server

Welcome to the Technotrove2.0! This is the backend API for the Technotrove user interface, built with Go!.

The Technotrove Server is responsible for managing user profiles, project creation, and image uploading. It provides an infrastructure for user creation with password encryption, user authentication using JWT, and project management linked to user profiles.

## Key Features

- User account creation and password encryption
- User authentication with JWT
- Project creation linked to user profiles
- Image uploading for projects and profile pictures

## Tech Stack

- Core: Go, Gin
- Database: GORM, PostgreSQL
- Security: bcrypt, JWT
- Image Handling: Cloudinary
- HTTP Request Logging: Gin
- File Upload: Go

Getting Started

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Nicolas-Rodriguez-Ch/technoTroveGo
   ```

2. **Install the dependencies**:

   ```bash
   cd technotrovego
   go mod download
   ```

3. **Start the development server:**:

   ```bash
   go run ./src
   ```

4. **Start the deployment server:**:

   ```bash
   go build ./src
   ./src.exe
   ```

## Postman Collection

To test out the Technotrove Server API, feel free to use our shared Postman requests. You can find the collection here.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://technotrovego-7132.postman.co/workspace/My-Workspace~2a7615bb-7af4-4c7c-8578-9ce7d14a24c9/collection/13473689-9e065235-20ed-4246-9d7f-b0ca8e6df414?action=share&creator=13473689)

## Environment Variables

The application uses the following environment variables which should be defined in a `.env` file at the root directory:

```bash
DATABASE_URL="<your-database-url>"
PORT=<your-port>
SECRET_KEY="<your-secret-key>"
CLOUDINARY_CLOUD_NAME="<your-cloudinary-cloud-name>"
CLOUDINARY_API_KEY="<your-cloudinary-api-key>"
CLOUDINARY_API_SECRET="<your-cloudinary-api-secret>"
GIN_MODE="<release-or-empty>"
AUTO_MIGRATE="<Reccomended-true-for-development-and-first-deploy"
```

In the root directory, you can find a [.env.example](.env.example) file with the environment variables that are being used throughout the project.

Please replace your-database-url, your-port, your-secret-key, your-cloudinary-cloud-name, your-cloudinary-api-key, and your-cloudinary-api-secret with your actual data.

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file for details.
