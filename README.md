# SWA-AWS-SAML-login
Golang app to login into AWS

## What is this for

CLI application to log into any aws accounts via AWS SAML, instead of running awssaml commands and remembering each flag to input.

## requirements

[aws cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)
[awssaml](https://docs.awssaml.ec.dev.aws.swacorp.com/index.html) installed and setup.
valid AWS credentials
golang

## How to use this

clone to your local machine.
adjust any accounts/name/ids in the script to your corresponding setup.

to check your own config run `awssaml list`

after you updated `awsLogin.go` according to your credentials run `go mod init awslogin && go tidy`

from that directory you can run `go run awsLogin.go login --name {name_of_role}` to test.

If working properly you can build the binary with `go build` and that will generate an executable you can run.

Next copy over the binary to your local bin to run from directory `cp awslogin /usr/local/bin`

Now you can run the aws login script from any directory ex `awslogin login --name devDev`