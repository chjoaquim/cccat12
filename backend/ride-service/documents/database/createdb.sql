create table passengers
(
    id         uuid                not null PRIMARY KEY,
    name       varchar(255)        NOT NULL,
    email      varchar(150)        NOT NULL,
    document   varchar(255) UNIQUE NOT NULL,
    created_at timestamp           NOT NULL,
    updated_at timestamp
);
