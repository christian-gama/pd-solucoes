BEGIN;

CREATE TABLE "course_subjects" (
    "id" SERIAL PRIMARY KEY,
    "course_id" INT NOT NULL,
    "subject_id" INT NOT NULL
);

ALTER TABLE "course_subjects"
ADD CONSTRAINT "fk__course_subjects__course_id"
FOREIGN KEY ("course_id")
REFERENCES "courses" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "course_subjects"
ADD CONSTRAINT "fk__course_subjects__subject_id"
FOREIGN KEY ("subject_id")
REFERENCES "subjects" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

CREATE UNIQUE INDEX "uidx__course_subjects__course_id__subject_id"
ON "course_subjects" ("course_id", "subject_id");

COMMIT;