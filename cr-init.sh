#!/bin/sh

./cockroach sql --insecure --host="$COCKROACH_HOSTNAME" < /cockroach/cr-init.sql