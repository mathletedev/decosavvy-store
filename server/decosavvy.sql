CREATE TABLE users (
	id UUID PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	avatar TEXT NOT NULL,
	cart UUID[] DEFAULT '{}',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
	id UUID PRIMARY KEY,
	image TEXT NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	price FLOAT NOT NULL
);

INSERT INTO products (id, image, name, description, price) VALUES (
	'9a6f8e95-7321-46ad-b5eb-43f8c6ae04a2',
	'https://lh3.googleusercontent.com/d/16T3H_mCuxyrEe-a9F4p40ie5QPo3lT16',
	'Steel Table',
	'This steel table is masterfully crafted for strength, durability, and longevity. With its minimalistic design, a steel table offers both functionality and contemporary appeal.',
	379.99
);

INSERT INTO products (id, image, name, description, price) VALUES (
	'246f1f47-ffdb-48da-9e24-93d7cba22e08',
	'https://lh3.googleusercontent.com/d/1vyQsDU7UVTwqTkHqQtGF4_hJwhzEwl8b',
	'Wooden Chair',
	'This wooden chair comes from the forests of Buffalo New York, this product was harvested and processed in the home country of the United States of America. This wooden chair comes with a nice sleek exterior, finished with various polished to ensure quality.',
	79.99
);

INSERT INTO products (id, image, name, description, price) VALUES (
	'6ab4b53e-4a44-4047-b91b-2fda38d2af11',
	'https://lh3.googleusercontent.com/d/1R10IukI43y4RLwodyG6e1eQ2XxUSS0Kf',
	'Large Closet',
	'These large closets are designed to maximize storage efficiency and keep your belongings organized and easily accessible. With spacious storage, the design allows for various accessories and features.',
	379.99
);

INSERT INTO products (id, image, name, description, price) VALUES (
	'f7e1e019-8103-4cc7-a494-7237b5ccec91',
	'https://lh3.googleusercontent.com/d/1rSw884KOQIob2j6Zl2WQAy1ZGKSwP8KS',
	'Double Seat',
	'This high end couch starts in the forests of New Zealand. With dye from white clovers, and white rara, this couch is made with the most pristine resources. Our couch also comes equipped with memory foam, the most popular mattress on the market.',
	639.99
);

INSERT INTO products (id, image, name, description, price) VALUES (
	'dd53fb22-83ad-46ab-8a51-46473aad2888',
	'https://lh3.googleusercontent.com/d/1qSChclm-xpY-VhH5xJjsFgsxAQDgE_ur',
	'Bistro Table',
	'A bistro table is a small, typically round table designed for casual dining. With their compact size, and homey aesthetic, bistro tables create an intimate and inviting atmosphere.',
	119.99
);

INSERT INTO products (id, image, name, description, price) VALUES (
	'ecdd7349-4957-4956-b38b-8beb18730757',
	'https://lh3.googleusercontent.com/d/1I0--EW9gIDu9vpTfxjycRRwrUoojCHBn',
	'Large Sofa',
	'Our large sofa is our finest product, with its Egyptian Blue pigment, its radiating blue is is simply immaculate. To experience the simple American life style while still adhering to luxury, we have hidden silk linings and gold patterns to make sure you feel that luxury. While average on the outside, our luxurious furniture is the best on the market bosting luxury status while adhering to affordable prices.',
	959.99
);
