FROM golang:alpine AS builder

ARG gym_controll_db_drive
ARG gym_controll_db_host
ARG gym_controll_db_port
ARG gym_controll_db_user
ARG gym_controll_db_pass
ARG gym_controll_db_name

ARG gym_controll_web_server_port

ENV GymControll_DBDRIVER $gym_controll_db_drive
ENV GymControll_DBHOST $gym_controll_db_host
ENV GymControll_DBPORT $gym_controll_db_port
ENV GymControll_DBNAME $gym_controll_db_port
ENV GymControll_DBUSER $gym_controll_db_name
ENV GymControll_DBPASS $gym_controll_db_pass

ENV GymControll_WEBSERVER_PORT $gym_controll_web_server_port



WORKDIR /app
COPY . .
WORKDIR /app/cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server .

CMD ["/app/server"]