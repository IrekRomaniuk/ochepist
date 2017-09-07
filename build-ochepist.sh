#!/usr/bin/bash
env GOOS=linux GOARCH=386 go build -o bin/ochepist_lin
env GOOS=windows GOARCH=386 go build -o bin/ochepist_win
dos2unix.exe templates/*
dos2unix.exe results/*
tar -zcvf ochepist.tar.gz bin/ templates/ results/g-ochepist-test-dbedit.txt