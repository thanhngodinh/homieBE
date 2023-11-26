CREATE TYPE gender AS ENUM ('Nam', 'Nữ', 'Nam hoặc Nữ');

CREATE TABLE "users" (
  "id" varchar(50) PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "phone" varchar(20) default '',
  "email" varchar(100) NOT NULL,
  "avatar_url" varchar(500) default 'https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/default-avatar.png',
  "is_verified_email" boolean default false,
  "is_verified_phone" boolean default false,
  "otp" varchar(10) default '',
  "expiration_time" timestamptz,
  "is_find_roommate" boolean default false,
  "display_name" varchar(100) default '',
  "gender" gender default 'Nam hoặc Nữ',
  "date_of_birth" timestamptz,
  "province_profile" varchar(40) default '',
  "district_profile" varchar(40)[] default '{}',
  "cost_from" int default 0,
  "cost_to" int default 0,
  "province_suggest" varchar(40) default '',
  "district_suggest" varchar(40) default '',
  "cost_suggest" int default 2000000,
  "capacity_suggest" int default 1,
  "created_at" timestamptz default (current_timestamp),
  "created_by" varchar(100),
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "admins" (
  "id" varchar(40) PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "phone" varchar(20) default '',
  "email" varchar(100) NOT NULL,
  "avatar_url" varchar(500) default 'https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/default-avatar.png',
  "display_name" varchar(100) default '',
  "created_at" timestamptz default (current_timestamp),
  "created_by" varchar(100) default 'system',
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TYPE hostel_types AS ENUM ('Ký túc xá', 'Phòng cho thuê', 'Nhà nguyên căn', 'Phòng ở ghép', 'Căn hộ');
CREATE TYPE post_status AS ENUM ('W', 'A', 'I', 'V');

CREATE TABLE "posts" (
  "id" varchar(50) PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "province" varchar(40) NOT NULL,
  "district" varchar(40) NOT NULL,
  "ward" varchar(40) NOT NULL,
  "street" varchar(100) NOT NULL,
  "status" post_status NOT NULL default 'W',
  "type" hostel_types NOT NULL default 'Phòng cho thuê',
  "gender" gender default 'Nam hoặc Nữ',
  "cost" int NOT NULL,
  "deposit" int default 0,
  "electricity_price" int default 0,
  "water_price" int default 0,
  "parking_price" int default 0,
  "service_price" int default 0,
  "capacity" int default 1,
  "area" int default 0,
  "description" text default '',
  "image_url" varchar[] NOT NULL,
  "view" int default 0,
  "longitude" text default '',
  "latitude" text default '',
  "created_at" timestamptz default (current_timestamp),
  "ended_at" timestamptz default (current_timestamp + interval '1' month),
  "created_by" varchar(100) NOT NULL,
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE post_utilities (
  "post_id" VARCHAR(50) NOT NULL,
	"utility_id" VARCHAR(40) NOT NULL,
  primary key (post_id, utility_id)
);

CREATE TABLE utilities (
  "id" VARCHAR(40) PRIMARY KEY,
	"icon" VARCHAR(255) default '',
	"name" VARCHAR(255) NOT NULL,
  "created_at" timestamptz default (current_timestamp),
  "created_by" varchar(100) default '',
  "updated_at" timestamptz,
  "updated_by" varchar(100) default '',
  "deleted_at" timestamptz
);

CREATE TABLE "user_like_posts" (
  "user_id" varchar(50),
  "post_id" varchar(50),
  primary key (user_id, post_id)
);

CREATE TABLE "rates" (
  "user_id" varchar(50) NOT NULL,
  "post_id" varchar(50) NOT NULL,
  "star" int NOT NULL,
  "comment" text default '',
  "created_at" timestamptz default (current_timestamp),
  "updated_at" timestamptz,
  "deleted_at" timestamptz,
  primary key (user_id, post_id)
);

CREATE TABLE "post_rate_info" (
  "id" bigserial PRIMARY KEY,
  "post_id" varchar(50) NOT NULL UNIQUE,
  "total" int NOT NULL default 0,
  "star1" int NOT NULL default 0,
  "star2" int NOT NULL default 0,
  "star3" int NOT NULL default 0,
  "star4" int NOT NULL default 0,
  "star5" int NOT NULL default 0
);
