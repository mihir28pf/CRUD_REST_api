# A simple CRUD REST API in GoLang with Mux & GORM

This repo has code for a sample REST API for `product`. You can create, update, delete and get products.
Take a look at the api.rest file for various APIs and how to invoke these.
The code uses Gorilla Mux for routing requests, GORM as the ORM to access the database,
Viper for Loading configurations, and MySQL as the database provider.
The code uses config.json file to connect to MySQL database.

# What is expected of you?
- Write a Dokerfile to build the app inside docker.
- Write a docker-compose.yml file to run the app locally with MySQL as a dependency inside docker
- Create a GitHub Actions CI workflow to build the docker image and push it to public registry (e.g. dockerhub)
- Record and share the video demo (<3 min) for the above steps 
- Write GitHub Actions CD workflow to deploy this app locally using Vagrant VMs (nice to have)
