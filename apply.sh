#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/play \
&& sftp pi@192.168.1.132:/home/pi/Golab <<< $'put bin/play'