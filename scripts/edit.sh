#!/usr/bin/env bash
curl -v -X POST 'http://127.0.0.1:8080/user/2' \
    -H 'Authorization: Basic dXNlcjE6cGFzc3dvcmQx' \
    -H 'Content-Type: application/json' \
    -d '{
    "id": "534",
    "email": "anothername@email.com",
    "username": "anotherUserName",
    "password": "ultraPassword",
    "is_admin": true
}'