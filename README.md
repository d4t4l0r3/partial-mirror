# partial-mirror

This software is aimed at hosting partial mirrors for any linux distro (and potentially any other system using similar HTTP(S) repositories).
The server acts as a regular mirror, but only syncs a selected list of packages rather than the entire repository, to preserve bandwidth and storage space. All other packages will only be synced on demand, but then stay cached alongside the pre-fetched packages in accordance with the configured retention policy.

## Configuration

| YAML name         | Env variable name                | Default   | Description                                                              |
| ---               | ---                              | ---       | ---                                                                      |
| `bind_address`    | `PARTIAL_MIRROR_BIND_ADDRESS`    | localhost | The address the mirror will bind to                                      |
| `bind_port`       | `PARTIAL_MIRROR_BIND_PORT`       | 80        | The port the mirror will bind to                                         |
| `upstreams`       |                                  | _empty_   | The upstream mirrors to sync from. The first one reachable will be used. |
| `packages`        |                                  | _empty_   | The packages to keep synced                                              |
| `retention_count` | `PARTIAL_MIRROR_RETENTION_COUNT` | 0         | How many latest versions of a package to keep. 0 = infinite              |
| `retention_time`  | `PARTIAL_MIRROR_RETENTION_TIME`  | 0         | How long to keep a package version after its last request. 0 = infinite  |
| `db.host`         | `PARTIAL_MIRROR_DB_HOST`         | localhost | The address of the database server                                       |
| `db.port`         | `PARTIAL_MIRROR_DB_PORT`         | 5432      | The port of the database server                                          |
| `db.user`         | `PARTIAL_MIRROR_DB_USER`         | mirror    | The user for the database                                                |
| `db.dbname`       | `PARTIAL_MIRROR_DB_DBNAME`       | mirror    | The name of the database                                                 |
| `db.password`     | `PARTIAL_MIRROR_DB_PASSWORD`     | _empty_   | The password for the database user                                       |
| `db.sslmode`      | `PARTIAL_MIRROR_DB_SSLMODE`      | require   | Postgres' [sslmode setting](https://www.postgresql.org/docs/current/libpq-connect.html) |
