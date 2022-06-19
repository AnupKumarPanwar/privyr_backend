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

Create uuid extension.

```sql
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Step 3

`conf.ini.sample` file contains environment variable name with dummy values. Copy it to `conf.ini` and replace with actual
values.

```shell script
cp pkg/conf/conf.ini.sample pkg/conf/conf.ini
```

### Step 4

Run server

```shell script
go run main.go
```