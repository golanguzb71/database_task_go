package mock

import "database/sql"

func CreateTables(db *sql.DB) bool {
	var sql = `
	CREATE TABLE branch
(
    id   serial PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE users
(
    id   serial PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE product
(
    id            serial PRIMARY KEY,
    name          varchar          NOT NULL,
    income_price  double precision NOT NULL,
    outcome_price double precision NOT NULL
);

CREATE TABLE orders
(
    id           serial PRIMARY KEY,
    user_id      bigint references users (id) NOT NULL,
    branch_id 	bigint references branch(id) NOT NULL ,
    discount_price double precision NOT NULL
);

CREATE TABLE order_items
(
    id         serial PRIMARY KEY,
    order_id   bigint references orders (id)  NOT Null,
    product_id bigint references product (id) NOT NULL,
);

`
	_, err := db.Exec(sql)
	if err != nil {
		return false
	}
	return true
}
