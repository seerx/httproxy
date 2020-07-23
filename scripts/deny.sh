#!/usr/bin/env bash

firewall-cmd --permanent --add-rich-rule="rule family=ipv4 source address=$1 port protocol=tcp port=80 drop"
#  --add-rich-rule='rule family=\"ipv4\" source address=\"$1\" port protocol=\"tcp\" port=\"80\"\' reject"
# echo $cmd
# out=`$cmd`
# echo out
firewall-cmd --reload
