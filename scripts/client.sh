#!/usr/bin/env bash
curl -s 'http://127.0.0.1:8010/user' \
    -H 'X-Forwarded-For: 173.242.102.241' \
    -d '{
    "id": "534",
    "email": "anothername@email.com",
    "username": "anotherUserName",
    "password": "ultraPassword",
    "is_admin": false
}'
