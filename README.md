# GOBLOG
This is a simple **golang** blog application. That has authentication system.
Part of the challenge of creating same application with same features in **laravel**, **ruby on rails**, **django**, **dotnet core**, **spring**, **golang** and **nodejs**. In the beginning we'll be using go's built in features as much as possible to better understand what go has to offer for web development out of the box.

## Tasks
* ~~TODO: HTTP server~~
* ~~TODO: HTTP request handling and routing~~
* ~~TODO: Template handing~~
* ~~TODO: Static resource handling~~
* ~~TODO: MySQL database integration~~
* ~~TODO: Better template management~~
* ~~TODO: Better error handling~~
* TODO: Better 404 error handling
* ~~TODO: Authentication~~
* ~~TODO: Handle login~~
* ~~TODO: Handle logout~~
* ~~TODO: User create~~
* TODO: User update
* TODO: User delete
* TODO: Post create
* TODO: Post update
* TODO: Post delete
* TODO: Post by user
* TODO: Highlight js code syntax highlight for rich tex editor code type
* TODO: Rich text editor for post
* TODO: Post image and user profile picture

## Content
* go version go1.14.2 linux/amd64

## Dependency
* Database: MySQL
* Database driver: github.com/go-sql-driver/mysql
* http router: github.com/julienschmidt/httprouter `because default handler doesn't support dynamic routing.`
* UI: Bootstrap and its dependencies.

## Instruction
#### Database setup:
* Login to mysql `mysql -u root -p`
* Create database `CREATE DATABASE dbname`
* Import sql data from `data/data.sql` into the database `mysql -u root -p dbname < data/data.sql`
**************************
#### Run server:
* Dev server: **go run main.go** and visit to 127.0.0.1:2000 on the browser.
