#!/bin/sh
/consul-template -consul consul:8500 -template '/haproxy.ctmpl:/etc/haproxy/haproxy.cfg:haproxy -f /etc/haproxy/haproxy.cfg -p /var/run/haproxy.pid&'
