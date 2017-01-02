#!/bin/bash

yum erase ruby20 ruby -y
yum install ruby23-devel ruby23 -y
gem update --system
echo "export PATH=$PATH:/home/ec2-user/bin" >> /home/ec2-user/.bashrc
