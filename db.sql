CREATE DATABASE xyz_multifinance

USE xyz_multifinance

CREATE TABLE customers(
 id INT NOT NULL AUTO_INCREMENT,
 nik VARCHAR(25) NOT NULL,
 full_name VARCHAR(200) NOT NULL,
 legal_name VARCHAR(200) NOT NULL,
 birth_place VARCHAR(200) NOT NULL,
 birth_date DATE NOT NULL,
 salary INT NOT NULL, 
 ktp_photo VARCHAR(255) NOT NULL,
 selfie_photo VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 deleted_at TIMESTAMP, 
 PRIMARY KEY (id)
)ENGINE InnoDB;

SELECT * FROM  customers;

CREATE TABLE customer_transactions(
	id INT NOT NULL AUTO_INCREMENT,
	customer_id INT,
	contract_number VARCHAR(255),
	otr_price float,
	admin_fee float,
	installment_amount float,
	interest_amount float,
	asset_name VARCHAR(255),
	status VARCHAR(50),
	created_at TIMESTAMP NOT NULL,
 	updated_at TIMESTAMP NOT NULL,
 	deleted_at TIMESTAMP, 
  	PRIMARY KEY (id)
)ENGINE InnoDB;

DESC customer_transactions;

CREATE TABLE customer_limits(
	id INT NOT NULL AUTO_INCREMENT,
	customer_id INT,
	limit_1 float,
	limit_2 float,
	limit_3 float,
	limit_4 float,
	  remaining_limit float,
	created_at TIMESTAMP NOT NULL,
 	updated_at TIMESTAMP NOT NULL,
 	deleted_at TIMESTAMP, 
  	PRIMARY KEY (id)
)ENGINE InnoDB;

DESC customer_limits;

ALTER TABLE customer_transactions ADD FOREIGN KEY (customer_id) REFERENCES customers(id);

ALTER TABLE customer_limits ADD FOREIGN KEY (customer_id) REFERENCES customers(id);

CREATE table customer_installments(
id INT NOT NULL AUTO_INCREMENT, 
customer_transaction_id integer,
  customer_limit_id integer,
  tenor int,
  total_amounts float,
  remaining_amounts float,
  created_at timestamp NOT NULL ,
  updated_at timestamp NOT null,
  deleted_at timestamp,
  	PRIMARY KEY (id)
)ENGINE InnoDB;

ALTER TABLE customer_installments ADD FOREIGN KEY (customer_transaction_id) REFERENCES customer_transactions(id);

ALTER TABLE customer_installments ADD FOREIGN KEY (customer_limit_id) REFERENCES customer_limits(id);

INSERT INTO xyz_multifinance.customers
(nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo, created_at, updated_at)
VALUES('123456789123', 'Budi Budiman', 'Budi Budiman', 'Jakarta', '1990-10-10', 15000000, 'url_ktp_photo', 'url_selfie_photo', now(), now()),
('789789123', 'Annisa Nissa', 'Annisa Nissa', 'Surabata', '1997-07-01', 10000000, 'url_ktp_photo', 'url_selfie_photo', now(), now());

INSERT INTO xyz_multifinance.customer_limits
(customer_id, limit_1, limit_2, limit_3, limit_4, remaining_limit, created_at, updated_at)
VALUES(1, 100000, 200000, 500000, 700000, 700000, now(), now()),
(2, 1000000, 1200000, 1500000, 2000000, 2000000, now(), now());
