go get -u github.com/lqs/sqlingo/sqlingo-gen-mysql
mkdir -p generated/sqlingo 
sqlingo-gen-mysql root:@/test >generated/sqlingo/test.dsl.go

go mod init sqlingo_study

https://www.v2ex.com/t/567482



