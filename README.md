## Getting Started
This is basic implementation of RestFull API based on GO Lang.


Download files in your web server directory and extract them or clone repository using git command:

```
git clone https://github.com/plamenpetrov/go-lang-rest-api.git
```

### How to run the project

Open terminal and navigate to project directory. Then run the following command:

```
go build . 
```

```
 .\webservice.exe
 ```
 
Open your brouser or Rest client (Postman) and you can test one of endpoints:

```
Get list of available users:
http://localhost:3000/users
Body: N/A
```


```
Post request to create new user:
Endpoint: http://localhost:3000/users
Body: {
    "firstname": "username",
    "lastname": "userlastname"
}
```

```
Get data for specific user (in this case ID = 1):
http://localhost:3000/users/1
Body: N/A
```

```
To update user data send PUT request to:
http://localhost:3000/users/1
Body: {
    "id": 1,
    "firstname": "Plamen",
    "lastname": "Petrov"
}
```

```
To Delete user send DELETE request to:
http://localhost:3000/users/1
Body: N/A
```
