# goとalpine
FROM mysql:latest

# update
RUN apt update -y && apt upgrade -y

# TimeZOneをJSTに変更
RUN apt-get install -y tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apt install locales-all -y

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

EXPOSE 3306