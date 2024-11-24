package password

go mod download github.com/Azure/go-ansiterm

import "github.com/baldavenger78/password/password"
func GeneratePlainTextPassword() (string, error) {
	pass, err := password.Generate(10, 4, 0, false, false)
	if err != nil {
		return "", err
	}

	return pass, nil
}

func GenerateRandomString() (string, error) {
	pass, err := password.Generate(16, 8, 0, true, true)
	if err != nil {
		return "", err
	}

	return pass, nil
}
