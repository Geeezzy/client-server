CREATE TABLE "users" (
  "id" serial NOT NULL,
  "username" character varying(300) NOT NULL,
  "first_name" text NULL,
  "last_name" text NULL
);

INSERT INTO users (username, first_name, last_name)
VALUES ('demo', 'Roman', 'Kuzmenko');

INSERT INTO users (username, first_name, last_name)
VALUES ('demo2', 'Roman1', 'Kuzmenko1');