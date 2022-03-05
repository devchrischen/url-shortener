#!/bin/bash

# workaround: wait for db to accept connection
sleep 15s

# Run DB migration
goose -dir ./goose mysql "chrischen:funnow@tcp(db:3306)/Url_Shortener?charset=utf8mb4&parseTime=True" up

# Run app
air