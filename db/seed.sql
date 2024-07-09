create table balances(
  id varchar(255) not null primary key,
  user_id varchar(255) not null unique,
  amount int not null default 1000,
  updated_at timestamp not null default current_timestamp
);

create table users(
    id varchar(255) not null primary key,
    username varchar(255) not null,
    password varchar(255) not null
);