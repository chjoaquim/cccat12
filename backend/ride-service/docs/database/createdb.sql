-- noinspection SqlNoDataSourceInspectionForFile

-- noinspection SqlDialectInspectionForFile

create table passengers
(
    id         uuid                not null PRIMARY KEY,
    name       varchar(255)        NOT NULL,
    email      varchar(150)        NOT NULL,
    document   varchar(255) UNIQUE NOT NULL,
    created_at timestamp           NOT NULL,
    updated_at timestamp
);

create table drivers
(
    id         uuid                not null PRIMARY KEY,
    name       varchar(255)        NOT NULL,
    email      varchar(150)        NOT NULL,
    document   varchar(255) UNIQUE NOT NULL,
    car_plate  varchar(10),
    created_at timestamp           NOT NULL,
    updated_at timestamp
);