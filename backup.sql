CREATE DATABASE taptalkdiaries;
USE taptalkdiaries;

CREATE TABLE master_users (
    id Bigint auto_increment primary key not null,
    username varchar(32) UNIQUE not null, 
    email varchar(64) UNIQUE not null,
    birthday varchar(10) not null,
    full_name varchar(16) not null,
    created_at datetime not null,
    updated_at datetime,
    deleted_at datetime
);


CREATE TABLE user_auths (
   user_id bigint primary key not null,
   `password` text not null,
   device_id text,
   last_login datetime,
   FOREIGN KEY (user_id) references master_users(ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE user_diaries (
	id bigint auto_increment primary key not null,
	user_id bigint not null,
	title text not null,
	body_text text not null,
	date_post date not null,
	created_at datetime not null, 
	updated_at datetime,
	deleted_at datetime,
	FOREIGN KEY (user_id) references master_users(id) ON UPDATE CASCADE ON DELETE CASCADE
);