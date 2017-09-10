#!/bin/bash

set -eux

apt-get update -y
curl -sL https://deb.nodesource.com/setup_7.x | sudo -E bash -
apt-get install -y nodejs build-essential
apt-get update -y
