CREATE TABLE IF NOT EXISTS banners (banner_id SERIAL PRIMARY KEY, banner_name TEXT, banner_category INT, banner_expired DATE, banner_start DATE, banner_image_url TEXT);

CREATE TABLE IF NOT EXISTS users (user_id SERIAL PRIMARY KEY, user_name TEXT, user_location INT, user_tier INT);

CREATE TABLE IF NOT EXISTS locations (location_id SERIAL PRIMARY KEY, location_name TEXT);

CREATE TABLE IF NOT EXISTS tiers (tier_id SERIAL PRIMARY KEY, tier_name TEXT);

CREATE TABLE IF NOT EXISTS bannercategories (banner_category_id SERIAL PRIMARY KEY, banner_category_name TEXT);

CREATE TABLE IF NOT EXISTS bannertiertable (bannertiertable_id SERIAL PRIMARY KEY, banner_id INT, tier_id INT);

CREATE TABLE IF NOT EXISTS bannerlocationtable (bannerlocationtable_id SERIAL PRIMARY KEY, banner_id INT, location_id INT);

INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(1, 1);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(1, 2);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(1, 3);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(2, 1);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(2, 2);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(3, 1);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(4, 2);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(5, 3);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(6, 1);
INSERT INTO bannerlocationtable (banner_id, location_id) VALUES(7, 1);

INSERT INTO bannertiertable (banner_id, tier_id) VALUES(1, 1);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(1, 2);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(1, 3);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(1, 4);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(1, 2);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(2, 2);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(3, 2);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(2, 3);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(3, 3);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(4, 3);
INSERT INTO bannertiertable (banner_id, tier_id) VALUES(5, 3);

INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('GoPayLater Diskon 35%',2, '2021-10-01', '2021-10-20', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Komisaris Tokopedia Amin', 4, '2021-10-01', '2021-11-22', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Hemat Ongkir Gojek 1 Rupiah', 1, '2021-10-20', '2021-10-22', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Cashback Hingga 700 Ribu', 4, '2021-10-04', '2021-10-12', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Top Up Cashback 2500 Rupiah', 3, '2021-11-06', '2021-12-20', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Cari Martabak Lewat GoFood 75% Promo', 3, '2022-01-01', '2022-03-20', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Ongkir Hanya 1 Rupiah', 3, '2021-10-11', '2021-10-23', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Paket Murah 35%', 4, '2021-11-01', '2021-12-22', 'www.google.com');
INSERT INTO banners (banner_name, banner_category, banner_expired, banner_start, banner_image_url) VALUES('Anti Lelet Anti Macet, GoCar', 1, '2021-10-01', '2021-11-20', 'www.google.com');

INSERT INTO bannerCategories (banner_category_name) VALUES('GoRide');
INSERT INTO bannerCategories (banner_category_name) VALUES('GoPay');
INSERT INTO bannerCategories (banner_category_name) VALUES('GoFood');
INSERT INTO bannerCategories (banner_category_name) VALUES('Tokopedia Mall');

INSERT INTO tiers (tier_name) VALUES('Bronze');
INSERT INTO tiers (tier_name) VALUES('Silver');
INSERT INTO tiers (tier_name) VALUES('Gold');
INSERT INTO tiers (tier_name) VALUES('Platinum');

INSERT INTO locations (location_name) VALUES('Jakarta');
INSERT INTO locations (location_name) VALUES('Medan');
INSERT INTO locations (location_name) VALUES('Bandung');

INSERT INTO users (user_name, user_location, user_tier) VALUES('Jono', 1, 1);
INSERT INTO users (user_name, user_location, user_tier) VALUES('Christy', 2, 2);
INSERT INTO users (user_name, user_location, user_tier) VALUES('Ellena Gomez', 3, 3);
INSERT INTO users (user_name, user_location, user_tier) VALUES('Latifa', 3, 4);
INSERT INTO users (user_name, user_location, user_tier) VALUES('Castley', 2, 3);
INSERT INTO users (user_name, user_location, user_tier) VALUES('Sarimi', 1, 2);
INSERT INTO users (user_name, user_location, user_tier) VALUES('Kokuri Yamate', 1, 1);