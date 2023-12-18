# This file is used to build and run docker container

build:
	docker build -t wthunder/pork-go .

run:
	docker run -ti --net=host -p 8000:8000 \
		-e SMTP_PASS=${SMTP_PASS} \
		-e DB_CONN=${DB_CONN} \
		-e JWT_SECRET=${JWT_SECRET} \
		-e KNOSTASH_BACKEND_URL=${KNOSTASH_BACKEND_URL} \
		wthunder/pork-go

