-- TODO: answer here

ALTER TABLE students ADD COLUMN date_of_birth DATE NOT NULL;

ALTER TABLE students
    ADD COLUMN street VARCHAR(255),
    ADD city VARCHAR(100),
    ADD COLUMN province VARCHAR(100),
    ADD COLUMN country VARCHAR(100),
    ADD COLUMN postal_code VARCHAR(50);
