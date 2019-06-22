#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/play \
&& sftp pi@192.168.1.122:/home/pi/Golab <<< $'put rpi-test'