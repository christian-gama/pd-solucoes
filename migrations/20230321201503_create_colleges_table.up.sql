BEGIN;

CREATE TABLE "colleges" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "cnpj" VARCHAR(14) NOT NULL
);

CREATE UNIQUE INDEX "uidx__colleges__cnpj"
ON "colleges" ("cnpj");

COMMIT;
