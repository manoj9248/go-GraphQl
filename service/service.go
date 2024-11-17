package service

import (
	"go-graphql/graph/model"

	"github.com/google/uuid"
)

var DBMap = make(map[string]model.JobListing)

func GetJob(id string) *model.JobListing {
	resp := DBMap[id]
	if resp.ID == "" {
		return nil
	}
	return &resp
}
func DeleteJob(id string) *model.DeleteJobResponse {
	delete(DBMap, id)
	return &model.DeleteJobResponse{
		DeleteJobID: id,
	}
}
func CreateJob(input model.CreateJobListingInput) *model.JobListing {
	id := uuid.New()
	DBMap[id.String()] = model.JobListing{
		ID:          id.String(),
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		URL:         input.URL,
	}
	resp := DBMap[id.String()]
	return &resp
}
func UpdateJob(id string, input model.UpdateJobListingInput) *model.JobListing {
	getData := DBMap[id]
	if input.Title == nil {
		input.Title = &getData.Title
	}
	if input.Description == nil {
		input.Description = &getData.Description
	}
	if input.Company == nil {
		input.Company = &getData.Company
	}
	if input.URL == nil {
		input.URL = &getData.URL
	}
	DBMap[id] = model.JobListing{
		ID:          id,
		Title:       *input.Title,
		Description: *input.Description,
		Company:     *input.Company,
		URL:         *input.URL,
	}
	resp := DBMap[id]
	return &resp
}
func GetJobs() []*model.JobListing {
	jobs := make([]*model.JobListing, 0, len(DBMap))
	for _, job := range DBMap {
		jobs = append(jobs, &job)
	}
	return jobs
}
