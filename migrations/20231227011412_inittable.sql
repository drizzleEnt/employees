-- +goose Up
CREATE TABLE employee(
    id serial primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    phone varchar(255) not null,
    companyId INTEGER not null
);

CREATE TABLE passport(
    id serial primary key,
    employee_id int REFERENCES employee (id) on DELETE  CASCADE NOT NULL,
    type varchar(255) not null,
    number varchar(255) not null
);

CREATE TABLE departament(
    id serial primary key,
    employee_id int REFERENCES employee (id) on DELETE  CASCADE NOT NULL,
    company_id INT NOT NULL DEFAULT '1',
    depname varchar(255) not null,
    depphone varchar(255) not null
);



-- +goose Down
DROP TABLE departament;
DROP TABLE passport;
DROP TABLE employee;

