#!/bin/bash
#
# Script to generate .env files for the project.
#
# Author: Christian Gama e Silva
# Email: christiangama.dev@gmail.com
# Date Created: 2023/03/22

create_env_file() {
    env_name="$1"
    env_file=".env.$env_name"

    if [ ! -f "$env_file" ]; then
        cp .env.example "$env_file"
        sed -i "s/APP_ENV=.*/APP_ENV=$env_name/" "$env_file"
        
        if [[ "$env_name" == "dev" || "$env_name" == "prod" ]]; then
            sed -i "s/DB_HOST=.*/DB_HOST=psql/" "$env_file"
        fi

        echo "Created $env_file"
    else
        echo "$env_file already exists. Skipping."
    fi
}

# Create .env.dev, .env.test, and .env.prod files
create_env_file "dev"
create_env_file "test"
create_env_file "prod"