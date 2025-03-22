package service

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

type PrivDirService struct{}

type PrivDir struct {
	Id      string `json:"id"`
	DirName string `json:"dirname"`
	Owner   string `json:"owner"`
}

func init() {
	db, err := Open()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(strings.TrimSpace(`
		create table PrivDir(
		    id varchar(36),
		    dirname varchar(250) unique,
			owner varchar(25),
			constraint PK_PrivDir_ID primary key(id),
			constraint FK_Owner_ID foreign key(owner)
			references(Account.username) on update cascade on delete cascade
		);
	`))
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.Exec(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}

func NewPrivDirService() *PrivDirService {
	return &PrivDirService{}
}

func (sv *PrivDirService) CreatePriv(dirname string, acc *Account) error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	id := uuid.NewString()
	stmt, err := db.Prepare("insert into PrivDir(id, name, owner) values (?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, dirname, acc.Username)
	return nil
}

func (sv *PrivDirService) ReadPriv(name string) (*PrivDir, error) {
	db, err := Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from PrivDir where name = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(name)
	var data PrivDir

	if err = row.Scan(&data.Id, &data.DirName, &data.Owner); err != nil {
		return nil, err
	}

	return &data, nil
}
