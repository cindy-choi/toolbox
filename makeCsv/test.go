package main
import (
  "fmt"
  "encoding/base64"
  "encoding/json"
  "net/http"
  "bytes"
)

type JiraRequest struct {
  Jql string `json:"jql"`
  Fields []string `json:"fields"`
}

type JiraResponse struct {
  StartAt int `json:"startAt"`
  MaxResults int `json:"maxResults"`
  Total int `json:"total"`
  Issues []Issue `json:"issues"`
}

type Issue struct {
  Id string `json:"id"`
  Key string `json:"key"`
  Points int
  Fields struct {
    FixVersion string `json:"fixVersion"`
  }
}

type UpdateField struct {
  Fields struct {
    FixVersion string
  }
}

func makeRequest() ([]byte, string){
  auth := base64.StdEncoding.EncodeToString([]byte("cindy.choi:chltpsk1"))
  request := JiraRequest {
    `project = NEA AND resolution = Unresolved AND fixVersion = EMPTY AND assignee != rosa.kim`,
    []string{"id", "key"},
  }

  query, err := json.Marshal(request)
  if err != nil {
    panic(err)
  }
  return query, auth
}

func listUp(jira *JiraResponse) {
  for _, issue := range jira.Issues {
    fmt.Printf("[%s] fixVersion = %s\n", issue.Key, issue.Fields.FixVersion)
  }
}

func updateVersion(jira *JiraResponse) {
  for _, issue := range jira.Issues {
    inputs := UpdateField {
      Fields {
        "FixVersion": "1.0"
      }
    }
    json, err := json.Marshal(inputs)
    if err != nil {
        panic(err)
    }

    request, err := http.NewRequest("PUT", `http://jira.nexrcorp.com/rest/api/2/issue/` + issue.Id, bytes.NewBuffer(json))
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Authorization", "Basic " + auth)
    if err != nil {
      panic(err)
    }

  }
}

func main() {
  query, auth := makeRequest()
  request, err := http.NewRequest("POST", "http://jira.nexrcorp.com/rest/api/2/search", bytes.NewBuffer(query))
  request.Header.Add("Content-Type", "application/json")
  request.Header.Add("Authorization", "Basic " + auth)
  if err != nil {
    panic(err)
  }

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    panic(err)
  }
  defer response.Body.Close()

  jiraResponse := new(JiraResponse)
  json.NewDecoder(response.Body).Decode(&jiraResponse)

  listUp(jiraResponse)
  updateList(jiraResponse)
}
