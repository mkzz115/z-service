OUTPUT=./output
API_NAME=api_service
THRIFT_NAME=thrift_server

all: install

api:
	@echo $(API_NAME)
	mkdir -p ${OUTPUT}/bin/$(API_NAME)
	go build -o ${OUTPUT}/bin/$(API_NAME)/$(API_NAME) api_service/main.go
	chmod u+x ${OUTPUT}/bin/$(API_NAME)
srv:
	mkdir -p ${OUTPUT}/bin/${THRIFT_NAME}
	go build -o ${OUTPUT}/bin/${THRIFT_NAME}/${THRIFT_NAME} api_service/main.go
	chmod u+x ${OUTPUT}/bin/$(THRIFT_NAME)
	
pre: clean
	mkdir -p ${OUTPUT}/bin

fmt:
	go fmt -l -w -s ./

install: api srv
	cp -r config ${OUTPUT}/
	@echo "success"

clean:
	rm -rf ${OUTPUT}

.PHONY: clean fmt


