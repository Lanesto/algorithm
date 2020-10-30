#!/bin/bash
# set -e

service rsyslog start

exec "$@"