#!/bin/sh

# Start backend in background
/app/backend/backend-app &

# Start nginx in foreground
nginx -g 'daemon off;'
