package models

type test struct {
	input  []User
	output string
}

// func (ms *ModelSuite) Test_CreateUser() {

// 	tests := []test{
// 		{
// 			input: []User{
// 				{
// 					FirstName:            "Test",
// 					LastName:             "User",
// 					Email:                "",
// 					Password:             "password",
// 					PasswordConfirmation: "password",
// 				},
// 			},
// 			output: "Email can not be blank",
// 		},
// 	}

// 	for _, test := range tests {
// 		for _, user := range test.input {
// 			err, _ := user.Validate(ms.DB)
// 			ms.Error(err)
// 			ms.True(err.HasAny())
// 			ms.Contains(err.Error(), test.output)
// 		}
// 	}
// }
