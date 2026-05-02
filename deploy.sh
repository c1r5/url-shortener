#!/usr/bin/env sh
set -eu

if [ -z "${DEPLOY_HOST:-}" ]; then
  echo "DEPLOY_HOST is required. Example: DEPLOY_HOST=my-vps ./deploy.sh" >&2
  exit 1
fi

DEPLOY_DIR="${DEPLOY_DIR:-~/url-shortener}"

make build

ssh "$DEPLOY_HOST" "mkdir -p $DEPLOY_DIR"
rsync -avz bin/server "$DEPLOY_HOST:$DEPLOY_DIR/server"
rsync -avz compose.yml .env "$DEPLOY_HOST:$DEPLOY_DIR/"
ssh "$DEPLOY_HOST" "cd $DEPLOY_DIR && docker compose up -d"
