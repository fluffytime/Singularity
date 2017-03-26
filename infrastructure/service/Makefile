#VID - Virtual Import Directory for project
VID=github.com/karolszmaj/gotrack/infrastructure/service
#Variables
all: protos

protos: shared/shared.pb.go auth/auth.pb.go

auth/auth.pb.go: auth/auth.proto
	protoc --go_out=plugins=grpc,Mshared/shared.proto=$(VID)/shared:. $<
	go install ./auth/

shared/shared.pb.go: shared/shared.proto
	protoc --go_out=. shared/shared.proto
	go install ./shared/
