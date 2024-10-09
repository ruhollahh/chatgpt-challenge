## Quick Start
### Where can I edit input prompts?
look inside the ```startLaptopWorkerQueue()``` function in the main.go file.

you can also edit all the other configurations inside the ```main()``` function in the main.go file.

### How can I run the app?
replace ```YOUR_AUTH_TOKEN``` in the following command and run it your terminal:
```bash
CHATGPT_CHALLENGE__AUTH_TOKEN=YOUR_AUTH_TOKEN make run
```
if you don't have ```make``` installed you can just run:
```bash
CHATGPT_CHALLENGE__AUTH_TOKEN=YOUR_AUTH_TOKEN go run main.go
```
#### In order To interact with the APIs visit this URL:
http://localhost:1414/swagger/index.html

### How can I run the unit tests?
```bash
make test
```
if you don't have ```make``` installed you can just run:
```bash
go test -v ./...
```
