# Montirku Backend

## Description
Aplikasi "Montirku" adalah sebuah platform digital yang dirancang untuk memberikan layanan jasa otomotif yang praktis, efisien, dan terpercaya kepada pemilik kendaraan. Aplikasi ini merupakan sebuah solusi inovatif dalam industri jasa otomotif yang bertujuan untuk mengubah cara pemilik kendaraan mengakses layanan perbaikan kendaraan dimana saja dan kapan saja. Pengembangan Backend aplikasi "Montirku" ditulis dalam bahasa pemograman GO (Golang) dan dibangun dengan menggunakan framework Echo, yang merupakan salah satu framework populer dalam pengembangan sistem Golang.

## Fitur Utama
-   To Be Defined

## Requirement & Stack

-   Golang >=1.18
-   MySQL 8.0.33

## Local Installation

1. Clone this project
    ```
    git clone https://github.com/Montirku/montirku-be.git
    ```

2. Copy `.env.example` to `.env`
    ```
    cp .env.example .env
    ```
3. Configure environment variables for database connection
    ```
    DB_CONNECTION=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=montirku_db
    DB_USERNAME=root
    DB_PASSWORD=
    ```

4.  Run the application
    ```
    go run main.go
    ```