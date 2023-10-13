CREATE TABLE "users" (
  "id" varchar(40) PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "phone" varchar(10) default '',
  "email" varchar(100) NOT NULL,
  "is_verified_email" char(1) DEFAULT 0,
  "is_find_roommate" boolean DEFAULT false,
  "avatar_url" varchar(500) default 'https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/default-avatar.png',
  "display_name" varchar(100) default '',
  "gender" char(1) default 'M',
  "date_of_birth" timestamptz,
  "province_profile" varchar(40) default '',
  "district_profile" varchar(40)[] default '{}',
  "cost_from" int default 0,
  "cost_to" int default 0,
  "province_suggest" varchar(40) default '',
  "district_suggest" varchar(40) default '',
  "cost_suggest" int default 2000000,
  "capacity_suggest" int default 1,
  "created_at" timestamptz DEFAULT (current_timestamp),
  "created_by" varchar(100),
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "hostels" (
  "id" varchar(40) PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "province" varchar(40) NOT NULL,
  "district" varchar(40) NOT NULL,
  "ward" varchar(40) NOT NULL,
  "street" varchar(100) NOT NULL,
  "status" char(1) NOT NULL DEFAULT 'W',
  "cost" int NOT NULL,
  "deposit" int DEFAULT 0,
  "electricity_price" int DEFAULT 0,
  "water_price" int DEFAULT 0,
  "parking_price" int DEFAULT 0,
  "service_price" int DEFAULT 0,
  "capacity" int DEFAULT 1,
  "area" int DEFAULT 0,
  "description" text default '',
  "phone" varchar(10) NOT NULL,
  "image_url" varchar[] NOT NULL,
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
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "user_like_posts" (
  user_id varchar(40),
  post_id varchar(40),
  primary key (user_id, post_id)
);
