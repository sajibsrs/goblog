drop table if exists posts;
drop table if exists sessions;
drop table if exists users;

create table users (
  id         int not null primary key,
  uuid       varchar(64) not null unique,
  fname      varchar(255),
  lname      varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);

create table sessions (
  id         int not null primary key,
  uuid       varchar(64) not null unique,
  fname      varchar(255),
  lname      varchar(255),
  email      varchar(255) not null unique,
  user_id    integer references users(id),
  created_at timestamp not null
);

create table posts (
  id         int not null primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  created_at timestamp not null
);