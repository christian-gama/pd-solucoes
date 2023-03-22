BEGIN;

CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "cpf" VARCHAR(11) NOT NULL
);

CREATE UNIQUE INDEX "uidx__students__cpf"
ON "students" ("cpf");

ALTER TABLE "students"
ADD CONSTRAINT "chk__cpf__must_be_valid_cpf"
CHECK ("cpf" ~ '^[0-9]{11}$');

COMMIT;