-- rambler up

CREATE TABLE articles (
	id integer primary key,
	title varchar(150) not null,
	slug varchar(150) not null,
	tagline varchar(450) not null,
	text text not null,
	tags text not null,
	created_at datetime not null,
	updated_at datetime not null,
	is_published boolean
);
