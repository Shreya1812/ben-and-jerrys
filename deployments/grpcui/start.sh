echo "$APP_HOST":"$APP_PORT"

grpcui  -port 9999 -bind 0.0.0.0 -plaintext "$APP_HOST":"$APP_PORT"