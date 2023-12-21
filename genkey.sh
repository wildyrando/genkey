#!/bin/bash

genkey() {
    lcnkey=""
    for ((anu=0; anu<4; anu++)); do
        block=$(cat /dev/urandom | tr -dc 'A-Z0-9' | fold -w 4 | head -n 1)
        lcnkey+="$block-"
    done
    lcnkey=${lcnkey::-1}
    echo "$lcnkey"
}

genkey
