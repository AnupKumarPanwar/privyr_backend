# privyr

## Setup

### Step 1

Create database and a user in postgres.

```sql
CREATE
DATABASE privyr;
```

```sql
CREATE
USER privyr WITH PASSWORD '#Privyr@123';
```

```sql
GRANT
ALL
PRIVILEGES
ON
DATABASE
privyr TO privyr;
```

```sql
ALTER
USER privyr WITH SUPERUSER;
```

### Step 2

Connect to `privy` database;

```
\c privyr
```

Create uuid extension.

```sql
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Step 3

`conf.ini.sample` file contains environment variable name with dummy values. Copy it to `conf.ini` and replace with actual
values.

```shell script
cp conf.ini.sample conf.ini
```

### Step 4

Run Migrations

```shell script
go run main.go db:migrate
```

Other db commands (NOT TO RUN UNLESS REQUIRED)
```shell script
go run main.go db:drop
```
```
go run main.go db:create
```


### Step 5

Run server

```shell script
go run main.go
```