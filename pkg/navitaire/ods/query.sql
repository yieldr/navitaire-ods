DECLARE @carrier_code VARCHAR(2) = @p1;

SELECT
  i.CarrierCode      AS carrier,
  i.DepartureStation AS origin,
  i.ArrivalStation   AS destination,
  i.FlightNumber     AS flight_number,
  i.STD              AS date_departure,
  i.SeatsSold        AS seats_sold,
  i.SeatsAvailable   AS seats_available,
  r.Revenue          AS revenue,
  ar.Revenue         AS ancillary_revenue
FROM (
       SELECT
         il.CarrierCode,
         il.DepartureStation,
         il.ArrivalStation,
         il.FlightNumber,
         il.STD,
         SUM(ic.ClassSold)               AS SeatsSold,
         il.Capacity - SUM(ic.ClassSold) AS SeatsAvailable
       FROM InventoryLeg AS il
         INNER JOIN InventoryLegClass AS ic WITH ( NOLOCK )
           ON il.InventoryLegID = ic.InventoryLegID
       WHERE
         il.CarrierCode = @carrier_code
         AND il.Status <= 1
       GROUP BY
         il.CarrierCode,
         il.DepartureStation,
         il.ArrivalStation,
         il.STD,
         il.FlightNumber,
         il.Capacity
     ) AS i
  INNER JOIN
  (
    SELECT
      il.CarrierCode,
      il.DepartureStation,
      il.ArrivalStation,
      il.FlightNumber,
      il.STD,
      SUM(pjc.ChargeAmount) AS Revenue
    FROM Booking AS b
      INNER JOIN BookingPassenger AS bp WITH ( NOLOCK )
        ON b.BookingID = bp.BookingID
      INNER JOIN PassengerJourneySegment AS pjs WITH ( NOLOCK )
        ON bp.PassengerID = pjs.PassengerID
      INNER JOIN PassengerJourneyCharge AS pjc WITH ( NOLOCK )
        ON pjs.PassengerID = pjc.PassengerID
           AND pjs.SegmentID = pjc.SegmentID
      INNER JOIN PassengerJourneyLeg AS pjl WITH ( NOLOCK )
        ON pjs.PassengerID = pjl.PassengerID
           AND pjs.SegmentID = pjl.SegmentID
      INNER JOIN InventoryLeg AS il WITH ( NOLOCK )
        ON pjl.InventoryLegID = il.InventoryLegID
           AND pjl.LegNumber = il.LegNumber
    WHERE
      il.CarrierCode = @carrier_code
      AND il.Status <= 1
    GROUP BY
      il.CarrierCode,
      il.DepartureStation,
      il.ArrivalStation,
      il.FlightNumber,
      il.STD
  ) AS r
    ON i.CarrierCode = r.CarrierCode
       AND i.DepartureStation = r.DepartureStation
       AND i.ArrivalStation = r.ArrivalStation
       AND i.FlightNumber = r.FlightNumber
       AND i.STD = r.STD
  INNER JOIN
  (
    SELECT
      il.CarrierCode,
      il.DepartureStation,
      il.ArrivalStation,
      il.FlightNumber,
      il.STD,
      SUM(pfc.ChargeAmount) AS Revenue
    FROM Booking AS b
      INNER JOIN BookingPassenger AS bp WITH ( NOLOCK )
        ON b.BookingID = bp.BookingID
      INNER JOIN PassengerFee AS pf WITH ( NOLOCK )
        ON bp.PassengerID = pf.PassengerID
      INNER JOIN PassengerFeeCharge AS pfc WITH ( NOLOCK )
        ON pf.PassengerID = pfc.PassengerID
           AND pf.FeeNumber = pfc.FeeNumber
      INNER JOIN PassengerJourneySegment AS pjs WITH ( NOLOCK )
        ON bp.PassengerID = pjs.PassengerID
      INNER JOIN PassengerJourneyLeg AS pjl WITH ( NOLOCK )
        ON pjs.PassengerID = pjl.PassengerID
           AND pjs.SegmentID = pjl.SegmentID
      INNER JOIN InventoryLeg AS il WITH ( NOLOCK )
        ON pf.InventoryLegID = il.InventoryLegID
           AND pjl.LegNumber = il.LegNumber
    WHERE
      il.CarrierCode = @carrier_code
      AND il.Status <= 1
    GROUP BY
      il.CarrierCode,
      il.DepartureStation,
      il.ArrivalStation,
      il.FlightNumber,
      il.STD

  ) AS ar
    ON i.CarrierCode = ar.CarrierCode
       AND i.DepartureStation = ar.DepartureStation
       AND i.ArrivalStation = ar.ArrivalStation
       AND i.FlightNumber = ar.FlightNumber
       AND i.STD = ar.STD;
