FROM golang:alpine AS builder

#gym_controll_db
ARG gym_controll_db_drive
ARG gym_controll_db_host
ARG gym_controll_db_port
ARG gym_controll_db_user
ARG gym_controll_db_pass
ARG gym_controll_db_name

ENV GymControll_DBDRIVER $gym_controll_db_drive
ENV GymControll_DBHOST $gym_controll_db_host
ENV GymControll_DBPORT $gym_controll_db_port
ENV GymControll_DBNAME $gym_controll_db_name
ENV GymControll_DBUSER $gym_controll_db_user
ENV GymControll_DBPASS $gym_controll_db_pass

#gym_controll_webroute
ARG gym_controll_web_server_port

ENV GymControll_WEBSERVER_PORT $gym_controll_web_server_port

#gym_controll_userenvs
ARG gym_controll_passsecret

ENV GymControll_PassSecret $gym_controll_passsecret

WORKDIR /app
COPY . .
WORKDIR /app/cmd/gym_menager/
RUN CGO_ENABLED=0 GOOS=linux go build -v -x -ldflags="-w -s" -o /app/server .

CMD ["/app/server"]