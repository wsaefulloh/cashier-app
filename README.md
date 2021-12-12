<h1 align="center">Cashier App Backend</h1>
<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/2560px-Go_Logo_Blue.svg.png" width="400px" alt="Golang.jpg" /></p>
<p align="center">
    <a href="https://golang.org/" target="blank">More about Golang</a>
</p>

## Built With

[![Golang](https://img.shields.io/badge/Golang-4.x-blue.svg?style=rounded-square)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v.13.3-blue.svg?style=rounded-square)](https://www.postgresql.org/)

## Description
This backend application is used by the user to record incoming orders and manage products. In this application, users can display products, add products, delete products, edit products, and calculate the total price of an order. This application is built with Golang using the gorilla/mux package for routing. The databases used in this application are PostgreSQL and Redis for the caching process.

## Architecture

<p align="center"><img src="https://res.cloudinary.com/dyli6i0pw/image/upload/v1638731396/Arsitektur_Backend_Golang_jxwl2r.png" width="700px" alt="Golang.jpg" /></p>

## Feature
- Authentication and Authorization
- JWT Web Token
- CRUD Product
- CRUD Category
- CRUD History
- Solid Principle
- Search Product Name
- Sort Product Name, Category, Date Update, and Price order by ASC or DESC

## Installation Steps

1. Install the Golang and GO Environment

```bash
https://golang.org/doc/install
```

2. Clone the repository

```bash
git clone git@github.com:wsaefulloh/cashier-backend.git (go get)
```

3. Run the app

```bash
go run main.go
```

4. You are all set!

```bash
View the website at: http://localhost:8080
```

## End Point

You can see all the end point [here](https://documenter.getpostman.com/view/16508598/UVC9i6As)
