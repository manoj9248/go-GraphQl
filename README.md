steps to create graphql project.
1. Get gql gen for your project go get github.com/99designs/gqlgen
2. Add gqlgen to tools.go printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
3. Get all the dependencies go mod tidy
4. Initialize your project go run github.com/99designs/gqlgen init
5. After you've written the graphql schema, run this - go run github.com/99designs/gqlgen generate

queries and mutations
1. Create Job
```
mutation CreateJobListing($input: createJobListingInput!) {
  createJob(input: $input) {
    _id
    title
    description
    company
    url
  }
}
```
Variable
```
{ "input": { "title": "Software Development Engineer - I", "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt", "company": "Google", "url": "www.google.com/" } }
```
GetJobs
```
query GetAllJobs{ jobs{ _id title description company url } }
```
GetJobByID
```
query GetJob($id: ID!){ job(id:$id){ _id title description url company } }
```
Variable
```
{ "id": "00da33e3-4ed1-4733-9fa4-88d516ab72e9" }
```
Update Job
```
mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) { updateJob(id:$id,input:$input){ title description _id company url } }
```
variables
```
{ "id": "00da33e3-4ed1-4733-9fa4-88d516ab72e9", "input": { "title": "Software Development Engineer - III" } }
```
DeleteJob
```
mutation DeleteQuery($id: ID!) { deleteJob(id:$id){  deleteJobId} }
```
variable
```
{ "id": "00da33e3-4ed1-4733-9fa4-88d516ab72e9" }
```