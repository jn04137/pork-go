#!/bin/bash

docker build -t wthunder/pork-go .
docker run -p 8000:8000 \
  -e SMTP_PASS=${SMTP_PASS} \
  -e DB_CONN=${DB_CONN} \
  -e JWT_SECRET=${JWT_SECRET} \ 
  wthunder/pork-go
