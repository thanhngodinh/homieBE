CREATE TABLE "users" (
  "id" varchar(20) PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "phone" varchar(10) default '',
  "password" varchar(100) NOT NULL,
  "email" varchar(100) NOT NULL,
  "is_email_verified" char(1) DEFAULT 0,
  "birth_date" timestamptz,
  "gender" varchar(10) default '',
  "first_name" varchar(50) default '',
  "last_name" varchar(50) default '',
  "middle_name" varchar(50) default '',
  "address" varchar(100) default '',
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100 NOT NULL),
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "hostels" (
  "id" varchar(20) PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "province" varchar(20) NOT NULL,
  "district" varchar(20) NOT NULL,
  "ward" varchar(20) NOT NULL,
  "street" varchar(100) NOT NULL,
  "status" char(1) NOT NULL DEFAULT 0,
  "cost" int NOT NULL,
  "electricity_price" int NOT NULL DEFAULT 0,
  "water_price" int NOT NULL DEFAULT 0,
  "parking_price" int NOT NULL DEFAULT 0,
  "wifi_price" int NOT NULL DEFAULT 0,
  "capacity" int NOT NULL DEFAULT 0,
  "area" int DEFAULT 0,
  "decription" text default '',
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100) NOT NULL,
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "user_like_posts" (
  user_id varchar(20),
  post_id varchar(20),
  primary key (user_id, post_id)
);
