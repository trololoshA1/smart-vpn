#!/bin/bash
set -e

export PATH=$PATH:/usr/local/go/bin

gomobile init
gomobile bind -target=android -o smartvpn.aar .