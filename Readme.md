## Getting the events

### Configure credentials
- Fill `<client_id>`, `<project_id>` and `<client_secret>` in `events/credentials.json`.

### Run the script
- Run `go run main.go` to serve the HTTP server.

### Get the events
- Go to `http://localhost:8888/events` on the browser.
- Copy the auth URL from the console and paste it on the browser.
- Authorize the app.
- Copy the code from the browser and paste it on the console.
- The events will be printed on the browser `http://localhost:8888/events`.