BEGIN;

CREATE TABLE "subjects" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "teacher_id" INT NOT NULL
);

ALTER TABLE "subjects"
ADD CONSTRAINT "fk__id__teachers.id"
FOREIGN KEY ("teacher_id")
REFERENCES "teachers" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

CREATE UNIQUE INDEX "uidx__subjects__name"
ON "subjects" ("name", "teacher_id");

COMMIT;