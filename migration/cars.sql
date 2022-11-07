CREATE TABLE "cars"(
    "id" serial Primary key,
    "image_url" VARCHAR(255),
    "marka" VARCHAR(255) NOT NULL,
    "model" VARCHAR(255) NOT NULL,
    "color" VARCHAR(255) NOT NULL,
    "mileage_km" INTEGER NOT NULL,
    "made_year" DATE NOT NULL,
    "cost" DECIMAL(18, 2) NOT NULL,
    "created_at" TIMESTAMP default current_timestamp
);