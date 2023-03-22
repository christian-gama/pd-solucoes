BEGIN;

CREATE TABLE "course_enrollments" (
  "id" SERIAL PRIMARY KEY,
  "student_id" INT NOT NULL,
  "course_subject_id" INT NOT NULL,
  "enrollment_date" DATE NOT NULL
);

CREATE UNIQUE INDEX "uidx__course_enrollments__student_id__course_subject_id"
ON "course_enrollments" ("student_id", "course_subject_id");

CREATE INDEX "idx__course_enrollments__student_id"
ON "course_enrollments" ("student_id");

CREATE INDEX "idx__course_enrollments__course_subject_id"
ON "course_enrollments" ("course_subject_id");

ALTER TABLE "course_enrollments"
ADD CONSTRAINT "fk__student_id__students.id"
FOREIGN KEY ("student_id") 
REFERENCES "students" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "course_enrollments"
ADD CONSTRAINT "fk__course_subject_id__course_subjects.id"
FOREIGN KEY ("course_subject_id")
REFERENCES "course_subjects" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;
