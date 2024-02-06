create user calculator;

create database calculator;
grant all privileges on database calculator to calculator;

create database calculator_test;
grant all privileges on database calculator_test to calculator;

\c calculator;

create table tasks
(
    task_id serial primary key,
    expression text not null,
    status varchar(10) not null,
    error text not null
);

