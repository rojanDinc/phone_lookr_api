package util

import (
	"os"
	"testing"
)

type fileReaderMock struct {
	values string
}

func (frm fileReaderMock) ReadFile() ([]byte, error) {
	return []byte(frm.values), nil
}

func TestToLoadEnvVarsIntoApp(t *testing.T) {
	// Arrange
	values := `MY_KEY=1234
DB_USER=test
DB_PASSWORD=mysecretpassword`
	fileReader := fileReaderMock{
		values: values,
	}
	os.Clearenv()
	// Act
	err := LoadDotEnvFile(fileReader)
	dbUser := os.Getenv("DB_USER")
	// Assert
	if err != nil {
		t.Errorf("Unexpected error when loading env file: %v", err)
	}
	if dbUser != "test" {
		t.Errorf("Unexpected value for DB_USER env variable. Expected test got: %s", dbUser)
	}
}
