# gotemplate

### From mozilla-services:

Containerized App Requirements
When the application is ran in the container it must:

1. Accept its configuration through environment variables.
2. Listen on environment variable $PORT for HTTP requests.
3. Must have a JSON version object at /app/version.json.
4. Respond to `/__version__` with the contents of /app/version.json.
5. Respond to `/__heartbeat__` with a HTTP 200 or 5xx on error. This should check backing services like a database for connectivity and may respond with the status of backing services and application components as a JSON payload.
6. Respond to `/__lbheartbeat__` with an HTTP 200. This is for load balancer checks and should not check backing services.
7. Send text logs to stdout or stderr.
8. Serve its own static content.
