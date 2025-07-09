#!/usr/bin/env bash

DOCKER_COMPOSE="docker-compose.yaml"

docker compose -f "$DOCKER_COMPOSE" -p "$GITHUB_SHA" down
