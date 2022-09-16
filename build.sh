#!/bin/bash

go build on/main.go
mv main notif-on

go build off/main.go
mv main notif-off