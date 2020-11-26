# Client Tracker - Technical test in Golang and Vue 

Web application to visualize information about customers' transactions. 
* Fetching information from 3 different endpoints that were given, each of them in a different format: Buyer in JSON format, Products in CSV format, and Transaction in No Standard format.
* First loading the Schema of the database with an Alter, then saving the data in Dgraph database using mutations, connecting the data with an Upsert, and finally making queries to retrieve the data a expose it to different endpoints to be consumed from the frontend.
* Then, from the frontend consuming the endpoints and displaying the information in a Dashboard, that allows you to interact with the data visualizing buyers, products, and transactions.

Technical test in Golang a Vuejs for a recruiting process

<img width="200" alt="ClientTracker" src="https://user-images.githubusercontent.com/53787841/100001129-a7483580-2d90-11eb-8274-6fcbdd062ec4.png">

## Architecture

<img width="500" alt="image" src="https://user-images.githubusercontent.com/53787841/100375008-25981800-2fdb-11eb-9e46-eaf8af71862c.png">

## Data model

<img width="300" alt="image" src="https://user-images.githubusercontent.com/53787841/100007872-b502b880-2d9a-11eb-8c52-7c0252d7bdb4.png">

## Project setup - Backend

### First, grab the latest version of Docker.

Then, pull the latest Dgraph version with:

```
docker pull dgraph/dgraph:v20.03.0
```

### Let’s create a folder for storing Dgraph data outside of the container:
```
mkdir -p ~/dgraph
```

### Now, to run Dgraph in Docker, it’s:
#### → Run Dgraph zero
```
docker run -it -p 5080:5080 -p 6080:6080 -p 8080:8080 \
  -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph --name dgraph \
  dgraph/dgraph:v20.03.0 dgraph zero
```
* Port 8000 will be used for the web-UI
* Port 8080 will be used for the RestAPI
* Port 9080 will be used for the grpc API 

#### → In another terminal, now run Dgraph alpha
```
docker exec -it dgraph dgraph alpha --lru_mb 2048 --zero localhost:5080 --whitelist 0.0.0.0/0
```
#### → And in another, run ratel (Dgraph UI)
```
docker exec -it dgraph dgraph-ratel
```

### Load the Dgraph Schema
cd into backend/dgraph_setup/ and then hit:
```
 go run main_schema.go
```

### Then, start the server
cd into backend/api/ and then hit:
```
 go run main.go
```
the web-app will be listening and serving on port 3000

## Project setup - Frontend

cd into the folder technical_test_Go/frontend and install all the packages running:
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

## Author

* [GitHub - Julian Franco Rua](https://github.com/julianfrancor)

* [LinkedIn - Julian Franco Rua](https://www.linkedin.com/in/julianfrancor/)