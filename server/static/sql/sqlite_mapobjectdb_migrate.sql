PRAGMA foreign_keys = ON;
PRAGMA journal_mode = MEMORY;

-- mapobjects

create table if not exists objects(
	id integer primary key autoincrement,
  x int,
  y int,
	z int,
	posx int,
	posy int,
	posz int,
	type varchar,
  mtime bigint
);

create index if not exists objects_pos on objects(posx,posy,posz);
create index if not exists objects_pos_type on objects(posx,posy,posz,type);

-- mapobject attributes

create table if not exists object_attributes(
	objectid integer not null,
	key varchar not null,
	value varchar not null,
	FOREIGN KEY (objectid) references objects(id) ON DELETE CASCADE
	primary key(objectid, key)
);

create index if not exists object_attributes_key_value on object_attributes(key, value);

-- settings

create table if not exists settings(
	key varchar primary key not null,
	value varchar not null
);

-- subscriptions

create table if not exists subscriptions(
	id integer primary key autoincrement,
	endpoint varchar not null,
	auth varchar not null,
	p256dh varchar not null,
	x int,
	y int,
	z int
);

create index if not exists subscriptions_pos on subscriptions(x,y,z);
create index if not exists subscriptions_endpoint on subscriptions(endpoint);
