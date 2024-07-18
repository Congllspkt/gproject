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




git add .
git commit -m 'Migrating tables'
git push

