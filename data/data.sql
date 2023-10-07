insert into users (id, username, phone, "password", email, is_email_verified, first_name , last_name, province_suggest, district_suggest, cost_suggest, capacity_suggest, gender) values
('user00001', 'user1', '0123456789', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user1@gmail.com', 1, 'A', 'Nguyen', 'Hà Nội', 'Ba Đình', 4000000, 2, 'women'),
('user00002', 'user2', '0123456782', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user2@gmail.com', 1, 'B', 'Nguyen', 'Hà Nội', 'Ba Đình', 2000000, 1, 'women'),
('user00003', 'user3', '0123456783', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user3@gmail.com', 1, 'C', 'Nguyen', 'Hà Nội', 'Hoàn Kiếm', 6000000, 3, 'men'),
('user00004', 'user4', '0123456784', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user4@gmail.com', 1, 'D', 'Nguyen', 'Hồ Chí Minh', 'Quận 1', 5000000, 2, 'men'),
('user00005', 'user5', '0123456785', '$2a$10$J4BEfdZT3rWL5lwcDOum6ugFvdTQ31zub0zJL2xMwvA3snH/KUCCy', 'user5@gmail.com', 1, 'E', 'Nguyen', 'Hồ Chí Minh', 'Quận 12', 4000000, 2, 'women');

insert into hostels (id, "name", province, district, ward, street, "status", cost, deposit, electricity_price, water_price, parking_price, service_price, capacity, area, description, created_by, phone, image_url) values
('post1', 'Nhà trọ cao cấp', 'Thành phố Hà Nội', 'Quận Ba Đình', 'Phường Trúc Bạch', '11 Hai Bà Trưng', 'A', 1500000, 1500000, 4000, 20000, 100000, 0, 2, 20, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00001', '0987654321', '{"https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg", "https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png"}'),
('post2', 'Căn hộ cao cấp', 'Thành phố Hồ Chí Minh', 'Quận 1', 'Phường Tân Định', '11 Hai Bà Trưng', 'A', 5500000, 5500000, 4500, 50000, 100000, 0, 3, 30, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00002', '0987654322', '{"https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg", "https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png"}'),
('post3', 'Ký túc xá', 'Thành phố Hồ Chí Minh', 'Quận 12', 'Phường Thạnh Lộc', '121 Quang Trung', 'A', 1500000, 2500000, 3000, 70000, 100000, 0, 4, 25, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00003', '0987654323', '{"https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg", "https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png"}'),
('post4', 'Nhà trọ cao cấp', 'Thành phố Hồ Chí Minh', 'Quận Phú Nhuận', 'Phường 05', '743 Nguyễn Kiệm', 'A', 3500000, 2500000, 4000, 20000, 100000, 0, 1, 15, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00004', '0987654324', '{"https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg", "https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png"}'),
('post5', 'Nhà trọ cao cấp', 'Thành phố Hà Nội', 'Long Biên', 'Phường 11', '66 Nguyễn Tri Phương', 'A', 2500000, 2500000, 3500, 20000, 100000, 0, 5, 20, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00001', '0987654321', '{"https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg", "https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png"}'),
('post6', 'Ký túc xá', 'Thành phố Hà Nội', 'Quận Ba Đình', 'Phường Phúc Xá', '11  Nguyễn Tri Phương', 'A', 1500000, 1500000, 4000, 20000, 100000, 0, 2, 20, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00001', '0987654321', '{"https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png","https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg"}'),
('post7', 'Chung cư A', 'Thành phố Hồ Chí Minh', 'Quận 1', 'Phường Tân Định', '321 Lê Đại Hành', 'A', 2500000, 2500000, 4500, 50000, 100000, 0, 3, 30, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00002', '0987654322', '{"https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png","https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg"}'),
('post8', 'Nhà trọ A', 'Thành phố Hồ Chí Minh', 'Quận 12', 'Phường Thạnh Lộc', '34 Hoàng Văn Thụ', 'A', 1500000, 1500000, 3000, 70000, 100000, 0, 4, 25, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00003', '0987654323', '{"https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png","https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg"}'),
('post9', 'Nhà nguyên căn', 'Thành phố Hồ Chí Minh', 'Quận Bình Thạnh', 'Phường 11', '66 Phạm Văn Đồng', 'A', 7500000, 5500000, 4000, 20000, 100000, 0, 1, 15, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00004', '0987654324', '{"https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png","https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg"}'),
('post10', 'Nhà trọ B', 'Thành phố Hà Nội', 'Quận Hoàn Kiếm', 'Phường Phúc Tâm', '432 Hai Bà Trưng', 'A', 2500000, 1500000, 3500, 20000, 100000, 0, 5, 20, 'Với kiến trúc nhà nhỏ khá hẹp như vậy thì gia chủ nên tối giản hóa thiết kế nhà, hạn chế những nội thất không cần thiết, và thay thế bằng những vật dụng đa năng. Ngoài ra, với không gian mẫu nhà nhỏ này, việc kết hợp nhà bếp và phòng khách, và xây dựng thêm gác lửng sẽ giúp tiết kiệm không gian và mang đến vẻ đẹp riêng cho căn nhà.', 'user00001', '0987654321', '{"https://vinhtuong.com/sites/default/files/inline-images/dac-diem-nha-cap-4.png","https://kantechpaint.com/wp-content/uploads/2023/02/thiet-ke-nha-pho-2-tang-mai-thai.jpg", "https://achi.vn/wp-content/uploads/2021/07/Thiet-ke-biet-thu-nha-vuon-2-tang-mai-thai-dep-300m2-tai-dong-nai-achi-22102-01.jpg"}');

insert into utilities ("id", "name", "icon") values
('wifi', 'Wifi', 'wifi'),
('washing_machine','Máy giặt', 'table-columns'),
('kitchen','Nhà bếp', 'kitchen-set'),
('air_conditioning','Điều hòa', 'temperature-half'),
('fridge','Tủ lạnh', 'snowflake'),
('pet','Cho nuôi pet', 'dog'),
('wardrobe','Tủ đồ', 'boxes-packing'),
('bed','Giường', 'bed'),
('balcony','Ban công', 'sun'),
('parking','Bãi giữ xe', 'motorcycle'),
('camera','Camera an ninh', 'camera'),
('bathroom','Nhà tắm riêng', 'restroom');

insert into user_like_posts (user_id, post_id) values
('user00001', 'post1'),
('user00001', 'post2'),
('user00002', 'post1'),
('user00003', 'post3'),
('user00004', 'post4');

insert into hostels_utilities (hostel_id, utilities_id) values
('post1', 'wifi'),
('post2', 'wifi'),
('post3', 'wifi'),
('post4', 'wifi'),
('post5', 'wifi'),
('post6', 'wifi'),
('post7', 'wifi'),
('post8', 'wifi'),
('post9', 'wifi'),
('post10', 'wifi'),
('post1', 'washing_machine'),
('post3', 'washing_machine'),
('post4', 'washing_machine'),
('post5', 'washing_machine'),
('post6', 'washing_machine'),
('post6', 'kitchen'),
('post7', 'kitchen'),
('post1', 'kitchen'),
('post2', 'kitchen'),
('post3', 'kitchen'),
('post2', 'air_conditioning'),
('post3', 'air_conditioning'),
('post2', 'fridge'),
('post3', 'pet'),
('post4', 'pet'),
('post5', 'pet'),
('post2', 'wardrobe'),
('post3', 'wardrobe'),
('post4', 'wardrobe'),
('post2', 'bed'),
('post1', 'bed'),
('post3', 'bed'),
('post4', 'bed'),
('post1', 'balcony'),
('post2', 'balcony'),
('post3', 'balcony'),
('post4', 'balcony'),
('post1', 'parking'),
('post2', 'parking'),
('post3', 'parking'),
('post4', 'parking'),
('post5', 'parking'),
('post6', 'parking'),
('post7', 'parking'),
('post8', 'parking'),
('post9', 'parking'),
('post2', 'camera'),
('post5', 'camera'),
('post2', 'bathroom'),
('post3', 'bathroom'),
('post4', 'bathroom');