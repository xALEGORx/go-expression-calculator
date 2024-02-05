CREATE USER calculator;

CREATE DATABASE calculator;
GRANT ALL PRIVILEGES ON DATABASE calculator TO calculator;

CREATE DATABASE calculator_test;
GRANT ALL PRIVILEGES ON DATABASE calculator_test TO calculator;

\c calculator;

create table tasks
(
    task_id serial primary key,
    expression text not null
);

