ROOTDIR=/Users/james/Dropbox/Development/go/src/Gateway311
ENGINE_SRC=/Users/james/Dropbox/Development/go/src/Gateway311/engine
ADP_CS_SRC=/Users/james/Dropbox/Development/go/src/Gateway311/adapters/citysourced
ADP_EMAIL_SRC=/Users/james/Dropbox/Development/go/src/Gateway311/adapters/email
CS_SIM_SRC=/Users/james/Dropbox/Development/go/src/CitySourcedAPI
CS_SIM_TEST=/Users/james/Dropbox/Work/CodeForSanJose/Open311/_test/CitySourced

default: builddocker

linuxcompile:
	env GOOS=linux GOARCH=amd64 go build -o gateway Gateway311/engine
	env GOOS=linux GOARCH=amd64 go build -o adapters/citysourced/adp_cs Gateway311/adapters/citysourced
	env GOOS=linux GOARCH=amd64 go build -o adapters/email/adp_email Gateway311/adapters/email
	env GOOS=linux GOARCH=amd64 go build -o simulators/citysourced/cs_sim CitySourcedAPI

maccompile:
	go build -o engine/gateway Gateway311/engine

	go build -o adapters/citysourced/adp_cs Gateway311/adapters/citysourced

	go build -o adapters/email/adp_email Gateway311/adapters/email

	go build -o $(CS_SIM_SRC)/cs_sim CitySourcedAPI
	cp $(CS_SIM_SRC)/cs_sim $(CS_SIM_TEST)/cfg1
	cp $(CS_SIM_SRC)/cs_sim $(CS_SIM_TEST)/cfg2
	cp $(CS_SIM_SRC)/cs_sim $(CS_SIM_TEST)/cfg3

	go build -o monitor/monitor Gateway311/monitor


builddocker: linux


runtest: maccompile

