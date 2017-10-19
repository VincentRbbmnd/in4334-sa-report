# Build instructions

- Make sure you installed golang and project is cloned in your `go workspace` in src/github.com/VincentRbbmnd/in4334-sa-report.
- Install dep (dependency manager for golang) `go get -u github.com/golang/dep/cmd/dep` (only once)
- Run `dep ensure` in the github-crawler-api folder
- Run `go build -o 'OUTPUT_FILE_NAME'`
- You now have a binary that can be run to host the API

# Running the API

- Make sure you have the binary file for execution (Followed the build instructions)
- The binary will connect to the raspberry pi containing the GitHub data
- Run the binary with flags `./binary -user DB_USER_NAME -pass DB_USER_PASS -port 8082` to run with the PI

- Other flags:
```json
{
    "host": "IP for the DB",
    "user": "Username for the DB",
    "pass": "Password for the DB",
    "port": "Port for DB connection",
    "db": "Name for the database" 
}
```


# Generation instructions

### If you want to generate new routes, mediatypes, models for the backend and/or payloads

- Put the github-crawler-api (Already done for building probably)
- Run `go get -u github.com/goadesign/goa/...` 
- Adjust the files in the design folder for your changes
- Make sure you are in the ROOT folder of the project, `$GOPATH/github-crawler/api`
- Run `goagen bootstrap -d github-crawler-api/design` for the route/mediatype/payload generation
- Run `goagen --design github-crawler-api/design gen --pkg-path=github.com/goadesign/gorma` for DB model generation
- Run `go build -o 'OUTPUT_FILE_NAME'` afterwards to compile the changes
