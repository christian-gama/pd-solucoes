package seed

import (
	"context"
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
	"github.com/christian-gama/pd-solucoes/pkg/log"
	"gorm.io/gorm"
)

func Execute(ctx context.Context, log log.Logger) {
	db := sql.MakePostgres()

	log.Infof("Seeding the database. It won't run if the database is not empty.")

	db.Transaction(func(tx *gorm.DB) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Seeding failed, aborting transaction:\n%v", r)
				err = fmt.Errorf("%v", r)
			}
		}()

		log.Infof("Seeding colleges")
		College(ctx, tx)

		log.Infof("Seeding courses")
		Course(ctx, tx)

		log.Infof("Seeding teachers")
		Teacher(ctx, tx)

		log.Infof("Seeding students")
		Student(ctx, tx)

		log.Infof("Seeding subjects")
		Subject(ctx, tx)

		log.Infof("Seeding course subjects")
		CourseSubject(ctx, tx)

		log.Infof("Seeding course enrollments")
		CourseEnrollment(ctx, tx)

		log.Infof("Seeding completed")

		return
	})
}
