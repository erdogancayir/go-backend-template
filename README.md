 Go (Golang) Backend Architecture project with ECHO, MongoDB, JWT Authentication Middleware, SMTP Mail


### Description

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) This project includes a backend service application template written in Go. This application contains various folders and files to follow a clean code structure. This project is designed to create a clean and understandable backend service application. This structure makes the code easier to read and maintain.

#### Accessing Data via Terminal
  ##### Obtain MongoDB IP address
```docker inspect <mongodbContainer> | grep IPAddress ``` <br>
```mongosh mongodb://root:example@<ipAddress>:27017```

### Contents

      api: These ledgers contain all routing and control operations.
          controller: These folders contain control functions for various API endpoints.
          route: These folders define the API's routes.
      bootstrap: These folders contain files for executing and configuring.
          app.go: Here are the words for the app's download-related actions.
          database.go: contains recipients and recipients.
          env.go: Contains the application's env slots.
      cmd: These folders contain run entry points.
          main.go: It is the entry point of the application. Starts the execution of the application.
      domain: This folder contains the runtime base objects and the data types used.
      internal: This folder contains files for surveillance internal surveillance.
          tokenutil: These ledgers contain files for token transactions.
      middleware: This folder contains files for middleware components.
          jwt_auth_middleware.go: Contains the JWT authentication middleware.
      repository: These folders contain files belonging to storage operations.
      usecase: This folder contains files containing detection clouds and job outputs.

### Features<br>

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)<br>
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white): MongoDB is a source-available, cross-platform, document-oriented database program.<br>
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)<br>
![Gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)<br>
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)<br>
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens): JWTs are a method for securely transmitting information between parties as a JSON object.<br>
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)<br>
There are Dockerfile and docker-compose.yml files required for the application to be run inside the Docker container.

### Architecture Layers of the project
 1. Router
 2. Controller
 3. Usecase
 4. Repository
 5. Domain

### Contributing

This project is open to all kinds of contributions. You can contribute in a variety of ways, including reporting issues, suggesting features, or contributing code directly.
