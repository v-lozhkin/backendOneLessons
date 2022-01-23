CREATE DATABASE onlineshop WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


\connect onlineshop

CREATE TABLE IF NOT EXISTS item
(
    id serial PRIMARY KEY,
    name VARCHAR ( 50 ) UNIQUE NOT NULL,
    description VARCHAR (1024),
    price INT NOT NULL,
    image_link VARCHAR (256)
);

CREATE ROLE shopapi WITH LOGIN PASSWORD 'jw8s0F4';
GRANT ALL PRIVILEGES ON DATABASE onlineshop TO shopapi;
GRANT ALL ON item to shopapi;
GRANT ALL ON item_id_seq to shopapi;