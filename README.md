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

- failed response:


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

- failed response:


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

- failed response:


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

- failed response:


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

- failed response:


### `GET` products
it's a get all product {by userId jwt} endpoint.  This endpoint only can be accessed by `admin` and `user`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

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

- failed response: