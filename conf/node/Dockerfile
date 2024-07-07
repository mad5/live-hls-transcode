FROM golang:1.17 AS downloader
WORKDIR /usr/app                      

#FROM node:14

# same as before
WORKDIR /usr/app

RUN apt-get update
RUN apt-get -y install curl gnupg
RUN curl -sL https://deb.nodesource.com/setup_20.x  | bash -
RUN apt-get -y install nodejs

RUN apt-get update && apt-get install -y ffmpeg
#RUN apt-get update && apt-get install -y  npm
#RUN apt-get update && apt-get install -y golang

COPY . .
RUN go install github.com/MaZderMind/live-hls-transcode@master
#RUN cd $GOPATH/src/github.com/MaZderMind/live-hls-transcode/

#RUN make dependencies
#RUN make