-- db/migrations/products_nnn.sql

DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(200),
    description TEXT,
    ean VARCHAR(13),
    price_out REAL
);

INSERT INTO products (name, description, ean, price_out) VALUES ('Epler', 'Friske, saftige røde epler.', '1234567890123', 29.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Bananer', 'Gule, modne bananer.', '2345678901234', 19.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Melk', 'Helmelk 1 liter, pasteurisert.', '3456789012345', 14.50);
INSERT INTO products (name, description, ean, price_out) VALUES ('Brød', 'Grovbrød med solsikkefrø.', '4567890123456', 29.00);
INSERT INTO products (name, description, ean, price_out) VALUES ('Egg', 'Økologiske egg, 12 stk.', '5678901234567', 39.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Ost', 'Norvegia, halvfast hvitost 500g.', '6789012345678', 79.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Smør', 'Tine usaltet smør 250g.', '7890123456789', 34.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Poteter', 'Norske mandelpoteter, 1 kg.', '8901234567890', 25.00);
INSERT INTO products (name, description, ean, price_out) VALUES ('Laks', 'Fersk laks, fileter 400g.', '9012345678901', 89.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Kyllingfilet', 'Kyllingbrystfilet, 500g.', '1123456789012', 64.50);
INSERT INTO products (name, description, ean, price_out) VALUES ('Spaghetti', 'Fullkorn spaghetti, 500g.', '2234567890123', 19.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Tomatsaus', 'Hjemmelaget tomatsaus på glass.', '3345678901234', 24.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Appelsiner', 'Saftige appelsiner fra Spania.', '4456789012345', 32.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Yoghurt', 'Naturell yoghurt, 1 liter.', '5567890123456', 19.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Havregryn', 'Grovkornet havregryn, 1 kg.', '6678901234567', 14.90);
