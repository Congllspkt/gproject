package main

import (
	initdb "gproject/Initdb"
	echobasic "gproject/echoBasic"
)

func main() {
	initdb.InitDB()
	echobasic.InitEcho()
}

/*
git add .
git commit -m 'init first 2'
git push


*/