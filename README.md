# gproject

go mod init test
go run server\main.go


git add .
git commit -m 'go Logs Handler'
git push






go env -w GO111MODULE="on" 

go env -w GOPATH=C:\Users\Welcome\go



docker 
docker run --name my-postgres -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mypassword -e POSTGRES_DB=mydb -e POSTGRES_SSLMODE=disable -p 5432:5432 -d postgres



sql-migrate up
sql-migrate down



go get -u go.uber.org/zap

go get github.com/spf13/viper
vscode icon, yaml

middlewares

curl --header "Authorization: valid_token" http://localhost:8082/v1/2024/user/

docker run -d -p 33306:3306 --name mysql-container1 -e MYSQL_ROOT_PASSWORD=root1234 -e MYSQL_DATABASE=shopdevgo mysql

docker exec -it mysql-container1 bash
mysql -uroot -proot1234
use shopdevgo
show tables;
desc go_db_user;


BenchmarkMaxOpenConns1-12
     159           6803135 ns/op            6642 B/op         87 allocs/op
PASS
ok      gproject/internal/tests/benchmark       2.377s


BenchmarkMaxOpenConns10
BenchmarkMaxOpenConns10-12
    1159            973816 ns/op            5891 B/op         79 allocs/op

BenchmarkMaxOpenConns100
BenchmarkMaxOpenConns100-12
     781           1308933 ns/op            6673 B/op         82 allocs/op

159 times
6803135 ns
6642 Byte
87 allocs




docker run -d -p 16379:6379 --name redis-container1 redis


producer: push/write, tạo message
offset: vị trí tin nhắn trong queue
LEO: log End Offset
consumer: pull/nhận tin nhắn
consumer group: 1 consumer -> n partitions nhưng 1 partition -> 1 consumer
vi dụ:
+ p: p1, p2, p3, p4
+ cgA: c1, c2, c3, c4
+ cgB: c5, c6

-> Phân chia:
- p1:
     + cgA: c1
     + cgB: c5
- p2:
     + cgA: c2
     + cgB: c6
- p3:
     + cgA: c3
     + cgB: c5
- p4:
     + cgA: c4
     + cgB: c6

topic: là nhiều queue, tăng hiệu suất
partition: 1 topic có 3 partition, các partition chạy song song lẫn nhau
broker: 1 kafka là 1 cluster broker gồm nhiều brokers

docker run -it -p 8080:8080 -e DYNAMIC_CONFIG_ENABLED=true provectuslabs/kafka-ui



go get -u github.com/segmentio/kafka-go
awesome-go.com

docker-compose down
docker-compose up -d
go run cmd/server/main.go


cloudinary to save and resize image 
upload many files -> sent many requests to service to save file -> improve performence

s3 -> upload 

read: s3-request presigner

web tool keycdn



git add .
git commit -m 'test insert 1M, 10M'
git push

