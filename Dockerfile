# Dockerfile for running goslack
FROM quay.io/packet/baseimage
MAINTAINER Aaron Welch "welch@packet.net"

ADD goslack /usr/local/bin/

USER root
# Install goslack deps
RUN \ 
  apt-get update -q && \
	apt-get install -qy \
		golang

CMD goslack -slackpath="$SLACK_PATH" -text="$SLACK_TEXT" -channel=$SLACK_CHANNEL -username="$SLACK_USERNAME" -emoji=$SLACK_EMOJI -letoken="$LETOKEN"
