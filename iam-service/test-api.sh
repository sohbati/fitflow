#!/bin/bash

# IAM Service API Test Script
# This script demonstrates the usage of the IAM Service API

BASE_URL="http://localhost:8091"

echo "🚀 Testing IAM Service API"
echo "=========================="

# Test health endpoint
echo "1. Testing health endpoint..."
curl -s "$BASE_URL/health" | jq .
echo ""

# Register a new user
echo "2. Registering a new user..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "testuser@example.com",
    "mobile": "+1234567890",
    "display_name": "Test User",
    "password": "testpassword123",
    "country": "US",
    "role": "user"
  }')

echo "$REGISTER_RESPONSE" | jq .
echo ""

# Login to get JWT token
echo "3. Logging in to get JWT token..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpassword123"
  }')

echo "$LOGIN_RESPONSE" | jq .

# Extract token from response
TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')
echo ""

if [ "$TOKEN" != "null" ] && [ "$TOKEN" != "" ]; then
  echo "4. Testing protected endpoint with JWT token..."
  curl -s -X GET "$BASE_URL/me" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
else
  echo "❌ Failed to get JWT token"
fi

echo "5. Testing Swagger documentation..."
echo "Visit: $BASE_URL/swagger/index.html"
echo ""

echo "✅ API testing completed!"
