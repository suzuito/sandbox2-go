# sandbox2-go

## common

### Test

```bash
make common-test
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
make blog-test
```

## crawler

Make develop environment

```bash
make crawler-init
```

### Run

Open filebase UI in local.
http://localhost:8082

Load environment variables.

```bash
cp ./.service/crawler/local.env.sh.sample ./.service/crawler/local.env.sh
vi ./.service/crawler/local.env.sh
source ./.service/crawler/local.env.sh
```

Run crawler app

```bash
./crawler-crawl.exe
```

Run notifier app

```bash
go run internal/crawler/crawler/cmd/local/crawl/*.go
go run internal/crawler/notifier/cmd/local/notify/*.go
```
