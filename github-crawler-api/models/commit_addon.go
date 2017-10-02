package models

import (
	"context"
	"encoding/json"
	"fmt"
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
	ParentSha  string
	Lat        float64
	Lon        float64
	LocationID int
}

type ParentCommitList []ParentCommit

type ParentCommit struct {
	Sha string
}

// ListCommitWithUsersWithLocationForRepo returns an array of view: default.
func (m *CommitDB) ListCommitWithUsersWithLocationForRepo(ctx context.Context, repoID int, from *time.Time, till *time.Time, limit int) []*app.Commit {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "listcommit"}, time.Now())
	var objs []*app.Commit

	// We want to filter on time to not crash the system
	if from == nil || till == nil {
		return objs
	}

	var native []*CommitWithEverything
	err := m.Db.Scopes().Table("repositories").
		Select(`users.*, location_id, commit_date, sha, "Commits".raw#>'{parents}' as parent, ST_Y(point) as lat, ST_X(point) as lon`).
		Joins(`LEFT JOIN "Commits" on repository_id = repositories.project_id`).
		Joins(`LEFT JOIN users on github_user_id = author_id`).
		Joins(`LEFT JOIN locations on locations.id = location_id`).
		Where("project_id = ? AND commit_date > ? AND commit_date < ? AND ST_Y(point) != 0 ", repoID, from, till).
		Order("commit_date desc").
		Limit(limit).
		Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Commit", "error", err.Error())
		return objs
	}
	for _, t := range native {
		var pList ParentCommitList
		err := json.Unmarshal(t.Parent, &pList)
		if err != nil {
			fmt.Println("Unmarshall went wrong for parentcommit list", err)
		}
		if len(pList) > 0 {
			t.ParentSha = pList[0].Sha
		}
		fmt.Println("Current sha: ", t.Sha)
		fmt.Println("Parent sha: ", t.ParentSha)
		fmt.Println("loc", t.Lat)
		fmt.Println("location id", t.LocationID)
		// objs = append(objs, t.CommitToCommit())
	}

	return objs
}
