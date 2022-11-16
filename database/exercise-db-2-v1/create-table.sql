-- TODO: answer here
/*
TABLE = persons
- `id` bertipe `INTEGER` dan merupakan _primary key_
- `NIK` bertipe `VARCHAR(255)` yang tidak boleh `null` dan harus unik, kalian bisa menggunakan _constraints_ `UNIQUE`.
- `fullname` bertipe `VARCHAR(255)` yang tidak boleh `null`
- `gender` bertipe `VARCHAR(50)` yang tidak boleh `null`
- `birth_date` bertipe `DATE` yang tidak boleh `null`
- `is_married` bertipe `BOOLEAN`
- `height` beripe `FLOAT`
- `weight` bertipe `FLOAT`
- `address` bertipe `TEXT`
*/

CREATE TABLE persons (
  id INTEGER PRIMARY KEY,
  nik VARCHAR(255) UNIQUE NOT NULL,
  fullname VARCHAR(255) NOT NULL,
  gender VARCHAR(50) NOT NULL,
  birth_date DATE NOT NULL,
  is_married BOOLEAN,
  height FLOAT,
  weight FLOAT,
  address TEXT
);
