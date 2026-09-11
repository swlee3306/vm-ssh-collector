# vm-ssh-collector

Go experiment for collecting VM filesystem information over SSH and periodically forwarding results.

## Scope

The collector reads target connection information from a database, connects over SSH, collects filesystem data and forwards a structured payload. HTTP handlers can execute commands on one or multiple targets.

**The command-execution endpoints are a sensitive development interface.** This README does not claim production-ready authorization or isolation. Review authentication, allowed commands, SSH host-key verification and logging before any controlled deployment.

## Configuration

The default filename is `setting.yml`. Keep credentials in an untracked local file or environment variables. An illustrative, nonworking configuration:

```yaml
database:
  dsn: "example_user:REPLACE_LOCALLY@tcp(db.example.com:3306)/example_db?parseTime=true"
collector:
  collectorTime: 1
```

`BATON_SETTING_FILENAME` selects the local configuration file; `BATON_DB_DSN` can override the DSN. Do not commit resolved values. Live execution requires the expected database schema and controlled SSH targets.

## Inspect before running

```sh
go build -mod=readonly ./...
```

Review `collector/`, `internal/ssh/`, `internal/dblinker/` and `utils/router/`. Build success does not verify actual connectivity. This repository is not presented as a secure remote-command service.

[Project map](https://github.com/swlee3306#selected-work)
