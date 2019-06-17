OUTPUT=./output
API_NAME=api_service
THRIFT_NAME=thrift_server

BUILD = go build
GOFLAGS = -mod=vendor

all: install

api:
	@echo $(API_NAME)
	$(BUILD) $(GOFLAGS) -o ${OUTPUT}/${API_NAME} api_service/main.go
	chmod u+x ${OUTPUT}/${API_NAME}
srv:
	$(BUILD) $(GOFLAGS) -o ${OUTPUT}/${THRIFT_NAME} server/main.go
	chmod u+x ${OUTPUT}/${THRIFT_NAME}
	
pre: clean
	mkdir -p ${OUTPUT}

fmt:
	go fmt -l -w -s ./

install: api srv
	cp -r config ${OUTPUT}/
	@echo "success"

clean:
	rm -rf ${OUTPUT}

.PHONY: clean fmt


