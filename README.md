# Yieldr - Navitaire New Skies&reg; ODS

Upload flight performance from [Navitaire New Skies](http://www.navitaire.com/p_new_skies.aspx) [Operational Data Store](http://www.navitaire.com/Styles/Images/PDFs/Data%20Store.pdf) to Yieldr.

## Download

Head over to the [releases](https://github.com/yieldr/navitaire-ods/releases) page and download the latest version appropriate for your operating system and architecture.

## Usage

The `run` command queries the New Skies ODS database and uploads the flights to Yieldr using the [Yieldr API](https://api.yieldr.com/#272b3d39-2dfc-e7fe-7f65-f2cd4d0e841c).

```bash
navitaire-ods run [options]
```

Refer to the [`run` command documentation](doc/navitaire-ods_run.md) for configuration options.

Command line options can alternatively be replaced with environment variables. The convention is `YIELDR_` followed by the flag name in capital letters with dashes (`-`) replaced by underscores (`_`). For example the `--api-addr` flag can be defined as the `YIELDR_API_ADDR` environment variable. For example

```bash
export YIELDR_API_ADDR="airline.yieldr.com"
export YIELDR_API_CLIENT_ID="2CF99XIIcAUm52O4p2C16RcdExEoCB7D"
export YIELDR_API_CLIENT_SECRET="2Z5XCHysvYM4tZwIxEvj6xnkqlTjSZyvhzU4K8eYsM1Y7d8LQlswhFJMjjC0HgHO"

navitaire-ods run
```

### Running custom SQL queries

A [sample query](pkg/navitaire/ods/query.sql) is supplied for guidance, but in most cases you would want to customise it to match your use case.

The result of the query should contain the following fields in specific order.

|Column|Description|
|-|-|
|`CarrierCode`|Property of the `Booking` table|
|`DepartureStation`|Property of `InventoryLeg` or `PassengerJourneySegment`|
|`ArrivalStation`|Property of `InventoryLeg` or `PassengerJourneySegment`|
|`FlightNumber`|Property of `InventoryLeg`|
|`STD`|Property of `InventoryLeg`|
|`SeatsSold`|Sum of `InventoryLegClass.ClassSold`|
|`SeatsAvailable`|Calculated as the division of `InventoryLeg.Capacity` by the sum of `InventoryLegClass.ClassSold`|
|`Revenue`|Sum of `PassengerJourneyCharge.ChargeAmount`|

You can select a file containing your desired SQL query using the `--db-query` flag or the `YIELDR_DB_QUERY` environment variable. Assuming we have a file called `file.sql` containing the following:

Running the command with the following options will execute a given query instead of the [sample query](pkg/navitaire/ods/query.sql).

```bash
navitaire-ods run --db-query=file.sql --db-query-args=OA
```

The SQL query could also be passed via standard input. Using the `--db-query` flag set to `-` instructs the program to read the query from stdin.

```bash
cat file.sql | navitaire-ods run --db-query="-" --db-query-args=OA
```

## Yieldr API

To find your `client_id` and `client_secret` you will need to create a [Yieldr API Integration](https://help.yieldr.com/yieldr-api/section-heading/step-2-integrate-with-the-yieldr-api).

## SFTP

We plan to add support for uploading files using SFTP. As of this writing this is not supported, but might become available in the future.

For more information on the `run` command, check the reference [documentation](doc/navitaire-ods_run.md).
