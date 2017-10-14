package models

import (
	"context"

	"github-crawler-api/app"

	"time"

	"github.com/goadesign/goa"
)

// MediaType Retrieval Functions
type CommitWithEverything struct {
	User
	CommitDate time.Time
	Parent     []byte
	Sha        string
	Message    string
	ParentSha  string
	Lat        float64
	Lng        float64
	LocationID int
}

type ParentCommitList []ParentCommit

type ParentCommit struct {
	Sha string
}

// ListCommitWithUsersWithLocationForRepo returns an array of view: default.
func (m *CommitDB) ListCommitWithUsersWithLocationForRepo(ctx context.Context, repoID int, from *time.Time, till *time.Time, limit *int) app.CommitCollection {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "listcommit"}, time.Now())
	var objs []*app.Commit
	l := 2000

	// We want to filter on time to not crash the system
	if from == nil || till == nil {
		return objs
	}
	if limit == nil {
		limit = &l
	}

	var native []*CommitWithEverything
	err := m.Db.Scopes().Table("repositories").
		Select(`users.login,users.type, location_id, message,commit_date, sha, ST_X(point) as lat, ST_Y(point) as lng`).
		Joins(`LEFT JOIN "Commits" on repository_id = repositories.project_id`).
		Joins(`LEFT JOIN users on github_user_id = author_id`).
		Joins(`LEFT JOIN locations on locations.id = location_id`).
		Where("repositories.id = ? AND commit_date > ? AND commit_date < ? AND ST_Y(point) != 0 ", repoID, from, till).
		Order("commit_date desc").
		Limit(limit).
		Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Commit", "error", err.Error())
		return objs
	}

	var list app.CommitCollection
	for _, t := range native {
		commit := app.Commit{}
		commit.Sha = t.Sha
		commit.Message = &t.Message
		loc := &app.Location{Lat: t.Lat, Lng: t.Lng}
		commit.Author = &app.Ghuser{ID: t.User.ID, Location: loc, Login: t.User.Login, Type: t.User.Type}
		commit.Timestamp = t.CommitDate
		list = append(list, &commit)
	}

	return list
}
