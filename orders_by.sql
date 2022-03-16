CREATE DATABASE orders_by;

CREATE IF NOT EXISTS SCHEMA orders;

CREATE SEQUENCE IF NOT EXISTS orders.orders_id_seq;

CREATE SEQUENCE IF NOT EXISTS orders.item_id_seq;

CREATE TABLE orders.orders (
	"orders_id" int4 NOT NULL DEFAULT nextval('orders.orders_id_seq'::regclass),
	"customer_name" VARCHAR(120),
	"ordered_at" TIMESTAMP,
	PRIMARY KEY(orders_id)
);

CREATE TABLE orders.items (
	"item_id" int4 NOT NULL DEFAULT nextval('orders.item_id_seq'::regclass),
	"item_code" VARCHAR(120),
	"description" VARCHAR(120),
	"quantity" int4,
	"orders_id" int4,
	PRIMARY KEY(item_id),
	CONSTRAINT fk_orders
		FOREIGN KEY(orders_id)
			REFERENCES orders.orders(orders_id)
);