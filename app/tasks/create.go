package tasks

import (
	"TodoBuffalo/app/models"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/markbates/grift/grift"
	"github.com/wawandco/fako"
	"golang.org/x/crypto/bcrypt"
)

var _ = grift.Add("create:task", func(c *grift.Context) error {

	for i := 0; i < 100; i++ {
		var task models.Task
		fako.Fill(&task)
		task.Must = time.Now()
		task.Status = false
		task.UserID = uuid.FromStringOrNil("8b04e3a0-853c-417e-aa19-b098e11e7123")
		if err := models.DB().Eager().Create(&task); err != nil {
			return err
		}
	}
	return nil
})

var _ = grift.Add("create:users", func(c *grift.Context) error {

	for i := 0; i < 30; i++ {
		var user models.User
		fako.Fill(&user)
		if err := models.DB().Eager().Create(&user); err != nil {
			return err
		}
	}
	return nil
})

var _ = grift.Add("create:users:admin", func(c *grift.Context) error {

	user := &models.User{
		Email:     "admin@gmail.com",
		Password:  "admin12345",
		Rol:       "admin",
		LastName:  "admin",
		FirstName: "admin",
	}
	tx := models.DB()
	userTemp := &models.User{}
	tx.Where("email = ?", user.Email).First(userTemp)

	if userTemp.Email != "" {
		return errors.New("User already exists=> Email: " + userTemp.Email + " Password: " + user.Password)
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.PasswordHash = string(pass)
	if err := tx.Eager().Create(user); err != nil {
		return err
	}
	return nil
})
