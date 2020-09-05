#!/usr/bin/env bash
make docker
docker-compose -f docker-composer.yml up -d