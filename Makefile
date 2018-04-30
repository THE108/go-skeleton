ifeq ($(strip $(GO)),)
GO=$(shell which go)
endif

app_version='$(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)'
go_version='$(shell ${GO} version | sed -E 's|.*go(([0-9]+\.){1,2}[0-9]+) .*|\1|g')'
build_date='$(shell date -u "+%Y-%m-%dT%H:%M:%S")'
git_log='$(shell git log --decorate --oneline -n1 | awk '{print $$1}')'
ldflags="-X=main.GitLog=${git_log} -X=main.AppVersion=${app_version} -X=main.GoVersion=${go_version} -X=main.BuildDate=${build_date}"

build:
	${GO} build -ldflags=butler{ .Vars.repoPath }/butler{ .Project.Name }=${ldflags} -o ${GOPATH}/bin/butler{ .Project.Name }

release:
	GOOS=linux ${GO} install -ldflags=butler{ .Vars.repoPath }/butler{ .Project.Name }=${ldflags} -v ./...
