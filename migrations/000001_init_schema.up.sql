CREATE TYPE user_role AS ENUM ('employee', 'moderator');
CREATE TYPE city AS ENUM ('Москва', 'Санкт-Петербург', 'Казань');
CREATE TYPE reception_status AS ENUM ('in_progress', 'close');
CREATE TYPE product_type AS ENUM ('электроника', 'одежда', 'обувь');

CREATE TABLE "users" (
                        id UUID PRIMARY KEY,
                        email VARCHAR(255) UNIQUE NOT NULL CHECK("email" LIKE '%@%'),
                        password_hash VARCHAR(255) NOT NULL,
                        role user_role NOT NULL
);

CREATE TABLE "pvz" (
                       id UUID PRIMARY KEY,
                       "registrationDate" TIMESTAMP NOT NULL DEFAULT NOW(),
                       city city NOT NULL,
                       "moderatorId" UUID NOT NULL REFERENCES "users"(id)
);

CREATE TABLE "reception" (
                             id UUID PRIMARY KEY,
                             "dateTime" TIMESTAMP NOT NULL DEFAULT NOW(),
                             "pvzId" UUID NOT NULL REFERENCES "pvz"(id),
                             status reception_status NOT NULL DEFAULT 'in_progress',
                             "employeeId" UUID NOT NULL REFERENCES "users"(id)
);

CREATE TABLE "product" (
                           id UUID PRIMARY KEY,
                           "dateTime" TIMESTAMP NOT NULL DEFAULT NOW(),
                           type product_type NOT NULL,
                           "receptionId" UUID NOT NULL REFERENCES "reception"(id),
                           sequence_number INTEGER NOT NULL
);
