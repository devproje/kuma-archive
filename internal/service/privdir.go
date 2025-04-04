package service

import (
	"fmt"
	"git.wh64.net/devproje/kuma-archive/internal/util"
	"github.com/google/uuid"
	"os"
	"strings"
)

type PrivDirService struct {
	acc *Account
}

type PrivDir struct {
	Id      string `json:"id"`
	DirName string `json:"dirname"`
	Owner   string `json:"owner"`
}

type test interface {
	Create()
	Read()
	Update()
	Delete()
}

func init() {
	db, err := util.OpenDatabase()
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
			    references Account(username) on update cascade on delete cascade
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

func NewPrivDirService(acc *Account) *PrivDirService {
	return &PrivDirService{
		acc: acc,
	}
}

func (sv *PrivDirService) Create(dirname string) (string, error) {
	db, err := util.OpenDatabase()
	if err != nil {
		return "", err
	}
	defer db.Close()

	id := uuid.NewString()
	stmt, err := db.Prepare("insert into PrivDir(id, dirname, owner) values (?, ?, ?);")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, dirname, sv.acc.Username)
	return id, nil
}

func (sv *PrivDirService) Read(dirname string) (*PrivDir, error) {
	db, err := util.OpenDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from PrivDir where dirname = ? and owner = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(dirname, sv.acc.Username)
	var data PrivDir

	if err = row.Scan(&data.Id, &data.DirName, &data.Owner); err != nil {
		return nil, err
	}

	return &data, nil
}

func (sv *PrivDirService) Delete(dirname string) error {
	db, err := util.OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from PrivDir where dirname = ? and owner = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(dirname, sv.acc.Username)
	if err != nil {
		return err
	}

	return nil
}

func (sv *PrivDirService) Query() []PrivDir {
	db, err := util.OpenDatabase()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from PrivDir;")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	defer rows.Close()

	var dirs []PrivDir
	for rows.Next() {
		var data PrivDir
		if err = rows.Scan(&data.Id, &data.DirName, &data.Owner); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		dirs = append(dirs, data)
	}

	return dirs
}

func (sv *PrivDirService) QueryUser() []PrivDir {
	db, err := util.OpenDatabase()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from PrivDir where owner = ?;")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query(sv.acc.Username)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	defer rows.Close()

	var dirs []PrivDir
	for rows.Next() {
		var data PrivDir
		if err = rows.Scan(&data.Id, &data.DirName, &data.Owner); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		dirs = append(dirs, data)
	}

	return dirs
}
