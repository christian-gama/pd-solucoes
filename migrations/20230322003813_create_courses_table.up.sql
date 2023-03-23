BEGIN;

CREATE TABLE "courses" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "college_id" INT NOT NULL
);

ALTER TABLE "courses"
ADD CONSTRAINT "fk__id__college.id"
FOREIGN KEY ("college_id")
REFERENCES "colleges" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

CREATE UNIQUE INDEX "uidx__courses__name"
ON "courses" ("name", "college_id");

COMMIT;