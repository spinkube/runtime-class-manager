#!/usr/bin/env bash
set -euo pipefail

declare -A levels=([DEBUG]=0 [INFO]=1 [WARN]=2 [ERROR]=3)
script_logging_level="INFO"

log() {
    local log_message=$1
    local log_priority=$2

    #check if level exists
    [[ ${levels[$log_priority]} ]] || return 1

    #check if level is enough
    (( ${levels[$log_priority]} < ${levels[$script_logging_level]} )) && return 2

    #log here
    d=$(date '+%Y-%m-%dT%H:%M:%S')
    echo -e "${d}\t${log_priority}\t${log_message}"
}

log "start downloading shim from  ${SHIM_LOCATION}..." "INFO"

mkdir -p /assets

curl -sL "${SHIM_LOCATION}"  | tar -xzf - -C /assets
log "download successful:" "INFO"

ls -lah /assets
