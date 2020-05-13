# loganalyzer

Analyze access logs (Apache, nginx , load balancers) for quick troubleshoot

## Contentes

1. [Usage](#usage)
    1. [Binaries](#binaries)
        1. [Install](#install)
        1. [Run](#run)
            1. [`top` List top 10 IPs, requests, etc](#top)
            1. [`bandwidth` Calculate total bandwidth](#bandwidth)
            1. [`custom` search for a custom field](#custom)

## Usage

### Binaries

#### Install

Download the most recent release binaries for your OS from https://github.com/cloudsark/loganalyzer/releases

Install each binary as follows,

1. Move binary to '/usr/local/bin'
```bash
$ sudo mv loganalyzer /usr/local/bin 
```
2. Verify the binary file is working
```bash
$ loganalyzer --help
Access log analyzer can be used to analyze web servers , load balancers access logs

Usage:
  loganalyzer [flags]
  loganalyzer [command]

Available Commands:
  bandwidth   Print total bandwidth
  custom      Search for a custom field
  help        Help about any command
  top         Print top 10 IPs, Requests , etc

Flags:
  -f, --file string   access log filename
  -h, --help          help for loganalyzer
  -v, --version       version for loganalyzer
Use "loganalyzer [command] --help" for more information about a command.
```

#### Run

##### top

1. Get help
```bash
$ loganalyzer top --help
Get useful information about your web server, top command is used to print top 10 IPs, methods, requests and status codes

Usage:
  loganalyzer top [command]

Available Commands:
  ip          Print Top 10 IP addresses accessing your web server
  ip2loc      Print Top 10 IP addresses accessing your web server with their location
  method      Print Top 10 HTTP methods used
  request     Print Top 10 requests
  status      Print Top 10 status codes

Flags:
  -h, --help   help for top

Global Flags:
  -f, --file string   access log filename

Use "loganalyzer top [command] --help" for more information about a command.
```
2. Print Top 10 IP addresses accessing your web server
```bash
$ loganalyzer -f access.log top ip
```
3. Print Top 10 IP addresses accessing your web server with their location
```bash
$ loganalyzer -f access.log top ip2loc
```
4. Print Top 10 HTTP methods used
```bash
$ loganalyzer -f access.log top method
```
5. Print Top 10 requests
```bash
$ loganalyzer -f access.log top request
```
6. Print Top 10 status codes
```bash
$ loganalyzer -f access.log top status
```


##### bandwidth
1. Pring total bandwidth
```bash
$ loganalyzer -f access.log bandwidth
```

##### custom
1. Search for a custom field. Provide the regex for the field and loganalyzer will look for it
```bash
$ loganalyzer -f apache-daily-access.log custom --field-regex '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}'
```

