# https://hub.docker.com/r/phusion/baseimage/
FROM phusion/baseimage:0.9.22

# Add mattes/migrate package (for DB migrations) to apt-get sources
RUN curl -L https://packagecloud.io/mattes/migrate/gpgkey | apt-key add - && \
  echo "deb https://packagecloud.io/mattes/migrate/ubuntu/ xenial main" > /etc/apt/sources.list.d/migrate.list

# Install apt-get packages
RUN apt-get update && apt-get -y install wget vim curl git ntp migrate

# Install Go 1.9.0
RUN mkdir -p /tmp && cd /tmp && \
  wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go1.9.linux-amd64.tar.gz
ENV PATH "/usr/local/go/bin:/go/bin:${PATH}"
ENV GOPATH "/go"

# Install Go Dep (vendoring)
RUN curl -L -o /usr/local/go/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
  chmod +x /usr/local/go/bin/dep

# Cleanup
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Timezone https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List
ENV TZ Asia/Singapore

# Setup cron
ADD ./crontab /etc/cron.d/
RUN chmod 0644 /etc/cron.d/crontab

# Logging dir
RUN mkdir -p /var/log/app

# Add all source code
ADD ./app /go/src/domain/user/app

# Build all binaries in /go/bin/*
WORKDIR /go/src/domain/user/app/cmd/server
RUN go get && go install

# App daemon, assuming this is the main server process.
RUN mkdir -p /etc/service/server
ADD ./runit/server.sh /etc/service/server/run
RUN chmod u+x /etc/service/server/run

WORKDIR /go/src/domain/user/app/cmd/server

