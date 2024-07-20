### Set authorized network

Set authorized network. https://github.com/suzuito/sandbox2-terraform/blob/main/organizations/tach.dev/products/products-common/services/mysql_sandbox/environments/common/main.tf

Check IP address of src.

- https://www.ugtop.com/spill.shtml

### Execute go migrate

```bash
export DB_HOST=SetThis(Public IP Address of sql instance)
export DB_USER=app
export DB_PASSWORD=SetThis(in SecretManager)

# For prd
export DB_NAME=photodx-prd

# For stg
undefined
```

```bash
migrate \
-source file://./.service/photodx/db/.schema/ \
-database mysql://${DB_USER}:${DB_PASSWORD}@tcp($DB_HOST:3306)/${DB_NAME} up
```

### Check DB

```bash
mysql --host=${DB_HOST} --user=${DB_USER} -p${DB_PASSWORD} -D${DB_NAME}
```
