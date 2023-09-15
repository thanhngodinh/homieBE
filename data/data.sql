CREATE TABLE "users" (
  "id" varchar(100) PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "phone" varchar(15),
  "password" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "is_email_verified" char(1) DEFAULT 0,
  "birth_date" timestamptz,
  "gender" varchar(20),
  "first_name" varchar(50),
  "last_name" varchar(50),
  "middle_name" varchar(50),
  "address" varchar(100),
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100),
  "updated_at" timestamptz,
  "updated_by" varchar(100)
);

CREATE TABLE "posts" (
  "id" varchar(100) PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "province" varchar(20) NOT NULL,
  "district" varchar(20) NOT NULL,
  "ward" varchar(20) NOT NULL,
  "street" varchar(100) NOT NULL,
  "post_type" char(1) NOT NULL DEFAULT 'H',
  "status" char(1) NOT NULL DEFAULT 0,
  "cost" int NOT NULL,
  "electricity_price" int NOT NULL DEFAULT 0,
  "water_price" int NOT NULL DEFAULT 0,
  "parking_price" int NOT NULL DEFAULT 0,
  "wifi_price" int NOT NULL DEFAULT 0,
  "capacity" int,
  "area" int,
  "decription" varchar,
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100) NOT NULL,
  "updated_at" timestamptz,
  "updated_by" varchar(100)
);

insert into users (id, username, phone, password_hash, email, is_email_verified, birth_date, first_name , last_name, "address") values
('user00001', 'user1', '0123456789', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user1@gmail.com', 0, '01-01-2000', 'A', 'Nguyen', 'Quan 1, Tp HCM'),
('user00002', 'user2', '0123456782', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user2@gmail.com', 0, '01-01-2000', 'B', 'Nguyen', 'Ha Noi'),
('user00003', 'user3', '0123456783', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user3@gmail.com', 0, '01-01-2000', 'C', 'Nguyen', 'Da Nang'),
('user00004', 'user4', '0123456784', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user4@gmail.com', 0, '01-01-2000', 'D', 'Nguyen', 'Quan 1, Tp HCM'),
('user00005', 'user5', '0123456785', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user5@gmail.com', 0, '01-01-2000', 'E', 'Nguyen', 'Can Tho');


insert into posts (id, name, province, district, ward, street, status, cost , electricity_price, water_price, parking_price, wifi_price, capacity, area, decription) values
('post1', 'First Hostel', 'Ha Noi', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 4000, 20000, 100000, 0, 2, 20, 'Nha dep'),
('post2', 'Second Hostel', 'Tp HCM', 'Quan 1', 'ward1', 'street1', 1, 2500000, 4500, 50000, 100000, 0, 3, 30, 'Nha dep'),
('post3', 'Third Hostel', 'Da Nang', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 3000, 70000, 100000, 0, 4, 25, 'Nha dep'),
('post4', 'Four Hostel', 'Can Tho', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 4000, 20000, 100000, 0, 1, 15, 'Nha dep'),
('post5', 'Five Hostel', 'Ha Noi', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 3500, 20000, 100000, 0, 5, 20, 'Nha dep');


