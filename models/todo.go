package models

import (
	"gin_todo/forms"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

const (
	Delete = 0
	Undo   = 1
	Doing  = 2
	Down   = 3
)

type Todo struct {
	Id       int64  `gorm:"primary_key" json:"id"  binding:"-"`
	Title    string `json:"title" binding:"required"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
	Status   int    `json:"status" orm:"column(status)"`
	UserId   int64  `json:"user_id" orm:"column(user_id);size(11)" binding:"-"`
}

func AddTodo(m *forms.PostTodo, user_id int64) (t *Todo, err error) {
	CreatedAt := time.Now().UTC().Unix()
	UpdatedAt := CreatedAt

	todo := Todo{
		Title:    m.Title,
		CreateAt: CreatedAt,
		UpdateAt: UpdatedAt,
		Status:   Undo,
		UserId:   user_id,
	}
	db.Create(&todo)
	flag := db.NewRecord(&todo)
	if flag == false {
		return &todo, err
	}

	return nil, err
}

func GetTodo(id int64) (*Todo, error) {
	var m Todo
	err := db.Where("id = ? ", id).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &m, nil
}

func GetAllTodo(
	query map[string]interface{},
	fields []string,
	sortby []string,
	offset int64,
	limit int64,
	userId int64) (ml *[]Todo, totalCount int64, err error) {
	var todos []Todo

	if userId > 0 {
		query["user_id"] = userId
	}

	//qs := db.Select(fields).Where(query).Offset(offset).Limit(limit).Find(&todos)
	qs := db.Where(query).Offset(offset).Limit(limit)

	//fields = []string{"title", "status", "create_at"}
	if len(fields) > 0 {
		qs = qs.Select(fields)
	}

	if len(sortby) > 0 {
		for _, field := range (sortby) {
			qs = qs.Order(field, !strings.HasPrefix(field, "-"))
		}
	}
	qs = qs.Find(&todos)

	qs.Count(&totalCount)

	err = qs.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return &todos, totalCount, err
}

func UpdateTodoById(m *Todo, post *forms.PostTodo) (err error) {
	post.UpdateAt = time.Now().UTC().Unix()
	err = db.Model(m).Updates(post).Error
	return err
}

func DeleteTodo(id int64) (err error) {
	err = db.Delete(&Todo{Id: id}).Error
	return
}
