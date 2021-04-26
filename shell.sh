#!/usr/bin/env bash

case $1 in
"commit")
  git add .
  git commit -m "$2"
  git pull
  git push
;;
"dhcp")
  go build cmd/dhcp/$2/main.go
  sudo ./main
  rm -rf main
;;
"netlink")
  go build cmd/netlink/main.go
  sudo ./main
  rm -rf main
;;
"dns")
  go build cmd/dns/$2/main.go
  sudo ./main --help
  rm -rf main
;;
"vlan")
vconfig add tap0 100 
vconfig add tap0 200
vconfig set_flag tap0.100 1 1 
vconfig set_flag tap0.200 2 2
#ifconfig tap0 0.0.0.0 
ifconfig tap0.100 192.168.147.50 netmask 255.255.255.0 up
ifconfig tap0.200 192.168.147.51 netmask 255.255.255.0 up
#vconfig rem tap0.100
#vconfig rem tap0.200
;;
esac

