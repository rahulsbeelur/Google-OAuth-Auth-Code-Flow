## Google OAuth with Auth Code Flow in React and Go lang

Tech used - 
- react
- Vite
- Golang
- Gin
- net/http
- gin-contrib/cors
- joho/godotenv
- /x/oauth2
- /x/oauth2/google


This is to generate a jwt token with the code received from Google OAuth. Once the code is received in the frontend, a network call to endpoint exposed which exhanges this code with the provider to get back the token. 
This token has access token which inturn brings back user info. This is used to generate jwt and signed and sent back to the frontend.

This jwt can be used in every request, where the backend can confirm the user is legitimate.

**This is just with the access token and does not contain refresh token**
