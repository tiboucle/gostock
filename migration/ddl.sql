CREATE TABLE colors(
	code 			VARCHAR(3) PRIMARY KEY,
  	name 			VARCHAR(50),
  	created_at 		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at 		DATETIME
);

CREATE TABLE items(
	sku 			VARCHAR(50) PRIMARY KEY,
  	name 			VARCHAR(100),
	size 			VARCHAR(5),
	color 			VARCHAR(3),
	stock 			INTEGER,
	created_at		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at		DATETIME,
	FOREIGN KEY(color) REFERENCES colors(code)
);

CREATE TABLE purchases(
	invoice_id		VARCHAR(50) PRIMARY KEY,
  	invoice_date	DATETIME,
  	created_at		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at		DATETIME
);

CREATE TABLE purchase_details(
	id				INTEGER PRIMARY KEY,
	invoice_id		VARCHAR(50),
  	item_sku		VARCHAR(50),
  	price			NUMERIC,
  	amount			INTEGER,
  	created_at		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at		DATETIME,
	FOREIGN KEY(invoice_id) REFERENCES purchases(invoice_id),
	FOREIGN KEY(item_sku) REFERENCES items(sku)
);

CREATE TABLE purchase_lines(
	id 				INTEGER PRIMARY KEY AUTOINCREMENT,
	invoice_id		VARCHAR(50),
	item_sku 		VARCHAR(50),
	item_received	INTEGER,
	created_at		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at		DATETIME,
	FOREIGN KEY(invoice_id) REFERENCES purchases(invoice_id),
	FOREIGN KEY(item_sku) REFERENCES items(sku)
);

CREATE TABLE orders(
	invoice_id		VARCHAR(50) PRIMARY KEY,
  	invoice_date	DATETIME,
  	created_at		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at		DATETIME
);

CREATE TABLE order_details(
	id 				INTEGER PRIMARY KEY,
	invoice_id		VARCHAR(50),
  	item_sku		VARCHAR(50),
  	price			NUMERIC,
  	amount			INTEGER,
  	created_at		DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at 		DATETIME,
	deleted_at		DATETIME,
	FOREIGN KEY(invoice_id) REFERENCES orders(invoice_id),
	FOREIGN KEY(item_sku) REFERENCES items(sku)
);