FROM golang:1.17 AS downloader

WORKDIR /usr/app

RUN apt-get update
RUN apt-get -y install curl gnupg
RUN curl -sL https://deb.nodesource.com/setup_20.x  | bash -
RUN apt-get update && apt-get -y install nodejs

RUN apt-get update && apt-get install -y ffmpeg

COPY . .
RUN go install github.com/mad5/live-hls-transcode@master

