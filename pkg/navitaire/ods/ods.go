//go:generate go-bindata -pkg ods -o query.go query.sql
package ods

import (
	"database/sql"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/pkg/errors"
)

type ODSConfig struct {
	Driver   string
	Addr     string
	User     string
	Password string
	Database string
}

type ODS struct {
	db *sql.DB
}

func New(c *ODSConfig) (*ODS, error) {

	q := url.Values{}
	q.Set("database", c.Database)

	url := url.URL{
		Scheme:   c.Driver,
		User:     url.UserPassword(c.User, c.Password),
		Host:     c.Addr,
		RawQuery: q.Encode(),
	}

	db, err := sql.Open(c.Driver, url.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed connecting to the database server")
	}

	return &ODS{db}, nil
}

func (ods *ODS) Query(query string, args ...string) ([]*Flight, error) {

	// Execute the query against the database.
	iargs := make([]interface{}, len(args))
	for i := range args {
		iargs[i] = args[i]
	}

	rows, err := ods.db.Query(query, iargs...)
	if err != nil {
		return nil, errors.Wrap(err, "failed querying the database")
	}
	defer rows.Close()

	// Make a slice of Flight structs to store the result in.
	flights := make([]*Flight, 0, 100)

	for rows.Next() {
		// Iterate over every row in the result set. Scan each row into a Flight
		// struct and append it to the slice.
		f := &Flight{}
		err := rows.Scan(
			&f.carrier,
			&f.origin,
			&f.destination,
			&f.flightNumber,
			&f.dateDeparture,
			&f.seatsSold,
			&f.seatsAvailable,
			&f.revenue,
			&f.ancillaryRevenue)
		if err != nil {
			return nil, errors.Wrap(err, "failed scanning row into a flight struct")
		}

		if len(flights) == cap(flights) {
			tmpFlights := make([]*Flight, len(flights), (cap(flights)+1)*2)
			copy(tmpFlights, flights)
			flights = tmpFlights
		}

		flights = append(flights, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return flights, nil
}

func DefaultQuery() []byte {
	return MustAsset("query.sql")
}
