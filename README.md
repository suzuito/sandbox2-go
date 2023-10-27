# sandbox2-go

## common

### Test

```bash
make test
```

## blog

```bash
make blog-init
make blog-init-rdb
```

### Run

```bash
# env
cp ./.service/blog/local.env.sh.sample ./.service/blog/local.env.sh
## Add GH Token
vim ./.service/blog/local.env.sh
source ./.service/blog/local.env.sh

# server
air -c ./.service/blog/.air.server.toml
curl http://localhost:8080/ping

# check rdb
docker compose exec blog-mysql mysql
```

Insert test articles into local db

- Login as admin
  - Access /admin/login
  - Input password
  - Click Login
- Import markdown on repository
  - Access /admin/
  - Click `*`

Migration

```bash
# Create new migration
migrate create -dir .schema -ext sql init
```

### Test

```bash
make test
```
