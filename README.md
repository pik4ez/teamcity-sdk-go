TeamCity API bindings
=====================

This is a simple wrapper around the TeamCity API.

[![GoDoc](https://godoc.org/github.com/Cardfree/teamcity-sdk-go?status.png)](https://godoc.org/github.com/Cardfree/teamcity-sdk-go)

Sample usage:

```go
package main

import "github.com/Cardfree/teamcity-go-sdk/teamcity"

func main() {
	client := teamcity.New("myinstance.example.com", "username", "password")

	b, err := client.QueueBuild("Project_build_task", "master", nil)
	if err != nil {
		fmt.Printf("You're outta luck: %s\n", err)
		return
	}

	fmt.Printf("Build: %#v\n", b)
}
```
## Teamcity Rest API Docs
- [teamcity-rest-api](https://dploeger.github.io/teamcity-rest-api/)
- [perl5-teamcity-api](http://eilara.github.io/perl5-teamcity-api/)

## Starting the Docker Container for Testing
```bash
docker-compose up teamcity10
```

## Upgrading Teamcity

### Test Data
When Upgrading from one version of Teamcity to Another the Test Data needs to be upgraded as well.

1. Update the docker-compose.yml and Dockerfile's to the new version of teamcity
1. 
```bash
docker exec -it ${CONTAINER_ID} bash
cp -r /data/teamcity_server/datadir/config /test-data
cp -r /data/teamcity_server/datadir/system /test-data
```


## Debugging Rest Calls on Teamcity

```bash
docker exec -it ${CONTAINER_ID} bash
tail -f /opt/teamcity/logs/teamcity-rest.log
```
