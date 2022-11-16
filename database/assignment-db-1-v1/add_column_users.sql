-- TODO: answer here


/*
phone bertipe varchar(50) boleh diisi data kosong / null
address bertipe varchar(255) boleh diisi data kosong / null
department bertipe varchar(255) boleh diisi data kosong / null
division bertipe varchar(255) boleh diisi data kosong / null
position bertipe varchar(255) boleh diisi data kosong / null

*/

ALTER TABLE users
    ADD COLUMN phone VARCHAR(50),
    ADD COLUMN address VARCHAR(255),
    ADD COLUMN department VARCHAR(255),
    ADD COLUMN division VARCHAR(255),
    ADD COLUMN position VARCHAR(255);
