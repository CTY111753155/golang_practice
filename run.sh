#!/bin/bash

go build -o booking cmd/web/*.go && ./booking
./booking -dbname=bookings -dbuser=yangzhicen -cache=false -production=false