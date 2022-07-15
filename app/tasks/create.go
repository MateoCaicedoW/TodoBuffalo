package tasks

import (
	"TodoBuffalo/app/models"
	"errors"
	"time"

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
		if err := models.DB().Create(&task); err != nil {
			return err
		}
	}
	return nil
})

var _ = grift.Add("create:users", func(c *grift.Context) error {

	for i := 0; i < 30; i++ {
		var user models.User
		fako.Fill(&user)
		if err := models.DB().Create(&user); err != nil {
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
		return errors.New("User already exists=> Email: " + user.Email + " Password: " + user.Password)
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.PasswordHash = string(pass)
	if err := tx.Create(user); err != nil {
		return err
	}
	return nil
})
