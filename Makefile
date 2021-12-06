APP = gosolid
APP_EXE = "./build/$(APP)"

build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=linux go build -o ${APP_EXE}

test:
	export CGO_ENABLED=0 && go test -cover -v ./...