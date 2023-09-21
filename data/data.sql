insert into users (id, username, phone, password_hash, email, is_email_verified, birth_date, first_name , last_name, "address") values
('user00001', 'user1', '0123456789', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user1@gmail.com', 0, '01-01-2000', 'A', 'Nguyen', 'Quan 1, Tp HCM'),
('user00002', 'user2', '0123456782', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user2@gmail.com', 0, '01-01-2000', 'B', 'Nguyen', 'Ha Noi'),
('user00003', 'user3', '0123456783', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user3@gmail.com', 0, '01-01-2000', 'C', 'Nguyen', 'Da Nang'),
('user00004', 'user4', '0123456784', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user4@gmail.com', 0, '01-01-2000', 'D', 'Nguyen', 'Quan 1, Tp HCM'),
('user00005', 'user5', '0123456785', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user5@gmail.com', 0, '01-01-2000', 'E', 'Nguyen', 'Can Tho');


insert into hostels (id, name, province, district, ward, street, status, cost , electricity_price, water_price, parking_price, wifi_price, capacity, area, decription, created_by) values
('post1', 'First Hostel', 'Ha Noi', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 4000, 20000, 100000, 0, 2, 20, 'Nha dep', 'user00001'),
('post2', 'Second Hostel', 'Tp HCM', 'Quan 1', 'ward1', 'street1', 1, 2500000, 4500, 50000, 100000, 0, 3, 30, 'Nha dep', 'user00002'),
('post3', 'Third Hostel', 'Da Nang', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 3000, 70000, 100000, 0, 4, 25, 'Nha dep', 'user00003'),
('post4', 'Four Hostel', 'Can Tho', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 4000, 20000, 100000, 0, 1, 15, 'Nha dep', 'user00004'),
('post5', 'Five Hostel', 'Ha Noi', 'Ba Dinh', 'ward1', 'street1', 1, 1500000, 3500, 20000, 100000, 0, 5, 20, 'Nha dep', 'user00001');

insert into user_like_posts (user_id, post_id) values
('user00001', 'post1'),
('user00001', 'post2'),
('user00002', 'post1'),
('user00003', 'post3'),
('user00004', 'post4');
