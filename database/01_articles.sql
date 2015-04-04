-- rambler up

CREATE TABLE articles (
	id integer primary key,
	title varchar(150) not null,
	slug varchar(150) not null,
	tagline varchar(450) not null,
	text text not null,
	tags text not null,
	is_published boolean,
	created_at datetime not null,
	updated_at datetime not null,
	published_at datetime
);
