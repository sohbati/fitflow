#!/bin/bash

# Test script for Google OAuth authentication
BASE_URL="http://localhost:8091"

echo "Testing Google OAuth Authentication"
echo "=================================="

# Test 1: Get Google OAuth URL
echo "1. Getting Google OAuth URL..."
curl -X GET "$BASE_URL/auth/google/url" \
  -H "Content-Type: application/json" \
  -w "\nStatus: %{http_code}\n" \
  -s

echo -e "\n"

# Test 2: Test Google OAuth login (this will fail without a valid code)
echo "2. Testing Google OAuth login (without valid code)..."
curl -X POST "$BASE_URL/auth/google" \
  -H "Content-Type: application/json" \
  -d '{"code": "invalid_code"}' \
  -w "\nStatus: %{http_code}\n" \
  -s

echo -e "\n"

# Test 3: Health check
echo "3. Health check..."
curl -X GET "$BASE_URL/health" \
  -H "Content-Type: application/json" \
  -w "\nStatus: %{http_code}\n" \
  -s

echo -e "\n"

echo "Google OAuth endpoints available:"
echo "- GET  $BASE_URL/auth/google/url  - Get Google OAuth authorization URL"
echo "- POST $BASE_URL/auth/google       - Login with Google OAuth code"
echo "- GET  $BASE_URL/swagger/index.html - View API documentation"
