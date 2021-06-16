create table pokemon
(
	id serial PRIMARY KEY,
	number integer NOT NULL,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	type1 VARCHAR ( 50 ) NOT NULL,
	type2 VARCHAR ( 50 ),
	color VARCHAR ( 50 ) NOT NULL,
	ability1 VARCHAR ( 50 ) NOT NULL,
	ability2 VARCHAR ( 50 ),
	hiddenability VARCHAR ( 50 ),
	generation integer NOT NULL,
	height VARCHAR ( 50 ) NOT NULL,
	weigth VARCHAR ( 50 ) NOT NULL,
	hp integer NOT NULL,
	attack integer NOT NULL,
	defense integer NOT NULL,
	specialattack integer NOT NULL,
	specialdefense integer NOT NULL,
	speed integer NOT NULL
)