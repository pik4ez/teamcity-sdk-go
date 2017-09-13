TeamCity API bindings
=====================

This is a simple wrapper around the TeamCity API.

[![GoDoc](https://godoc.org/github.com/Cardfree/teamcity-sdk-go?status.png)](https://godoc.org/github.com/Cardfree/teamcity-sdk-go)

Sample usage:

```
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
# Teamcity Rest API Docs
https://dploeger.github.io/teamcity-rest-api/#addTrigger
http://eilara.github.io/perl5-teamcity-api/


# Upgrading Teamcity Test Data

1. Update the docker-compose.yml and Dockerfile's to the new version
1. `docker exec -it ${CONTAINER_ID} bash`
1. `cp -r /data/teamcity_server/datadir/config /test-data/config`
1. `cp -r /data/teamcity_server/datadir/system /test-data/system`



# Debugging Rest Calls on Teamcity
```bash
tail -f /opt/teamcity/logs/teamcity-rest.log
```
