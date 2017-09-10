# Fasten

Fasten makes it easy to setup and deploy to an Ubuntu or Amazon Linux instance easily.  

### How it works
Fasten manages everything over `scp` and `ssh`, this means that no additional software except for Fasten is required.  Fasten does the following

 - Provisions Ubuntu and Amazon Linux servers with your chosen language
 - Deploys your code
 - Manages your remote application processes
 - Installs your application dependencies

### Support Languages

 - Ruby  
 - Golang
 - Node.js

### Support Operating Systems

- Ubuntu 16.04
- Amazon Linux

### Installation

Download the binary for your operating system from the releases page or compile locally by running `make`

### Quick Start

1. Run the `fasten init` command.  This will walk you through setting up an application.  If you get stuck, there is an examples directory with a ruby and node.js application.  It is important to note that you are responsible for completing the config to your application needs, so continuing on from this point without editing the `fasten.yaml` will not result in a working application.  If you are running a Go application, the binary must be called main, see the examples directory.

1. Edit the `fasten.yaml` file to include your application start command and if needed, a pre start command.  This is important if your application needs static assets compiled or you need to run a IO operation before your main app starts.

1. Ensure to set the correct permissions on your key pair, `chmod 400 keypair.pem`

1. Once you have your `fasten.yaml` complete and you server is running, you can install the software needed for your applications by running `fasten install`

1. This will provision the server with all the dependencies you'll need.  Once complete, you can deploy your code using `fasten deploy`

Fasten will copy your code to the server, install dependencies and start the process.  The copy operation uses `scp`, so all files are securely copied to the server.  The dependencies are installed based on the package manager of choice, bundler or npm for example.  The process of the application is managed by a PID file in the directory of your application.  That pid is used by fasten to start and stop the process.  

### Ignore Files

An optional feature of fasten is the ability to ignore files from the deploy.  This is useful for when you are install local dependencies and don't want them copied to the server. Simply create a `.fastenignore` file in the same directory as the fasten.yaml file and the files or directories will be ignore from deployments.

## Commands
```sh
install
  #Provisions your instance with languages needed to run your application code.  

deploy
  #Deploys your code to server.

stop
  #Stops all applications.  

init
  #Creates Fasten config file, this is for when you are already running and AWS instance and have your key pair
```

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
