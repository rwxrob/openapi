FROM ubuntu

LABEL MAINTAINER "Rob Muhlestein <rob@rwx.gg>"
LABEL SOURCE "https://github.com/rwxrob/openapi"

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get -y --no-install-recommends upgrade
RUN apt-get install -y ca-certificates apt-utils build-essential
RUN update-ca-certificates
RUN apt-get install -y --no-install-recommends openjdk-18-jdk-headless \
    golang curl vim tmux perl python-is-python3 git gh jq sudo shellcheck \
    ssh rsync cifs-utils smbclient bash-completion less

COPY ./files/. ./Dockerfile /
COPY build /usr/share/workspace

# prompts for user, group, ids, and shell
ENTRYPOINT ["sh","/entry"] 
