#!/usr/bin/bash

BASE_URL="http://localhost:8080"

function curl_api() {
    method=$1
    url=$2
    user_id=$3

    output=$(curl -s -X "$method" --cookie "user-id=$user_id" "$url")
    echo -e "\t$output"
}

echo -e "GET Products with authorized user (user1)"
curl_api "GET" "$BASE_URL/api/v1/products" "user1"

echo -e "\nPOST Products with unauthorized (user1)"
curl_api "POST" "$BASE_URL/api/v1/products" "user1"

echo -e "\nPOST Products with authorized user (user2)"
curl_api "POST" "$BASE_URL/api/v1/products" "user2"

echo -e "\nPOST Products with non-existent user (user3)"
curl_api "POST" "$BASE_URL/api/v1/products" "user3"
