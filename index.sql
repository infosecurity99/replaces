create table  countries(
    id serial  primary key,
    name varchar(30),
    currency varchar(10)
);

create table cities(
    id serial,
    name varchar(30),
    population int,
    country_id  int  references countries(id)
);