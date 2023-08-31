CREATE TABLE users(
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY
);

CREATE TABLE segments(
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name varchar(255) UNIQUE NOT NULL
);

CREATE TABLE segments_users(
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  segment_id bigint REFERENCES segments ON UPDATE CASCADE ON DELETE CASCADE,
  user_id bigint REFERENCES users ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE segments_users_history(
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id bigint REFERENCES users ON UPDATE CASCADE ON DELETE CASCADE,
  segment_name varchar(255) NOT NULL,
  operation varchar(15) NOT NULL,
  created_at timestamp DEFAULT now()
);
