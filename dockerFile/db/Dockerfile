# goとalpine
FROM mysql:latest

# update
RUN apt update -y && apt upgrade -y

# TimeZOneをJSTに変更
RUN apt-get install -y tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

EXPOSE 3306