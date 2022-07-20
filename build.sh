#!/bin/bash

echo '==========================ビルド開始=========================='

select VAR in notification scraping-girl-info scraping-new-face scraping-rank scraping-schedule-update scraping-schedule-identify reply
do 
	echo "$VARのビルドを開始します。"
	break
done

cd ./src/"$VAR"/

echo `pwd`

git pull
GOOS=linux go build main.go
GOOS=linux CGO_ENABLED=0 go build main.go
zip function.zip main
echo '==========================ビルド終了=========================='
