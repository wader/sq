package kind_test

import (
	stdj "encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/neilotoole/sq/libsq/core/kind"
)

func TestKind(t *testing.T) {
	testCases := map[kind.Kind]string{
		kind.Unknown:  "unknown",
		kind.Null:     "null",
		kind.Text:     "text",
		kind.Int:      "int",
		kind.Float:    "float",
		kind.Decimal:  "decimal",
		kind.Bool:     "bool",
		kind.Datetime: "datetime",
		kind.Date:     "date",
		kind.Time:     "time",
		kind.Bytes:    "bytes",
	}

	for knd, testText := range testCases {
		knd, testText := knd, testText

		t.Run(knd.String(), func(t *testing.T) {
			gotBytes, err := knd.MarshalText()
			require.NoError(t, err)
			require.Equal(t, testText, string(gotBytes))

			gotString := knd.String()
			require.Equal(t, testText, gotString)

			gotJSON, err := knd.MarshalJSON()
			require.NoError(t, err)
			require.Equal(t, `"`+testText+`"`, string(gotJSON))

			var dt2 kind.Kind
			require.NoError(t, dt2.UnmarshalText([]byte(testText)))
			require.True(t, knd == dt2)
		})
	}

	d := kind.Kind(666)
	bytes, err := d.MarshalText()
	require.Error(t, err)
	require.Nil(t, bytes)

	bytes, err = d.MarshalJSON()
	require.Error(t, err)
	require.Nil(t, bytes)

	d = kind.Bytes // pick any valid type
	require.Error(t, d.UnmarshalText([]byte("invalid_text")))
	require.Equal(t, kind.Bytes, d, "d should not be mutated on UnmarshalText err")
}

func TestKindDetector(t *testing.T) {
	const (
		fixtTime1              = "00:00:00"
		fixtTime2              = "08:30:05"
		fixtTime3              = "15:30"
		fixtTime4              = "7:15PM"
		fixtDate1              = "1970-01-01"
		fixtDate2              = "1989-11-09"
		fixtDate3              = "02 Jan 2006"
		fixtDate4              = "2006/01/02"
		fixtDatetime1          = "1970-01-01T00:00:00Z" // RFC3339Nano
		fixtDatetime2          = "1989-11-09T00:00:00Z"
		fixtDatetimeAnsic      = "Mon Jan 2 15:04:05 2006"
		fixtDatetimeUnix       = "Mon Jan 2 15:04:05 MST 2006"
		fixtDatetimeRFC3339    = "2002-10-02T10:00:00-05:00"
		fixtDatetimeStamp      = "Jan 2 15:04:05"
		fixtDatetimeStampMilli = "Jan 2 15:04:05.000"
		fixtDatetimeStampMicro = "Jan 2 15:04:05.000000"
		fixtDatetimeStampNano  = "Jan 2 15:04:05.000000000"
	)

	testCases := []struct {
		in        []interface{}
		want      kind.Kind
		wantMunge bool
		wantErr   bool
	}{
		{in: nil, want: kind.Null},
		{in: []interface{}{}, want: kind.Null},
		{in: []interface{}{""}, want: kind.Text},
		{in: []interface{}{nil}, want: kind.Null},
		{in: []interface{}{nil, ""}, want: kind.Text},
		{in: []interface{}{int(1), int8(8), int16(16), int32(32), int64(64)}, want: kind.Int},
		{in: []interface{}{1, "2", "3"}, want: kind.Decimal},
		{in: []interface{}{"99999999999999999999999999999999999999999999999999999999"}, want: kind.Decimal},
		{in: []interface{}{"99999999999999999999999999999999999999999999999999999999xxx"}, want: kind.Text},
		{in: []interface{}{1, "2", stdj.Number("1000")}, want: kind.Decimal},
		{in: []interface{}{1.0, "2.0"}, want: kind.Decimal},
		{in: []interface{}{1, float64(2.0), float32(7.7), int32(3)}, want: kind.Float},
		{in: []interface{}{nil, nil, nil}, want: kind.Null},
		{in: []interface{}{"1.0", "2.0", "3.0", "4", nil, int64(6)}, want: kind.Decimal},
		{in: []interface{}{true, false, nil, "true", "false", "yes", "no", ""}, want: kind.Bool},
		{in: []interface{}{"0", "1"}, want: kind.Decimal},
		{in: []interface{}{fixtTime1, nil, ""}, want: kind.Time, wantMunge: true},
		{in: []interface{}{fixtTime2}, want: kind.Time, wantMunge: true},
		{in: []interface{}{fixtTime3}, want: kind.Time, wantMunge: true},
		{in: []interface{}{fixtTime4}, want: kind.Time, wantMunge: true},
		{in: []interface{}{fixtDate1, nil, ""}, want: kind.Date, wantMunge: true},
		{in: []interface{}{fixtDate2}, want: kind.Date, wantMunge: true},
		{in: []interface{}{fixtDate3}, want: kind.Date, wantMunge: true},
		{in: []interface{}{fixtDate4}, want: kind.Date, wantMunge: true},
		{in: []interface{}{fixtDatetime1, nil, ""}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetime2}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeAnsic}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeUnix}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{time.RubyDate}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{time.RFC822}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{time.RFC822Z}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{time.RFC850}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{time.RFC1123}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{time.RFC1123Z}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeRFC3339}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeStamp}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeStampMilli}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeStampMicro}, want: kind.Datetime, wantMunge: true},
		{in: []interface{}{fixtDatetimeStampNano}, want: kind.Datetime, wantMunge: true},
	}

	for i, tc := range testCases {
		tc := tc

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			kd := kind.NewDetector()

			for _, val := range tc.in {
				kd.Sample(val)
			}

			gotKind, gotMungeFn, gotErr := kd.Detect()
			if tc.wantErr {
				require.Error(t, gotErr)
				return
			}

			require.Equal(t, tc.want.String(), gotKind.String(), tc.in)

			if !tc.wantMunge {
				require.Nil(t, gotMungeFn)
			} else {
				require.NotNil(t, gotMungeFn)
				for _, val := range tc.in {
					_, err := gotMungeFn(val)
					require.NoError(t, err)
				}
			}
		})
	}
}
