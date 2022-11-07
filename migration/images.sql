CREATE TABLE "car_images"(
    "id" serial PRIMARY KEY,
    "cars_id" INTEGER NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "sequence_number" INTEGER PRIMARY KEY
);