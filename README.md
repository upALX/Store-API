# Store-API
This is an api to compose a e-commerce with microsservices.

![giphy](https://github.com/upALX/All-Assets/blob/main/on-mvp.webp)

🎇**This project is on MVP**🎇

---

## Tech stack
![GO](https://img.shields.io/badge/-Go-05122A?style=flat&logo=Go)&nbsp;

## How to use 🫁

**1 - Clone this repo**
```
git@github.com:upALX/Store-API.git
```

**2 - This project use ```.env``` to set the DATABASE variables for good practice, so after clone create an .env file and put this variables:**
```
DB_USER="alx"
DB_PASSWORD="alxroot"
DB_NAME="store"
DB_ROOT_PASSWORD="root"
DB_PORT="3306"
```
*you can change the values, but MySQL run on 3306 port.*

**3 - Install all GO dependencies&modules:**
```
go mod tidy
```

**4 - Run the database:**
``` 
docker-compose up
```

**5 - In a new terminal, access the bash of SQL and set the tables:**
- 5.1: accessing MySQL bash:
  ``` 
  docker-compose exec mysql bash
  ```
- 5.2: accessing MySQL admin (when call password,add the pasword on DB_ROOT_PASSWORD[into .env file]):
  ```
  mysql -uroot -p 
  ```
- 5.3: search the database (DB_NAME), use them and create the tables "products" and "categories":
  ```
    CREATE TABLE `categories` (
        `id` varchar(36) NOT NULL,
        `name` varchar(255) NOT NULL,
        PRIMARY KEY(`id`)
    );
    
    CREATE TABLE `products`(
        `id` varchar(36) NOT NULL,
        `name` varchar(255) NOT NULL,
        `description`  varchar(255) NOT NULL,
        `price` decimal(10,2) NOT NULL,
        `category_id` varchar(36) NOT NULL,
        `image_url` varchar(255) NOT NULL,
        PRIMARY KEY(`id`),
        KEY `fk_products_categories_idx` (`category_id`)
    );
  ```
*if the DATABASE don't exist, just create :p*

**6 - With compose running, open new terminal and run the app(make sure you are in the folder: cmd/catalog):**
```
go run main.go
```

## Make your mark :triangular_flag_on_post:      

**If you have any problems with this app or have an idea that contributes, open a [issue](https://github.com/upALX/Store-API/issues), [pull request](https://github.com/upALX/Store-API/pulls) or find me on [Linkedin](https://www.linkedin.com/in/alxinc/). Don't forget to give the project a star 🌟 :D**

## License :unlock:

This project is under the [MIT license](https://github.com/upALX/Store-API/blob/main/LICENSE).

---

**Developed with 💜 by ME**
