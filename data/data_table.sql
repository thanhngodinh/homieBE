CREATE TABLE "users" (
  "id" varchar(40) PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "phone" varchar(10) default '',
  "email" varchar(100) NOT NULL,
  "is_email_verified" char(1) DEFAULT 0,
  "first_name" varchar(50) default '',
  "last_name" varchar(50) default '',
  "province_suggest" varchar(20) default '',
  "district_suggest" varchar(20) default '',
  "cost_suggest" int default 2000000,
  "capacity_suggest" int default 1,
  "gender" varchar(10) default '',
  "date_of_birth" timestamptz,
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100),
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "hostels" (
  "id" varchar(40) PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "province" varchar(20) NOT NULL,
  "district" varchar(20) NOT NULL,
  "ward" varchar(20) NOT NULL,
  "street" varchar(100) NOT NULL,
  "status" char(1) NOT NULL DEFAULT 'W',
  "cost" int NOT NULL,
  "electricity_price" int NOT NULL DEFAULT 0,
  "water_price" int NOT NULL DEFAULT 0,
  "parking_price" int NOT NULL DEFAULT 0,
  "wifi_price" int NOT NULL DEFAULT 0,
  "capacity" int NOT NULL DEFAULT 0,
  "area" int DEFAULT 0,
  "decription" text default '',
  "view" int DEFAULT 0,
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100) NOT NULL,
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE hostels_utilities (
  "hostel_id" VARCHAR(40) NOT NULL,
	"utilities_id" VARCHAR(40) NOT NULL,
  primary key (hostel_id, utilities_id)
);

CREATE TABLE utilities (
  "id" VARCHAR(40) PRIMARY KEY,
	"icon" VARCHAR(255) default '',
	"name" VARCHAR(255) NOT NULL,
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100) default '',
  "updated_at" timestamptz,
  "updated_by" varchar(100) default ''
);

CREATE TABLE "user_like_posts" (
  user_id varchar(40),
  post_id varchar(40),
  primary key (user_id, post_id)
);
