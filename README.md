 Go (Golang) Backend Architecture project with ECHO, MongoDB, JWT Authentication Middleware, SMTP Mail



This project includes a backend service application template written in Go. This application contains various folders and files to follow a clean code structure.

contents

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

There are Dockerfile and docker-compose.yml files required for the application to be run inside the Docker container.


Katkıda Bulunma

Bu proje, her türlü katkıya açıktır. Sorunları bildirme, özellikler önerme veya doğrudan kod katkısında bulunma gibi çeşitli yollarla katkıda bulunabilirsiniz. 
