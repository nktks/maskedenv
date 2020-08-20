# maskedenv

Print masked env.

# why?

Sometimes in Kubernetes Pods, when `STDOUT`, `STDERR` attached at `/proc/1/fd/1`,
outputs printed to host log files, and exported to cloud log platform by log agent.
(ie Stackdriver logging agent in GKE, Cloudwatch logging agent in EKS)

If you want to check whether some environment is set,
you need carefully to use env command.
`maskedenv` print masked environment variables, you can check target env variables was set with masked print.

# How to install
## by go get
```
# go get github.com/nakatamixi/maskedenv
# maskedenv
HOSTNAME=************
HOME=*****
TERM=*****
PATH=**************************************************************************************
GOPATH=***
PWD=***
GOLANG_VERSION=****
#
```
## binary
download binary from github release.
```
/ # VERSION=0.0.0
/ # curl -L -s https://github.com/nakatamixi/maskedenv/releases/download/${VERSION}/maskedenv_linux_arm64 -o maskedenv
/ # chmod +x ./maskedenv
/ # ./maskedenv
PWD=*
PATH=************************************************************
TERM=*****
HOME=*****
SHLVL=*
HOSTNAME=************
/ #
```

# Usage
```
Usage of maskedenv:
  -h int
    	unmask head size
  -t int
    	unmask tail size
```
```
$ HOGE=abc ./maskedenv |grep HOGE
HOGE=***
$ HOGE=abc ./maskedenv -h 1  |grep HOGE
HOGE=a**
$ HOGE=abc ./maskedenv -t 1  |grep HOGE
HOGE=**c
```
