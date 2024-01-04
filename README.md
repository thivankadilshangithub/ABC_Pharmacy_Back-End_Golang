# Project name --> ABC Pharmacy System 

## Table of Contents
- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Folder Structure](#folder-structure)
- [Frontend](#frontend)
- [Backend](#backend)
- [Database](#database)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)

## Overview
This project is full-stack web system for the ABC Pharmcy.
Utilized PostgreSQL for the database, Golang for the backend, and ReactJS for the frontend.

Developed the fuctinalities to managing items (add, edit, delete) and invoices.

## Prerequisites

Please ensure the below requirements.

- [Node.js](https://nodejs.org/)
- [Golang](https://golang.org/)
- [Database (PostgreSQL)](https://www.postgresql.org/)
- [VScode](https://code.visualstudio.com)

## Getting Started

1. Download the zip folder (ABC_Pharmacy_Thivanka) and locate any suitable location to run the Golang project.
(I have used the github.com folder in the go folder)

2. Unzip the folder and run the frontend and backend separately.

## Folder Structure

Front End - client folder
Back End - server folder

## Front End

Run the frontend folder

--> cd client
--> npm start

## Back End

Run the Backend folder

--> cd server
--> go run main.go

## Database

Postgres Database

Database Name - ABC_Pharmacy_DB

tables name - items and invoices

## API Endpoints

create item - http://127.0.0.1:8081/api/items (POST methosd)
show items - http://127.0.0.1:8081/api/items (GET method)
edit item - http://127.0.0.1:8081/api/items/id (PATCH method)
delete item - http://127.0.0.1:8081/api/items/id (DELETE method)

create invoice - http://127.0.0.1:8081/api/invoices (POST methosd)
show invoices - http://127.0.0.1:8081/api/invoices (GET method)
edit invoice - http://127.0.0.1:8081/api/invoices/id (PATCH method)
delete invoice - http://127.0.0.1:8081/api/invoices/id (DELETE method)

## Testing

Postman API testing
