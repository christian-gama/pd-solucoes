#!/bin/bash
#
# Script to wait for the database to be up and running
#
# Author: Christian Gama e Silva
# Email: christiangama.dev@gmail.com
# Date Created: 2023/03/21

set -e

container_name="$1"
shift
test_mode="$1"
shift
cmd="$@"

until docker inspect --format "{{.State.Health.Status}}" $container_name | grep "healthy" > /dev/null; do
  >&2 echo "Postgres is unavailable - waiting..."
  sleep 2
done

export $test_mode
exec $cmd