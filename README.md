# fasten


Fasten makes it easy to setup and deploy to an AWS Lightsail instance.  

### How it works
Fasten manages everything over SCP and SSH, this means that no additional software except for Fasten is required.  Fasten does the following

 - Starts AWS Lightsail instance
 - Provisions Ubuntu and Amazon Linux servers with your chosen Language
 - Deploys your code
 - Manages your remote application processes
 - Installs your application dependencies


If you have a server that is running Ubuntu or Amazon Linux, Fasten will also manage deployments.

## Commands
```sh
Setup
  #This will set everything for you, launch instance, create Fasten config file and download instance key pair.  

Install
  #Provisions your instance with languages needed to run your application code.  

Deploy
  #Deploys your code to instance.

Status
  #Shows useful information about your Fasten instance and connection information.  

Generate
  #Creates Fasten config file, this is for when you are already running and AWS instance and have your key pair
```

## Description

## Usage

## Install

To install, use `go get`:

```bash
$ go get -d github.com/DaveBlooman/fasten
```

## Contribution

1. Fork ([https://github.com/DaveBlooman/fasten/fork](https://github.com/DaveBlooman/fasten/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[DaveBlooman](https://github.com/DaveBlooman)
