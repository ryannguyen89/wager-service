FROM debian:stable-slim
WORKDIR /app
RUN mkdir -p /usr/share/
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN apt-get install -y tzdata
ENV TZ=Asia/Ho_Chi_Minh
COPY swagger /app/swagger
COPY wager-service /app/
COPY entry-point.sh /app/
ENTRYPOINT ["./entry-point.sh"]