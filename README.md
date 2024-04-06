# KINEXON Container Runtime (Coding Challenge)

## What is this about?
This challenge is part of the KINEXON interview process for full stack developers. It gives you the opportunity to show
us your skills, interests, motivation and how you work in general.


## What's included
This repository contains a skeleton for a Container Runtime Service that is responsible for visualizing and managing the 
lifecycle of docker containers. The service is split into two parts: a backend written in Go and a frontend based on Vue.

The backend connects to your local docker daemon through a UNIX socket and exposes runtime information through a REST API. 
Initially, it only contains a single endpoint "GET /info" that returns basic information about the docker daemon.

The frontend is a simple Vue application that connects to the backend and is responsible for visualizing the information 
returned by the API. For now, it only contains a single page that displays some information about the docker daemon and 
the host in a table.


## How to run the code

### Prerequisites
The only strict dependency for running this code are docker and docker-compose.
This code was successfully tested on an Ubuntu 20.04 machine with docker v25.0.3 and docker-compose v2.24.5 installed.

Optional requirements are node/npm/yarn if you want to install additional frontend packages and go/golang (v1.22) if
you want to run the backend code outside of docker or run backend unit tests locally.

### Running the code

The code can be run using the following command:
```bash
docker compose up --build
```

This will start the backend and frontend services and make them available at http://localhost:5000

The backend API will automatically be proxied through the Vite development server (see vite.config.js) and is
available at http://localhost:5000/server

If you want to make use of hot reloading for the frontend and automatically rebuild the backend when code
changes are detected you can use the following command:
```bash
# Requires docker compose version 2.22 or later
docker compose watch

# Run in a different terminal to view logs
docker compose logs -f
```


## What we expect from you
You are invited to extend the functionality of the application by adding the following features:

#### Refactoring of initial frontend
The current implementation of the frontend only consists of a single page that displays some information about
the docker daemon in a basic unstyled table. Take your time to refactor the page to make it more user-friendly and 
visually appealing. Also add basic layout elements (menu bar, header, footer etc.) that will be required for
the other features.

#### Container List
Add a new page to the frontend that displays a list of all running containers on the host (think of "docker ps"). For 
each container it should show at least the following properties: container id, name, image, status, and creation date.
It should be possible to refresh the list of containers without reloading the page, and the list should support 
filtering the containers by name or image.

If you like, you can also add additional features like restarting, stopping, or removing containers, but it's also fine
to just display the list of containers.

The backend should be extended with new API endpoints to complement the functionality of the frontend.

#### Container Resource Stats
Add a new page or modal that displays resource usage statistics for a specific container ("docker stats <container-id>"). 
The page should show at least the following information: CPU usage, memory usage, and network I/O. 
The statistics should be updated automatically and in real-time. Think of ways how to best visualize the data and how
to update the data in an efficient way.


### General hints
- The code is not particularly well-structured so please feel free to refactor and improve it as you see fit
- We would like to see an intuitive and user-friendly UI that also looks great. If you like to use a CSS framework 
  or UI components to achieve this, feel free to do so.
- To fetch information about running containers etc. you will need to use the Docker API which is documented here:
  https://docs.docker.com/engine/api/v1.44/


### What to deliver
Your submission should include the entire source code of your solution consisting of frontend and backend. 
Please make sure to include all necessary files and resources. It should be possible to build and run your solution 
locally using only "docker compose up". No other build and install steps should be necessary.

The challenge is not time-boxed. You can invest as much or as little time as you want and need to.

If you have any questions about the problem statement, please do not hesitate to contact us at jochen.hartl@kinexon.com.

Looking forward to your submission! Good luck and have fun! 🚀
