# crawler

## Development

Make develop environment

```bash
make init
```

Open filebase UI in local.
http://localhost:8082

Load environment variables.

```bash
cp ./local.env.sh.sample ./local.env.sh
vi ./local.env.sh
source ./local.env.sh
```

Run crawler app

```bash
go run crawler/cmd/local/*.go
```

Run notifier app

```bash
go run notifier/cmd/local/*.go
```
