package cef_test

import (
	"errors"
	"testing"

	"go.m8.ru/cef"
)

func TestCEF_UnmarshalText(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		text    string
		wantErr error
	}{
		{
			name:    "number of header fields less than 7",
			text:    "",
			wantErr: cef.ErrHeaderFieldsNum,
		},
		{
			name:    "number of header fields with extension less than 7",
			text:    "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|dst=2.1.2.2 spt=1232 src=10.0.0.1",
			wantErr: cef.ErrHeaderFieldsNum,
		},
		{
			name:    "bad CEF version",
			text:    "BAD|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1",
			wantErr: cef.ErrBadCEFVersion,
		},
		{
			name:    "bad CEF version",
			text:    "CEF:BAD|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1",
			wantErr: cef.ErrBadCEFVersion,
		},
		{
			name:    "bad agent severity",
			text:    "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|Very-High|dst=2.1.2.2 spt=1232 src=10.0.0.1",
			wantErr: cef.ErrBadAgentSeverity,
		},
		{
			name:    "bad extension",
			text:    "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1 cfp1=bad",
			wantErr: cef.ErrBadExtension,
		},
		{
			name:    "ok without extension",
			text:    "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10",
			wantErr: nil,
		},
		{
			name:    "ok with extension with space",
			text:    "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|fname=C:\\Program Files\\ArcSight",
			wantErr: nil,
		},
		{
			name:    "ok",
			text:    "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1",
			wantErr: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			log := new(cef.CEF)

			gotErr := log.UnmarshalText([]byte(tc.text))
			t.Log(log.FName())
			if !errors.Is(gotErr, tc.wantErr) {
				t.Fatalf("got error %v, want %v", gotErr, tc.wantErr)
			}
		})
	}
}
