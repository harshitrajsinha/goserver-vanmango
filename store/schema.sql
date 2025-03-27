-- Create ENUM type for van category
-- CREATE TYPE IF NOT EXISTS category AS ENUM ('simple', 'rugged', 'luxury');

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'category') THEN
        CREATE TYPE category AS ENUM ('simple', 'rugged', 'luxury');
    END IF;
END $$;

-- Create ENUM type for engine material
-- CREATE TYPE IF NOT EXISTS engine_material AS ENUM ('aluminium', 'iron');

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'engine_material') THEN
        CREATE TYPE engine_material AS ENUM ('aluminium', 'iron');
    END IF;
END $$;

-- Create ENUM type for cylinders
-- CREATE TYPE cylinders AS ENUM ('4', '6', '8');

-- Create table engine
CREATE TABLE IF NOT EXISTS engine (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    displacement_in_cc INT NOT NULL,
    no_of_cylinders INT NOT NULL,
    material  engine_material,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create table van
CREATE TABLE IF NOT EXISTS van (
    van_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category category NOT NULL,
    fuel_type VARCHAR(50) NOT NULL,
    engine_id UUID NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_engine_id FOREIGN KEY (engine_id) REFERENCES engine(id) ON DELETE CASCADE
);

-- Create or replace trigger function for updating updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$ 
BEGIN
    IF EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'engine_set_timestamp') THEN
        DROP TRIGGER engine_set_timestamp ON engine;
    END IF;
END $$;

-- Apply trigger to engine table
CREATE TRIGGER engine_set_timestamp
BEFORE UPDATE ON engine
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

DO $$ 
BEGIN
    IF EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'van_set_timestamp') THEN
        DROP TRIGGER van_set_timestamp ON van;
    END IF;
END $$;

-- Apply trigger to van table
CREATE TRIGGER van_set_timestamp
BEFORE UPDATE ON van
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Clear existing data before inserting new data
TRUNCATE TABLE van CASCADE;
TRUNCATE TABLE engine CASCADE;

-- Insert data into the engine table
INSERT INTO engine (id, displacement_in_cc, no_of_cylinders, material)
VALUES
    ('e1f86b1a-0873-4c19-bae2-fc60329d0140', 2000, 4, 'aluminium'),
    ('f4a9c66b-8e38-419b-93c4-215d5cefb318', 1600, 4, 'iron'),
    ('cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c', 3000, 6, 'aluminium'),
    ('9746be12-07b7-42a3-b8ab-7d1f209b63d7', 1800, 4, 'aluminium');

-- Insert data into the van table
INSERT INTO van (name, brand, description, category, fuel_type, engine_id, price)
VALUES
    ('Modest Explorer', 'Honda', 'The Modest Explorer is a van designed to get you out of the house and into nature. This beauty is equipped with solar panels, a composting toilet, a water tank and kitchenette. The idea is that you can pack up your home and escape for a weekend or even longer!', 'simple', 'Gasoline', 'e1f86b1a-0873-4c19-bae2-fc60329d0140', 5136.00),
    ('Beach Bum', 'Volkswagen', 'Beach Bum is a van inspired by surfers and travelers. It was created to be a portable home away from home, but with some cool features in it you wont find in an ordinary camper.','rugged','Petrol', 'f4a9c66b-8e38-419b-93c4-215d5cefb318', 6849.04),
    ('Reliable Red', 'Toyota', 'Reliable Red is a van that was made for traveling. The inside is comfortable and cozy, with plenty of space to stretch out in. There is a small kitchen, so you can cook if you need to. Youll feel like home as soon as you step out of it.','luxury', 'Diesel', 'cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c', 8561.30),
    ('Dreamfinder', 'Nissan', 'Dreamfinder is the perfect van to travel in and experience. With a ceiling height of 2.1m, you can stand up in this van and there is great head room. The floor is a beautiful glass-reinforced plastic (GRP) which is easy to clean and very hard wearing. A large rear window and large side windows make it really light inside and keep it well ventilated.','simple','Gasoline', '9746be12-07b7-42a3-b8ab-7d1f209b63d7', 5564.84),
    ('The Cruiser', 'Mercedes-Benz', 'The Cruiser is a van for those who love to travel in comfort and luxury. With its many windows, spacious interior and ample storage space, the Cruiser offers a beautiful view wherever you go.','luxury','Diesel', 'cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c', 10273.36),
    ('Green Wonder', 'Ford', 'With this van, you can take your travel life to the next level. The Green Wonder is a sustainable vehicle that is perfect for people who are looking for a stylish, eco-friendly mode of transport that can go anywhere.','rugged','Gasoline', '9746be12-07b7-42a3-b8ab-7d1f209b63d7', 5992.91);