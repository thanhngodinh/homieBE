# go-sql-rest-api

## How to run

#### To run the application

```shell
go run main.go
```

## API Design

### Common HTTP methods

- GET: retrieve a representation of the resource
- POST: create a new resource
- PUT: update the resource
- DELETE: delete a resource

## API design for health check

To check if the service is available.

#### _Request:_ GET /health

#### _Response:_

```json
{
  "status": "UP",
  "details": {
    "sql": {
      "status": "UP"
    }
  }
}
```

# Swagger Example API

This is a sample server celler server.

## Version: 1.0

### Terms of service

<http://swagger.io/terms/>

**Contact information:**  
API Support  
<http://www.swagger.io/support>  
support@swagger.io

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### /auth/login

#### POST

##### Summary

Login

##### Parameters

| Name | Located in | Description   | Required | Schema                     |
| ---- | ---------- | ------------- | -------- | -------------------------- |
| User | body       | User to login | Yes      | [domain.User](#domainuser) |

##### Responses

| Code | Description                                          | Schema                                  |
| ---- | ---------------------------------------------------- | --------------------------------------- |
| 201  | Created                                              | [util.Response](#utilresponse) & object |
| 400  | Invalid character 's' looking for beginning of value | string                                  |
| 500  | Internal Server Error                                | string                                  |

### /auth/register

#### POST

##### Summary

Create a User

##### Parameters

| Name | Located in | Description    | Required | Schema                     |
| ---- | ---------- | -------------- | -------- | -------------------------- |
| User | body       | User to create | Yes      | [domain.User](#domainuser) |

##### Responses

| Code | Description                                          | Schema                                  |
| ---- | ---------------------------------------------------- | --------------------------------------- |
| 201  | Created                                              | [util.Response](#utilresponse) & object |
| 400  | Invalid character 's' looking for beginning of value | string                                  |
| 500  | Internal Server Error                                | string                                  |

### /hostels

#### GET

##### Summary

Get Hostels

##### Description

Return a list of the Hostels included the pagination

##### Parameters

| Name          | Located in | Description                               | Required | Schema |
| ------------- | ---------- | ----------------------------------------- | -------- | ------ |
| Authorization | header     | The Authorization                         | Yes      | string |
| pageIdx       | query      | The index of the page start from 0        | Yes      | string |
| pageSize      | query      | The number of Hostels return on each page | Yes      | string |

##### Responses

| Code | Description           | Schema                                                 |
| ---- | --------------------- | ------------------------------------------------------ |
| 200  | OK                    | [domain.GetHostelsResponse](#domaingethostelsresponse) |
| 500  | Internal Server Error | string                                                 |

#### POST

##### Summary

Create a Hostel

##### Parameters

| Name          | Located in | Description       | Required | Schema                         |
| ------------- | ---------- | ----------------- | -------- | ------------------------------ |
| Authorization | header     | The Authorization | Yes      | string                         |
| Hostel        | body       | Hostel to create  | Yes      | [domain.Hostel](#domainhostel) |

##### Responses

| Code | Description                                          | Schema                                  |
| ---- | ---------------------------------------------------- | --------------------------------------- |
| 201  | Created                                              | [util.Response](#utilresponse) & object |
| 400  | Invalid character 's' looking for beginning of value | string                                  |
| 500  | Internal Server Error                                | string                                  |

### /hostels/{code}

#### DELETE

##### Summary

Delete a Hostel

##### Parameters

| Name          | Located in | Description            | Required | Schema |
| ------------- | ---------- | ---------------------- | -------- | ------ |
| Authorization | header     | The Authorization      | Yes      | string |
| code          | path       | The code of The Hostel | Yes      | string |

##### Responses

| Code | Description           | Schema |
| ---- | --------------------- | ------ |
| 200  | 1                     | string |
| 404  | 0                     | string |
| 500  | Internal Server Error | string |

#### GET

##### Summary

Get a Hostel

##### Description

Return a Hostel with the code

##### Parameters

| Name          | Located in | Description            | Required | Schema |
| ------------- | ---------- | ---------------------- | -------- | ------ |
| Authorization | header     | The Authorization      | Yes      | string |
| code          | path       | The code of the Hostel | Yes      | string |

##### Responses

| Code | Description           | Schema                         |
| ---- | --------------------- | ------------------------------ |
| 200  | OK                    | [domain.Hostel](#domainhostel) |
| 500  | Internal Server Error | string                         |

#### PUT

##### Summary

Update a Hostel

##### Parameters

| Name          | Located in | Description            | Required | Schema                         |
| ------------- | ---------- | ---------------------- | -------- | ------------------------------ |
| Authorization | header     | The Authorization      | Yes      | string                         |
| code          | path       | The code of the Hostel | Yes      | string                         |
| Hostel        | body       | Hostel to update       | Yes      | [domain.Hostel](#domainhostel) |

##### Responses

| Code | Description                                          | Schema                                  |
| ---- | ---------------------------------------------------- | --------------------------------------- |
| 200  | OK                                                   | [util.Response](#utilresponse) & object |
| 400  | Invalid character 's' looking for beginning of value | string                                  |
| 500  | Internal Server Error                                | string                                  |

### /team_members

#### GET

##### Summary

Get TeamMembers

##### Description

Return a list of the TeamMembers included the pagination

##### Parameters

| Name          | Located in | Description                                   | Required | Schema |
| ------------- | ---------- | --------------------------------------------- | -------- | ------ |
| Authorization | header     | The Authorization                             | Yes      | string |
| pageIdx       | query      | The index of the page start from 0            | Yes      | string |
| pageSize      | query      | The number of TeamMembers return on each page | Yes      | string |

##### Responses

| Code | Description           | Schema                                                         |
| ---- | --------------------- | -------------------------------------------------------------- |
| 200  | OK                    | [domain.GetTeamMembersResponse](#domaingetteammembersresponse) |
| 500  | Internal Server Error | string                                                         |

#### POST

##### Summary

Create a TeamMember

##### Parameters

| Name          | Located in | Description          | Required | Schema                                 |
| ------------- | ---------- | -------------------- | -------- | -------------------------------------- |
| Authorization | header     | The Authorization    | Yes      | string                                 |
| TeamMember    | body       | TeamMember to create | Yes      | [domain.TeamMember](#domainteammember) |

##### Responses

| Code | Description                                          | Schema                                  |
| ---- | ---------------------------------------------------- | --------------------------------------- |
| 201  | Created                                              | [util.Response](#utilresponse) & object |
| 400  | Invalid character 's' looking for beginning of value | string                                  |
| 500  | Internal Server Error                                | string                                  |

### /team_members/{code}

#### DELETE

##### Summary

Delete a TeamMember

##### Parameters

| Name          | Located in | Description                | Required | Schema |
| ------------- | ---------- | -------------------------- | -------- | ------ |
| Authorization | header     | The Authorization          | Yes      | string |
| code          | path       | The code of The TeamMember | Yes      | string |

##### Responses

| Code | Description           | Schema |
| ---- | --------------------- | ------ |
| 200  | 1                     | string |
| 404  | 0                     | string |
| 500  | Internal Server Error | string |

### /team_members/{id}

#### GET

##### Summary

Get a TeamMember

##### Description

Return a TeamMember with the id

##### Parameters

| Name          | Located in | Description              | Required | Schema |
| ------------- | ---------- | ------------------------ | -------- | ------ |
| Authorization | header     | The Authorization        | Yes      | string |
| id            | path       | The id of the TeamMember | Yes      | string |

##### Responses

| Code | Description           | Schema                                 |
| ---- | --------------------- | -------------------------------------- |
| 200  | OK                    | [domain.TeamMember](#domainteammember) |
| 500  | Internal Server Error | string                                 |

#### PUT

##### Summary

Update a TeamMember

##### Parameters

| Name          | Located in | Description              | Required | Schema                                 |
| ------------- | ---------- | ------------------------ | -------- | -------------------------------------- |
| Authorization | header     | The Authorization        | Yes      | string                                 |
| id            | path       | The id of the TeamMember | Yes      | string                                 |
| TeamMember    | body       | TeamMember to update     | Yes      | [domain.TeamMember](#domainteammember) |

##### Responses

| Code | Description                                          | Schema                                  |
| ---- | ---------------------------------------------------- | --------------------------------------- |
| 200  | OK                                                   | [util.Response](#utilresponse) & object |
| 400  | Invalid character 's' looking for beginning of value | string                                  |
| 500  | Internal Server Error                                | string                                  |

### /team_members/hostels/{code}

#### GET

##### Summary

Get TeamMembers by Hostel's code

##### Description

Return a list of the TeamMember belong to a Hostel included the pagination

##### Parameters

| Name          | Located in | Description                                   | Required | Schema |
| ------------- | ---------- | --------------------------------------------- | -------- | ------ |
| Authorization | header     | The Authorization                             | Yes      | string |
| code          | path       | The code of the Hostel                        | Yes      | string |
| pageIdx       | query      | The index of the page start from 0            | Yes      | string |
| pageSize      | query      | The number of TeamMembers return on each page | Yes      | string |

##### Responses

| Code | Description           | Schema                                                         |
| ---- | --------------------- | -------------------------------------------------------------- |
| 200  | OK                    | [domain.GetTeamMembersResponse](#domaingetteammembersresponse) |
| 500  | Internal Server Error | string                                                         |

### Models

#### domain.GetHostelsResponse

| Name       | Type                               | Description | Required |
| ---------- | ---------------------------------- | ----------- | -------- |
| data       | [ [domain.Hostel](#domainhostel) ] |             | No       |
| pagination | [util.Pagination](#utilpagination) |             | No       |

#### domain.GetTeamMembersResponse

| Name       | Type                                       | Description | Required |
| ---------- | ------------------------------------------ | ----------- | -------- |
| data       | [ [domain.TeamMember](#domainteammember) ] |             | No       |
| pagination | [util.Pagination](#utilpagination)         |             | No       |

#### domain.Hostel

| Name     | Type   | Description                                         | Required |
| -------- | ------ | --------------------------------------------------- | -------- |
| district | string | _Example:_ `"Bedford"`                              | No       |
| email    | string | _Example:_ `"peter.p@zylker.com"`                   | No       |
| name     | string | _Example:_ `"Robert Robertson"`                     | No       |
| ownerId  | string | _Example:_ `"07e7a76c-1bbb-11ed-861d-0242ac120002"` | No       |
| phone    | string | _Example:_ `"09832209761"`                          | No       |
| province | string | _Example:_ `"Titao"`                                | No       |
| status   | string | _Example:_ `"Active"`                               | No       |
| street   | string | _Example:_ `"144 J B Hazra Road"`                   | No       |

#### domain.TeamMember

| Name        | Type   | Description                                         | Required |
| ----------- | ------ | --------------------------------------------------- | -------- |
| birthDate   | string | _Example:_ `"1991-05-06"`                           | No       |
| district    | string | _Example:_ `"Bedford"`                              | No       |
| email       | string | _Example:_ `"peter.p@zylker.com"`                   | No       |
| firstName   | string | _Example:_ `"Robert"`                               | No       |
| id          | string | _Example:_ `"6a077d3c-1bbb-11ed-861d-0242ac120002"` | No       |
| lastName    | string | _Example:_ `"Robertson"`                            | No       |
| hostelCode  | string | _Example:_ `"07e7a76c-1bbb-11ed-861d-0242ac120002"` | No       |
| nationality | string | _Example:_ `"Afghanistan"`                          | No       |
| password    | string | _Example:_ `"Password@3"`                           | No       |
| phone       | string |                                                     | No       |
| province    | string | _Example:_ `"Titao"`                                | No       |
| role        | string | _Example:_ `"Owner or Manager or Staff"`            | No       |
| street      | string | _Example:_ `"144 J B Hazra Road"`                   | No       |
| username    | string | _Example:_ `"robertson"`                            | No       |

#### domain.User

| Name     | Type   | Description | Required |
| -------- | ------ | ----------- | -------- |
| password | string |             | No       |
| username | string |             | No       |

#### util.Pagination

| Name     | Type    | Description | Required |
| -------- | ------- | ----------- | -------- |
| pageIdx  | integer |             | No       |
| pageSize | integer |             | No       |
| total    | integer |             | No       |

#### util.Response

| Name   | Type    | Description | Required |
| ------ | ------- | ----------- | -------- |
| data   |         |             | No       |
| error  |         |             | No       |
| status | integer |             | No       |

## Common libraries

- [core-go/health](https://github.com/core-go/health): include HealthHandler, HealthChecker, SqlHealthChecker
- [core-go/config](https://github.com/core-go/config): to load the config file, and merge with other environments (SIT, UAT, ENV)
- [core-go/log](https://github.com/core-go/log): log and log middleware

### core-go/health

To check if the service is available, refer to [core-go/health](https://github.com/core-go/health)

#### _Request:_ GET /health

#### _Response:_

```json
{
  "status": "UP",
  "details": {
    "sql": {
      "status": "UP"
    }
  }
}
```

To create health checker, and health handler

```go
    db, err := sql.Open(conf.Driver, conf.DataSourceName)
    if err != nil {
        return nil, err
    }

    sqlChecker := s.NewSqlHealthChecker(db)
    healthHandler := health.NewHealthHandler(sqlChecker)
```

To handler routing

```go
    r := mux.NewRouter()
    r.HandleFunc("/health", healthHandler.Check).Methods("GET")
```

### core-go/config

To load the config from "config.yml", in "configs" folder

```go
package main

import "github.com/core-go/config"

type Root struct {
    DB DatabaseConfig `mapstructure:"db"`
}

type DatabaseConfig struct {
    Driver         string `mapstructure:"driver"`
    DataSourceName string `mapstructure:"data_source_name"`
}

func main() {
    var conf Root
    err := config.Load(&conf, "configs/config")
    if err != nil {
        panic(err)
    }
}
```

### core-go/log _&_ core-go/middleware

```go
import (
    "github.com/core-go/config"
    "github.com/core-go/log"
    m "github.com/core-go/middleware"
    "github.com/gorilla/mux"
)

func main() {
    var conf app.Root
    config.Load(&conf, "configs/config")

    r := mux.NewRouter()

    log.Initialize(conf.Log)
    r.Use(m.BuildContext)
    logger := m.NewLogger()
    r.Use(m.Logger(conf.MiddleWare, log.InfoFields, logger))
    r.Use(m.Recover(log.ErrorMsg))
}
```

To configure to ignore the health check, use "skips":

```yaml
middleware:
  skips: /health
```
