package entities

import (
	"strings"

	"gorm.io/gorm"
)

type IRepoWithState[E Entity, S IState] interface {
	FindByStates(map[string][]int) ([]E, error)
}

type RepoWithState[E Entity, S IState] struct {
	Entity E
	State  S
	db     *gorm.DB
}

func NewRepoWithStates[E Entity, S IState](db *gorm.DB, entity E, states S) *RepoWithState[E, S] {
	return &RepoWithState[E, S]{
		db:     db,
		Entity: entity,
		State:  states,
	}
}

func (rws RepoWithState[E, S]) FindByStates(states map[string][]int) ([]E, error) {
	var entities []E
	var entity E // to declare the type

	repo := NewRepository(rws.db, entity)
	var result *gorm.DB
	queryStrBuilder := strings.Builder{}
	queryString, params := repo.queryBuilder(entity)

	queryStr, stateParams := rws.createQueryForStates(states)

	if queryStrBuilder.String() != "" {
		queryStrBuilder.WriteString(" AND ")
	}
	queryStrBuilder.WriteString(queryStr)

	if queryString != "" {
		result = rws.db.Where(queryString, params).Where(queryStrBuilder.String(), stateParams).Find(&entities)
	} else {
		result = rws.db.Where(queryStrBuilder.String(), stateParams).Find(&entities)
	}
	return entities, result.Error
}

func (repo RepoWithState[E, S]) createQueryForStates(states map[string][]int) (string, []int) {
	queryStr := strings.Builder{}
	params := []int{}
	counter := 0
	for k, v := range states {
		if counter > 0 {
			queryStr.WriteString(" AND ")
		}
		queryStr.WriteString(turnStateTypeToQueryStr(k))
		queryStr.WriteString(" IN ?")
		for _, i := range v {
			params = append(params, i)
		}
		counter++
	}
	return queryStr.String(), params
}

func turnStateTypeToQueryStr(stateType string) string {
	// todo: extend for all states
	switch strings.ToLower(stateType) {
	case "taskstate", "task":
		return "task_state_id"
	case "tasktimestate", "tasktime":
		return "task_time_state_id"
	case "projectstate", "project":
		return "project_state_id"
	case "projecttimestate", "projecttime":
		return "project_time_id"
	case "batchstate", "batch":
		return "batch_id"
	case "batchtimestate", "batchtime":
		return "batch_time_state_id"
	case "clientofferstate", "clientoffer":
		return "client_offer_state_id"
	default:
		return ""
	}

}
