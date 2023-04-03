# omdb-api-client

### How To Run the Microservice App
 ```sh
 make run
 ```

Below is an example when app is successfully running:
```sh
2021/11/08 20:41:25 [main][omdbClient] Service successfully started on port:9090
INFO[0005] [Init][GRPC] starting grpc server on :9191    source="server.go:84"
2021/11/08 20:41:25 starting serve on  :9090
```

### Short Description
The Microservice OMDB App implements Clean Architecture. It has 4 layers:
1. `model` containing a set of data structure.
2. `delivery` is the entrypoint of the application.
3. `repository` is the persistence layer.
4. `usecase` containing the business logic.

# API Documentation

### Get List of Movies By Keyword
**URL** : `localhost:9090/omdb-client/search?page=1&keyword=Batman`

**Method** : `GET`

**Response Success Example** :
```json
{
  "header": {
    "process_time": "127.852862ms",
    "code": 200
  },
  "data": {
    "Search": [
      {
        "Title": "Batman Begins",
        "Year": "2005",
        "imdbID": "tt0372784",
        "Type": "movie",
        "Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
      }
    ],
    "totalResults": "470",
    "Response": "True"
  }
}
```


### Get Movie Details By imdbID
**URL** : `localhost:9090/omdb-client/movie/tt0372784`

**Method** : `GET`

**Response Success Example** :
```json
{
  "header": {
    "process_time": "704.563224ms",
    "code": 200
  },
  "data": {
    "Title": "Batman Begins",
    "Year": "2005",
    "Rated": "PG-13",
    "Released": "15 Jun 2005",
    "Runtime": "140 min",
    "Genre": "Action, Adventure",
    "Director": "Christopher Nolan",
    "Writer": "Bob Kane, David S. Goyer, Christopher Nolan",
    "Actors": "Christian Bale, Michael Caine, Ken Watanabe",
    "Plot": "After training with his mentor, Batman begins his fight to free crime-ridden Gotham City from corruption.",
    "Language": "English, Mandarin",
    "Country": "United Kingdom, United States",
    "Awards": "Nominated for 1 Oscar. 13 wins & 79 nominations total",
    "Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
    "Ratings": [
      {
        "Source": "Internet Movie Database",
        "Value": "8.2/10"
      },
      {
        "Source": "Rotten Tomatoes",
        "Value": "84%"
      },
      {
        "Source": "Metacritic",
        "Value": "70/100"
      }
    ],
    "Metascore": "70",
    "imdbRating": "8.2",
    "imdbVotes": "1,362,365",
    "imdbID": "tt0372784",
    "Type": "movie",
    "DVD": "18 Oct 2005",
    "BoxOffice": "$206,852,432",
    "Production": "N/A",
    "Website": "N/A",
    "Response": "True"
  }
}
```