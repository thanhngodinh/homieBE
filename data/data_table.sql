CREATE TYPE gender AS ENUM ('Nam', 'Nữ', 'Nam hoặc Nữ');
CREATE TYPE verify AS ENUM ('V', 'N');

CREATE TABLE "users" (
  "id" varchar(40) PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "phone" varchar(20) default '',
  "email" varchar(100) NOT NULL,
  "is_verified_email" verify default 'N',
  "is_verified_phone" verify default 'N',
  "is_find_roommate" boolean default false,
  "avatar_url" varchar(500) default 'https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/default-avatar.png',
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

CREATE TYPE hostel_types AS ENUM ('Ký túc xá', 'Phòng cho thuê', 'Nhà nguyên căn', 'Phòng ở ghép', 'Căn hộ');
CREATE TYPE post_status AS ENUM ('W', 'A', 'I');

CREATE TABLE "posts" (
  "id" varchar(40) PRIMARY KEY,
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

CREATE TABLE  post_utilities (
  "post_id" VARCHAR(40) NOT NULL,
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
  user_id varchar(40),
  post_id varchar(40),
  primary key (user_id, post_id)
);

CREATE TABLE "user_rate_posts" (
  user_id varchar(40),
  post_id varchar(40),
  star int,
  comment text default '',
  primary key (user_id, post_id)
);
