package ods

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Flight struct {
	carrier          sql.NullString
	origin           sql.NullString
	destination      sql.NullString
	flightNumber     sql.NullString
	dateDeparture    time.Time
	seatsSold        sql.NullInt64
	seatsAvailable   sql.NullInt64
	revenue          sql.NullFloat64
	ancillaryRevenue sql.NullFloat64
}

func (f *Flight) Carrier() string           { return f.carrier.String }
func (f *Flight) Origin() string            { return f.origin.String }
func (f *Flight) Destination() string       { return f.destination.String }
func (f *Flight) FlightNumber() string      { return f.flightNumber.String }
func (f *Flight) DateDeparture() time.Time  { return f.dateDeparture }
func (f *Flight) SeatsSold() int64          { return f.seatsSold.Int64 }
func (f *Flight) SeatsAvailable() int64     { return f.seatsAvailable.Int64 }
func (f *Flight) Revenue() float64          { return f.revenue.Float64 }
func (f *Flight) AncillaryRevenue() float64 { return f.ancillaryRevenue.Float64 }

func (f *Flight) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"carrier":           f.Carrier(),
		"origin":            f.Origin(),
		"destination":       f.Destination(),
		"flight_number":     f.FlightNumber(),
		"date_departure":    f.DateDeparture(),
		"seats_sold":        f.SeatsSold(),
		"seats_available":   f.SeatsAvailable(),
		"revenue":           f.Revenue(),
		"ancillary_revenue": f.AncillaryRevenue(),
	}
	return json.Marshal(v)
}
