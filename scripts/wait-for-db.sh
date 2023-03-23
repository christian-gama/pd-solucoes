#!/bin/bash
#
# Script to wait for the database to be up and running
#
# Author: Christian Gama e Silva
# Email: christiangama.dev@gmail.com
# Date Created: 2023/03/21

#!/bin/bash

container_name="$1"

is_first_time=true

until docker inspect --format "{{.State.Health.Status}}" $container_name | grep "healthy" > /dev/null; do
  if $is_first_time; then
    >&2 echo "Connecting to Postgres - the first time may take a while"
    is_first_time=false
  else
    >&2 echo "Waiting for Postgres to be up and running..."
  fi
  sleep 2
done