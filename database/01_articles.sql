-- rambler up

CREATE TABLE articles (
	id integer primary key,
	title varchar(50) not null,
	tagline varchar(50) not null,
	slug varchar(50) not null,
	text text not null,
	created_at datetime not null,
	updated_at datetime not null,
	published boolean
);
