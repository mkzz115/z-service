#!/usr/bin/env bash

function start() {
    exec_with_dump
}


function dump_env() {
    echo `go version`
}


function exec_with_dump() {
    dump_env
    exec $1 $2
}

ActionList="start"
function Usage() {
    echo "Usage: $0 {`echo "$ActionList" | sed -e 's|,|\||g'`}"
    exit 1
}

[ $# -ne 3 ] && Usage || action="$3"

if [[ "isAction" == $(echo "$action" |awk -v actionList=$ActionList '{if(index(actionList,$1)>0){print "isAction"}}') ]] ; then
        $action
else
        Usage ;
fi