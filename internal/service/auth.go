package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type AuthService struct{}

type Account struct {
	Username string
	Password string
	Salt     string
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func init() {
	db, err := Open()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(strings.TrimSpace(`
		create table Account(
			username varchar(25),
			password varchar(255),
			salt varchar(50),
			primary key (username)
		);
	`))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer stmt.Close()

	if _, err = stmt.Exec(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}

func (s *AuthService) Create(data *Account) error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into Account values(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	salt := genSalt()
	if _, err = stmt.Exec(data.Username, encrypt(data.Password, salt), salt); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Read(username string) (*Account, error) {
	db, err := Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from Account where username = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var account Account
	if err := stmt.QueryRow(username).Scan(&account.Username, &account.Password, &account.Salt); err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AuthService) Update(username, password string) error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("update Account set password = ?, salt = ? where username = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	salt := genSalt()
	if _, err = stmt.Exec(encrypt(password, salt), salt, username); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Delete(username string) error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from Account where username = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(username); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Verify(username, password string) (bool, error) {
	account, err := s.Read(username)
	if err != nil {
		return false, err
	}

	if encrypt(password, account.Salt) == account.Password {
		return true, nil
	}

	return false, nil
}

func (s *AuthService) VerifyToken(username, encryptPw string) (bool, error) {
	account, err := s.Read(username)
	if err != nil {
		return false, err
	}

	if encryptPw == account.Password {
		return true, nil
	}

	return false, nil
}

func (s *AuthService) Token(username, password string) string {
	raw := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

func encrypt(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

func genSalt() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return ""
	}
	return hex.EncodeToString(b)
}
