# Yieldr - Navitaire New Skies&reg; ODS

Upload flight performance from [Navitaire New Skies](http://www.navitaire.com/p_new_skies.aspx) [Operational Data Store](http://www.navitaire.com/Styles/Images/PDFs/Data%20Store.pdf) to Yieldr.

## Download

Head over to the [releases](https://github.com/yieldr/navitaire-ods/releases) page and download a version appropriate for your operating system and architecture.

```
curl -sSL -o navitaire-ods https://github.com/yieldr/navitaire-ods/releases/download/<VERSION>/navitaire-ods-<OS>-<ARCH>
```

## Usage

The `run` command queries the New Skies ODS database and uploads the flights to Yieldr using the [Yieldr API](https://api.yieldr.com/#272b3d39-2dfc-e7fe-7f65-f2cd4d0e841c).

```bash
navitaire-ods run \
	--api \
	--api-addr="<account>.yieldr.com" \
	--api-client-id=<client-id> \
	--api-client-secret=<client-secret> \
	--api-project-id=<project-id> \
	--db="sqlserver" \
	--db-addr="localhost:1234" \
	--db-name="NS34ODS" \
	--db-user=<username> \
	--db-pass=<password> \
	--db-query="" \ # use the default query, see below
	--db-query-args=<carrier-code>
```

Flags can alternatively be replaced with environment variables. The convention is `YIELDR_` followed by the flag name in capital letters with dashes (`-`) replaced by undescores (`_`). For example the `--api-addr` flag can be defined as the `YIELDR_API_ADDR` environment variable.

### SQL Query

A [sample SQL query](pkg/navitaire/ods/query.sql) is supplied for guidance, but in most cases you would want to customise it to match your use case.

## Yieldr API

To find your `client_id` and `client_secret` you will need to create a [Yieldr API Integration](https://help.yieldr.com/yieldr-api/section-heading/step-2-integrate-with-the-yieldr-api).

## SFTP

We plan to add support for uploading files using SFTP. As of this writing this is not supported, but might become available in the future.

For more information on the `run` command, check the reference [documentation](doc/navitaire-ods_run.md).

