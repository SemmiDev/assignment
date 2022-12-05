package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	err := u.db.Create(&teacher).Error
	return err
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	listTeacher := []model.Teacher{}
	err := u.db.Where("deleted_at IS NULL").Find(&listTeacher).Error
	viewTeacher := []model.ViewTeacher{}
	for _, v := range listTeacher {
		viewTeacher = append(viewTeacher, model.ViewTeacher{
			Name:         v.Name,
			FieldOfStudy: v.FieldOfStudy,
			Age:          v.Age,
		})
	}
	return viewTeacher, err
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	err := u.db.
		Model(&model.Teacher{}).
		Where("id = ?", id).
		Update("name", name).Error
	return err
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	err := u.db.Delete(&model.Teacher{}, id).Error
	return err
}
