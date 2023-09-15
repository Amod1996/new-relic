

## build docker images for all of. these can then be run via docker run
docker-build:
	docker build -f server/Dockerfile -t three_word_occurance_calculator .

## pulls the latest application and test dependencies and updates go.mod
update-deps:
	go get -t -u ./...
	go mod tidy
	go mod vendor
	go mod verify

## perform code formatting (and enhanced version of go fmt)
gofumpt:
	@echo "# Formatting code...."
	go install -mod=mod mvdan.cc/gofumpt@latest
	gofumpt -l -w .



start-go-program:
	docker run -i --name three_word_occurance_calculator three_word_occurance_calculator


#start-the-process: clean docker-build  start-go-program
start-the-process: clean docker-build start-go-program

clean: update-deps gofumpt

CONTAINERS_TO_REMOVE :=  three_word_occurance_calculator
IMAGES_TO_REMOVE :=  three_word_occurance_calculator

remove-containers:
	docker stop $(CONTAINERS_TO_REMOVE)
	docker rm -f $(CONTAINERS_TO_REMOVE)


remove-images:
	docker rmi -f $(IMAGES_TO_REMOVE)

stop-the-process: remove-containers remove-images









