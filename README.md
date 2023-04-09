# ORM JWT AUTH
it's a jwt authorization implementation. Here's a [API flow](https://import.cdn.thinkific.com/236035/courses/1973834/challengedrawio21-220816-175740.pdf) task details:

```text
Buatlah Rest API product (create, read, update, delete) dengan fitur login dan register, serta memiliki 3 fitur middleware antara lain :
- Authentication
- Authorization multi level user
- Authorization access product by id
```
**Notes** : buatlah authentication dengan JWT token golang, lalu gunakan token tersebut untuk setiap hit Rest API product.

# Documentation

### `POST` users/register
it's a register user endpoint. This endpoint can be accessed by everyone.

**Request Body** :
```json
{
    "fullname" : "string",
    "email" : "string", # must be a valid email
    "password" : "string",
	"role" : "string" #must be a valid role, but omitempty
}
```
**Notes** : On this endpoint, `password` will be hashed using `crypto` library 

**Response Body** :
- success response:
![Screenshot 2023-04-09 072413](https://user-images.githubusercontent.com/37493831/230748989-ffd325eb-9014-4578-a26d-1d3075d6d4a4.png)

- failed response:
![Screenshot 2023-04-09 072256](https://user-images.githubusercontent.com/37493831/230748988-a33acdc4-0efa-40d7-a4ad-035f652d90e2.png)


### `POST` users/login
it's a login user endpoint. This endpoint can be accessed by everyone.

**Request Body** :
```json
{
    "email" : "string", # must be a valid email
    "password" : "string",
}
```

**Response Body** :
- success response:
![Screenshot 2023-04-09 072545](https://user-images.githubusercontent.com/37493831/230748968-11507120-d732-4495-99d8-5887ef887d7d.png)

- failed response:
![Screenshot 2023-04-09 072505](https://user-images.githubusercontent.com/37493831/230748967-2c2706c7-d504-4953-99e1-908de200dba4.png)


### `POST` products
it's a create product endpoint. This endpoint only can be accessed by `admin` and `user`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "title" : "string",
    "description" : "string",
}
```

**Response Body** :
- success response:
![Screenshot 2023-04-09 072704](https://user-images.githubusercontent.com/37493831/230748969-de49f0df-7c2d-4a27-be36-f84bc527810c.png)

- failed response:
![Screenshot 2023-04-09 074246](https://user-images.githubusercontent.com/37493831/230748983-e3cc0402-bee9-4fb2-88ae-a9cffecf7663.png)


### `PUT` products/:productId
it's a update product endpoint. This endpoint only can be accessed by `admin`. 

**Params**
- productId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "title" : "string",
    "description" : "string",
}
```

**Response Body** :
- success response:
![Screenshot 2023-04-09 073327](https://user-images.githubusercontent.com/37493831/230748976-6cafc1d4-e3bf-4a49-84cc-9b26eec7a26c.png)

- failed response:
![Screenshot 2023-04-09 072842](https://user-images.githubusercontent.com/37493831/230748970-6982c703-8345-4c02-8c1f-e8d37cf600e7.png)
![Screenshot 2023-04-09 074152](https://user-images.githubusercontent.com/37493831/230748982-bd6b90ad-f5a6-42f7-9db7-b251a7e5af71.png)


### `DELETE` products/:productId
it's a delete product endpoint. This endpoint only can be accessed by `admin`.

**Params**
- productId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![Screenshot 2023-04-09 073730](https://user-images.githubusercontent.com/37493831/230748978-3a100680-8107-43c3-8aee-05f76f59588c.png)

- failed response:
![Screenshot 2023-04-09 073806](https://user-images.githubusercontent.com/37493831/230748981-0fee42e9-3403-4556-9f7f-c82bac5cb3c3.png)


### `GET` products
it's a get all product {by userId jwt} endpoint.  This endpoint only can be accessed by `admin` and `user`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![Screenshot 2023-04-09 073044](https://user-images.githubusercontent.com/37493831/230748973-5fa0b292-2a30-4937-86d6-dc99c35910bf.png)
![Screenshot 2023-04-09 073235](https://user-images.githubusercontent.com/37493831/230748974-641fd6c6-b6d4-4b9d-b6be-b54fb0e0a1e5.png)

- failed response:
![Screenshot 2023-04-09 074246](https://user-images.githubusercontent.com/37493831/230748983-e3cc0402-bee9-4fb2-88ae-a9cffecf7663.png)

### `GET` products/:productId
it's a get product by productId endpoint. This endpoint only can be accessed by `admin` and `user`.

**Params**
- productId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![Screenshot 2023-04-09 072926](https://user-images.githubusercontent.com/37493831/230748971-5efa5449-727f-4f20-aaa7-507341cca207.png)

- failed response:
![Screenshot 2023-04-09 073000](https://user-images.githubusercontent.com/37493831/230748972-4496e015-22cb-4810-8de9-e3dca9d9c121.png)