#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=5 go build \
&& sftp pi@192.168.1.122:/home/pi/Golab <<< $'put rpi-test'