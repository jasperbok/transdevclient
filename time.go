package transdevclient

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
	"time"
)

const TransdevDateFormat = "2006-01-02"

// UnixTimeString is a time.Time that has a JSON representation of a
// Unix timestamp as a string type.
type UnixTimeString time.Time

func (u *UnixTimeString) UnmarshalJSON(b []byte) (err error) {
	timestampString := strings.Trim(string(b), "\"")
	_number, err := strconv.Atoi(timestampString)
	if err != nil {
		return err
	}
	number := int64(_number)

	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return err
	}

	_t := time.Unix(number, 0)
	_t = _t.In(loc)
	*u = UnixTimeString(_t)

	return nil
}

// UnixTimeInt is a time.Time that has a JSON representation of a Unix
// timestamp as an integer type.
type UnixTimeInt time.Time

func (u *UnixTimeInt) UnmarshalJSON(b []byte) (err error) {
	buf := bytes.NewBuffer(b)
	number, err := binary.ReadVarint(buf)
	if err != nil {
		return err
	}

	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return err
	}

	_t := time.Unix(number, 0)
	_t = _t.In(loc)
	*u = UnixTimeInt(_t)

	return nil
}

// Date is t time.Time in which only the date matters.
type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	t, err := time.Parse(TransdevDateFormat, string(b))
	if err != nil {
		return err
	}
	*d = Date(t)

	return nil
}
