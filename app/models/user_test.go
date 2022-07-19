package models

import "fmt"

type test struct {
	input  []User
	output string
}

func (ms *ModelSuite) Test_CreateUser() {

	tests := []test{
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "User",
					Email:                "",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Email can not be blank",
		},
		{
			input: []User{
				{
					FirstName:            "",
					LastName:             "User",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "First Name can not be blank",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Last Name can not be blank",
		},
		{
			input: []User{
				{
					FirstName:            "TestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTest",
					LastName:             "user",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "First Name must be less than 50 characters.",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "TestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTestTest",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Last Name must be less than 50 characters.",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password2",
				},
			},
			output: "Passwords do not match.",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "test@gmail.com",
					Password:             "",
					PasswordConfirmation: "password2",
				},
			},
			output: "Password is required.",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "test@gmail.com",
					Password:             "123",
					PasswordConfirmation: "password2",
				},
			},
			output: "Password must be between 8 and 50 characters.",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test12",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Last Name must be letters only.",
		},
		{
			input: []User{
				{
					FirstName:            "Test12",
					LastName:             "Test",
					Email:                "test@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "First Name must be letters only.",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "testdsdsdsdsdsdsdsdsdsdsdsdssddsdsssdsdsdsdsdsdsdsdsdsdsddsdsdsdsdsdsdsdsdsdsdsd@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Before @ Email must be less or equal than 64 characters",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "test@gmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddedededgmailsdsddsdsdddddddededed.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "After @ Email must be less or equal than 255 characters",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "'@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Email is invalid",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "----@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Email is invalid",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Email is invalid",
		},
		{
			input: []User{
				{
					FirstName:            "Test",
					LastName:             "Test",
					Email:                "'@gmail.com",
					Password:             "password",
					PasswordConfirmation: "password",
				},
			},
			output: "Email is invalid",
		},
	}

	for i, test := range tests {
		for _, user := range test.input {
			fmt.Println("Testing: ", i)

			err, err2 := user.ValidateCreate(ms.DB)
			err.Append(err2)
			ms.Error(err)
			ms.True(err.HasAny())
			fmt.Println(test.output)
			ms.Contains(err.Error(), test.output)
		}
	}
}
