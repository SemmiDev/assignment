-- TODO: answer here


/*
id bertipe INT tidak boleh null
user_id bertipe INT tidak boleh null
presence_date bertipe DATE tidak boleh null
status bertipe VARCHAR(50) tidak boleh null
location bertipe VARCHAR(255) boleh diisi data kosong / null
description bertipe VARCHAR(255) boleh diisi data kosong / null
image_presence bertipe VARCHAR(255) boleh diisi data kosong / null
image_location bertipe VARCHAR(255) boleh diisi data kosong / null

*/
CREATE TABLE presences (
    id INT NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    presence_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    location VARCHAR(255),
    description VARCHAR(255),
    image_presence VARCHAR(255),
    image_location VARCHAR(255)
);
