BEGIN;

-- module to generate UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create enum for group_size
DO $$ BEGIN
    CREATE TYPE GROUP_SIZE AS ENUM('2', '4');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS church(
    id SERIAL PRIMARY KEY,
    pub_id UUID DEFAULT uuid_generate_v4() NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    address TEXT NOT NULL,
    website TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone TEXT NOT NULL,
    group_size GROUP_SIZE,
    same_gender BOOLEAN,
    min_age SMALLINT CHECK (min_age > 0),
    member_basic_info_update BOOLEAN NOT NULL,
    active BOOLEAN NOT NULL
);

-- create enums for gender and role
DO $$ BEGIN
    CREATE TYPE GENDER AS ENUM('male', 'female');
    CREATE TYPE ROLE AS ENUM ('member', 'leader', 'admin');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS app_user(
    id SERIAL PRIMARY KEY,
    pub_id UUID DEFAULT uuid_generate_v4() NOT NULL,
    church_id INTEGER REFERENCES church(id) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    gender GENDER NOT NULL,
    date_of_birth DATE,
    description TEXT,
    role ROLE NOT NULL,
    active BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS contact_method(
    id SERIAL PRIMARY KEY,
    pub_id UUID DEFAULT uuid_generate_v4() NOT NULL,
    user_id INTEGER REFERENCES app_user(id) NOT NULL,
    label TEXT NOT NULL,
    contact_details TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS question(
    id SERIAL PRIMARY KEY,
    pub_id UUID DEFAULT uuid_generate_v4() NOT NULL,
    church_id INTEGER REFERENCES church(id),
    question TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS match_group(
    id SERIAL PRIMARY KEY,
    pub_id UUID DEFAULT uuid_generate_v4() NOT NULL,
    created TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS match_group_app_user(
    match_group_id INTEGER REFERENCES match_group(id) NOT NULL,
    user_id INTEGER REFERENCES app_user(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS match_group_question(
    match_group_id INTEGER REFERENCES match_group(id) NOT NULL,
    question_id INTEGER REFERENCES question(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS bulletin(
    id SERIAL PRIMARY KEY,
    pub_id UUID DEFAULT uuid_generate_v4() NOT NULL,
    match_group_id INTEGER REFERENCES match_group(id) NOT NULL,
    user_id INTEGER REFERENCES app_user(id) NOT NULL,
    sent TIMESTAMPTZ NOT NULL,
    message TEXT NOT NULL
);

COMMIT;