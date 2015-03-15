# Dockerfile for a Rails application using Unicorn
FROM packethost/baseimage
MAINTAINER Aaron Welch "welch@packet.net"

USER root
# Install goslack deps
RUN \ 
  apt-get update -q && \
	apt-get install -qy \
		golang

USER root
