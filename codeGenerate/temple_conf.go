package main

var confTemp=`port: 12306
prometheus:
  switch_on: true
  port: 8080
limit:
  switch_on: true
  qps: 10
  size: 1000
  type: lpLimit
  timediff: -1
logs:
  chansize: 100
  loglevel: debug
  servername: test
`