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
	var user models.User
	fako.Fill(&user)
	user.Rol = "user"
	hash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	user.PasswordHash = string(hash)
	if err := models.DB().Eager().Create(&user); err != nil {
		return err
	}

	for i := 0; i < 26; i++ {
		var task models.Task
		fako.Fill(&task)
		task.Must = time.Now()
		task.Status = false
		task.UserID = user.ID
		task.User = &user
		if err := models.DB().Eager().Create(&task); err != nil {
			return err
		}
	}
	return nil
})

var _ = grift.Add("create:users", func(c *grift.Context) error {

	for i := 0; i < 2; i++ {
		var user models.User
		fako.Fill(&user)
		user.Rol = "user"
		hash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
		user.PasswordHash = string(hash)
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
