# blog

## Development

Make develop environment

```bash
# Set up local environment
make init
make init-rdb
```

Run & check server

```bash
# env
cp ./local.env.sh.sample ./local.env.sh
## Add GH Token
vim ./local.env.sh
source ./local.env.sh

# server
air -c .air.server.toml
curl http://localhost:8080/ping

# check rdb
docker compose exec rdb mysql
```

Test

```bash
make test
```

Migration

```bash
# Create new migration
migrate create -dir .schema -ext sql init
```

Insert test articles into local db

```bash
go run cmd/util/*.go convert
```

## Operation

### Deployment

#### Upload docker image (CloudBuild)

After a push on main branch, run cloud build and image is uploaded Google cloud container registry.
