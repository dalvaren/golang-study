## Steps to compile the test package

1. source .project_bash-profile
2. go install testPkg
3. testPkg

## for myLibrary
1. go build myLibrary
2. go install testPkg
3. testPkg

## to generate the godoc
1. godoc -http=:6060
2. go to Packages/myLibrary

## testing the myLibrary package
1. go test myLibrary