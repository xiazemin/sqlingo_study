go get -u github.com/lqs/sqlingo/sqlingo-gen-mysql
mkdir -p generated/sqlingo 
sqlingo-gen-mysql root:@/test >generated/sqlingo/test.dsl.go

sqlingo-gen-mysql root:@tcp(127.0.0.1:3306)/test 

go mod init sqlingo_study

https://www.v2ex.com/t/567482



