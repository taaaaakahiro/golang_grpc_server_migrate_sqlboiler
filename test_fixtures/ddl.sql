select current_schema;

CREATE TABLE IF NOT EXISTS "user" (
  id integer,
  name varchar(255),
  PRIMARY KEY (id)
);