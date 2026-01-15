# API Endpoints

Server runs on `http://localhost:8080`

## Get User

Retrieve a user by username.

```bash
# Get existing user (karim)
curl "http://localhost:8080/user?username=karim"

# Get non-existent user
curl "http://localhost:8080/user?username=rahim"
```

## Save User

Create or update a user with username and mobile.

```bash
# Save a new user
curl "http://localhost:8080/save?username=rakib&mobile=01933"

# Verify the saved user
curl "http://localhost:8080/user?username=rakib"
```
