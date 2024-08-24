# Versia-Go

Versia-Go is a experimental implementation of the (not yet renamed :P) [Versia](https://versia.pub) protocol written in
Go.

> Compatibility level: Versia Working Draft 4.0

> ⚠️ This project is still in development and is not ready for production use.
> In this phase no pull requests will be accepted and code may often break.

## Developing

### Requirements

- Go 1.22.5+
- Docker + Docker Compose v2

### Running

```shell
git clone https://github.com/lysand-org/versia-go.git
cd versia-go

docker compose up -d nats

touch .env.local
# Add the changed variables from .env to .env.local

go run .
```

## TODO

- [ ] Notes
  - [ ] API
    - [ ] Allow choosing the publishing user
  - [x] Federating notes
- [ ] Follows
  - [ ] API
  - [x] Automatic follows for public users
  - [ ] Unfollows (scheduled for Lysand Working Draft 4)
    - [ ] API
- [ ] Users
  - [ ] API
    - [x] Create user
  - [ ] Lysand API
    - [x] Get user (from local)
    - [x] Webfinger
    - [ ] Inbox handling
      - [ ] Federated notes
      - [ ] Federated unfollows
      - [x] Federated follows
  - [x] Receiving federated users
- [ ] Web
- Extensions
  - [ ] Emojis

## License

Versia-Go is licensed under the GNU Affero General Public License v3.0.

See [LICENSE](LICENSE) for more information.

> ℹ️ This project might get relicensed to a different license in the future.
