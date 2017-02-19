#!/bin/bash

set -eux

apt-get update -y
apt-get install ruby2.3 ruby2.3-dev git-core curl zlib1g-dev build-essential libssl-dev libreadline-dev libyaml-dev libsqlite3-dev sqlite3 libxml2-dev libxslt1-dev libcurl4-openssl-dev libffi-dev -y

chown -R ubuntu:ubuntu /usr/local/bin /var/lib/gems/2.3.0
