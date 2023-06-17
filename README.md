# Readme

Developed with : 
- go 1.19 & 1.20

Not tested on lower go versions.

## Guide to run functional test
### Prerequisite
1. Make sure you run installed go version >= `1.19`
2. Run your program to serve port `8080`.

### How to run project
1. `cd parking-lot`
2. `make run`

### Run Functional Test & Unit tests
1. `cd parking-lot`
2. open two terminal in the root project ("parking-lot")
3. if the project is already running skip this else in terminal 1 run `make run`
4. in terminal 2 run `make test-all`
