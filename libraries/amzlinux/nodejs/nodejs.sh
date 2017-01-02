#!/bin/bash

set -eux

curl --silent --location https://rpm.nodesource.com/setup_7.x | bash -
yum install -y gcc-c++ make -y
yum install -y nodejs -y
