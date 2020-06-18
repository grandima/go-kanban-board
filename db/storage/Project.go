package storage

import (
	"database/sql"
	"github.com/grandima/go-task-management-app/db/entity"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("Project not found")
	ErrInvalidId = errors.New("Invalid Id")
)

type ProjectStorage interface {
	Index() ([]entity.Project, error)
	ProjectBy(Id int) (*entity.Project, error)
}

type ProjectDB struct {
	db *sql.DB
}

func NewProjectService(db *sql.DB) *ProjectDB {
	return &ProjectDB{
		db: db,
	}
}

func (p *ProjectDB) Index() ([]entity.Project, error) {
	db := p.db
	const q = `SELECT * FROM project`

	rows, err := db.Query(q)
	if err != nil {
		return nil, errors.Wrap(err, "selecting projects")
	}
	defer rows.Close()

	projects := []entity.Project{}

	for rows.Next() {
		var p entity.Project
		err = rows.Scan(&p.Id, &p.Name, &p.Description)
		if err != nil {
			return nil, errors.Wrap(err, "mapping entity")
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (p *ProjectDB) ProjectBy(id int) (*entity.Project, error) {
	var project entity.Project

	const q = `SELECT * FROM project WHERE id = $1`
	row := p.db.QueryRow(q, id)

	if err := row.Scan(&project.Id, &project.Name, &project.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, errors.Wrap(err, "selecting single product")
	}

	return &project, nil
}
