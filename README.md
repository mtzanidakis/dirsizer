# dirsizer

Send a notification via e-mail when a directory exceeds a certain size.

## Configuration

The following configuration options are available, passed as environment variables:

- `DIRECTORY`: The path to the directory to monitor; defaults to `.`.
- `IDENTIFIER`: A unique identifier for the directory, used in the e-mail subject.
- `MAIL_FROM`: The e-mail address to send the notification from; defaults to `dirsizer@localhost`.
- `MAIL_TO`: The e-mail address to send the notification to; defaults to `root`.
- `SMTP_SERVER`: The SMTP server to use for sending the e-mail; defaults to `localhost:25`.
- `THRESHOLD`: The size in bytes at which to send the notification; defaults to `500M`.

## Usage

Dirsizer can be used in docker compose as a sidecar to any service. For an example check the [docker-compose.yml](docker-compose.yml).
