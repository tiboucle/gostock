INSERT INTO colors (code, name) VALUES
('BWH', 'Broken White'),
('NAV', 'Navy'),
('BLA', 'Black'),
('SAL', 'Salem'),
('YEL', 'Yellow'),
('WHI', 'White'),
('KHA', 'Khaki'),
('RED', 'Red');

INSERT INTO items(sku, name, size, color, stock) VALUES
('SSI-D00791077-MM-BWH', 'Zalekia Plain Casual Blouse', 'M', 'BWH', 10),
('SSI-D00864612-LL-NAV', 'Deklia Plain Casual Blouse', 'L', 'NAV', 10),
('SSI-D01037812-X3-BLA', 'Dellaya Plain Loose Big Blouse', 'XXXL', 'BLA', 10),
('SSI-D01220307-XL-SAL', 'Devibav Plain Trump Blouse', 'XL', 'SAL', 10),
('SSI-D01220357-SS-YEL', 'Devibav Plain Trump Blouse', 'S', 'YEL', 10),
('SSI-D01322234-LL-WHI', 'Thafqya Plain Raglan Blouse', 'L', 'WHI', 10),
('SSI-D01326201-XL-KHA', 'Siunfhi Ethnic Trump Blouse', 'XL', 'KHA', 10),
('SSI-D01401071-LL-RED', 'Zeomila Zipper Casual Blouse', 'L', 'RED', 10);

INSERT INTO purchases(invoice_id, invoice_date) VALUES
('20170823-75140', '2017-08-23 8:14:00');

INSERT INTO purchase_details(id, invoice_id, item_sku, price, amount) VALUES
(1, '20170823-75140', 'SSI-D00791077-MM-BWH', 64000, 56);

INSERT INTO incoming_items(id, invoice_id, item_sku, note, item_received, time_received) VALUES
(1, '20170823-75140', 'SSI-D00791077-MM-BWH', '2017/08/27 terima 51; 2017/08/29 terima 5', 56, '2017-08-23 8:14:00');

INSERT INTO orders(invoice_id, invoice_date) VALUES
('ID-20180107-267724', '2018-01-07 10:14:42'),
('ID-20180108-267724', '2018-01-08 10:14:42');

INSERT INTO order_details(id, invoice_id, item_sku, price, amount) VALUES
(1, 'ID-20180107-267724', 'SSI-D00791077-MM-BWH', 115000, 2),
(2, 'ID-20180108-267724', 'SSI-D00791077-MM-BWH', 100000, 3);