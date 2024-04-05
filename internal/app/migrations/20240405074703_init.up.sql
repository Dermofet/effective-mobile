CREATE TABLE IF NOT EXISTS "cars" (
    "id" UUID NOT NULL PRIMARY KEY,
    "reg_num" VARCHAR(9) NOT NULL,
    "mark" VARCHAR(255) NOT NULL,
    "model" VARCHAR(255) NOT NULL,
    "year" INTEGER NOT NULL,
    "owner_name" VARCHAR(255) NOT NULL,
    "owner_surname" VARCHAR(255) NOT NULL,
    "owner_patronymic" VARCHAR(255)
);