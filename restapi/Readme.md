<!-- server User -->
Created a simple API server using Gorilla Mux

<!-- About Apis -->
I tried creating a simple in-memory store for storing book information. Accomplished basic CRUD operations on books store using Go

<!-- Total APIS -->
Basic store has 3 predefined books with ids: 1, 2, 3, and we can perform CRUD operations on this store.

Below are the endpoints for all the deployed operations.

<!-- gives all the available books -->
GET: https://books-store-312503.uc.r.appspot.com/api/books

<!-- add books to existing list -->
POST: https://books-store-312503.uc.r.appspot.com/api/books/ , add fileds in body as shown below
Body: {
    "Isbn": "8888",
    "Title": "5",
    "Author": {
        "FirstName": "sai",
        "LastName": "duddukuri"
    }
}

<!-- Delete book with id 1 from data -->
DELETE: https://books-store-312503.uc.r.appspot.com/api/books/1

<!-- update book with id 1-->
PUT: https://books-store-312503.uc.r.appspot.com/api/books/1, use below body to update data
Body:{
 Body: {
    "Isbn": "8888",
    "Title": "q5",
    "Author": {
        "FirstName": "sai",
        "LastName": "duddukuri"
    }
}

<!-- Testing -->
Haven't added any access controls to the backend. Thus can be tested with Postman.
