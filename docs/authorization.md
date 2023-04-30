# Authorization

This application uses the [OAuth2 for Go] library.

## User Authentication

We will use the [OAuth 2.0 Authorization Framework RFC] which sends the user to an authorization servers, then
redirects them back to this application. Some or all of the following URL query parameters may be required when
redirecting a user to the auth server:

Check wth your OAuth Provider.

| Parameter     | Description                                                                                                                  |
|---------------|------------------------------------------------------------------------------------------------------------------------------|
| client_id     | This application's client ID.                                                                                                |
| scope         | The space-separated scope values to request.                                                                                 |
| state         | A semi-random data blob to send back with when completing authorization.                                                     |
| redirect_uri  | The redirect location. Should either be https://localhost for testing or the redirect_uri configured in the API Access tool. |
| response_type | The code type being requested. For the authorization code flow, this value should be code.                                   |

1. Send the user to the auth token URL.
2. When the user is redirected back, and on success, parse the token returned.
3. Store the token in the user session.
4. Use the token to request access to their profile.
5. Display their profile.

NOTE: HTTPS is REQUIRED for the redirect URL, even running locally. So Use [Simple Golang HTTPS/TLS Server] as a
guide. On Windows 11 try ``

### Setup SSL

#### *nix

1. Make sure **openssl** is installed.
2. Run these command to generate a self-signed certificate:
    ```shell
    openssl genrsa -out certs/server.key 2048
    openssl ecparam -genkey -name secp384r1 -out certs/server.key
    openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.crt -days 3650
    ```
3. Enter the following info for the DN:
    ```shell
    Country Name (2 letter code) [AU]:US
    State or Province Name (full name) [Some-State]:Michigan
    Locality Name (eg, city) []:Detroit
    Organization Name (eg, company) [Internet Widgits Pty Ltd]:Diablo Assistant
    Organizational Unit Name (eg, section) []:Research & Development
    Common Name (e.g. server FQDN or YOUR name) []:localhost
    Email Address []:my@emai.com
    ```

---

[OAuth 2.0 Authorization Framework RFC]: https://datatracker.ietf.org/doc/html/rfc6749#section-1.3.1
[OAuth2 for Go]: https://pkg.go.dev/golang.org/x/oauth2
[Simple Golang HTTPS/TLS Server]: https://gist.github.com/denji/12b3a568f092ab951456
