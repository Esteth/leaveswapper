CREATE TYPE order_type AS ENUM('buy', 'sell');

CREATE TABLE orders(
  user_id varchar(256) NOT NULL REFERENCES users(email),
  date date NOT NULL,
  type order_type NOT NULL,
  PRIMARY KEY(user_id, date)
);