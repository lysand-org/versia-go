#!/bin/sh

sudo socat TCP-LISTEN:443,reuseaddr,fork TCP:localhost:8443