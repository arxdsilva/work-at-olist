package record

import "testing"

func TestRecord_DataChecks(t *testing.T) {
	type fields struct {
		ID          string
		Type        string
		TimeStamp   string
		CallID      string
		Source      string
		Destination string
		Month       int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "valid end record",
			wantErr: false,
			fields: fields{
				ID:        "123",
				Type:      "end",
				TimeStamp: "2016-02-29T12:00:00Z",
				CallID:    "someID",
			},
		},
		{
			name:    "Valid start record",
			wantErr: false,
			fields: fields{
				ID:          "123",
				Type:        "start",
				TimeStamp:   "2016-02-29T12:00:00Z",
				CallID:      "someID",
				Source:      "1234567890",
				Destination: "1234567890",
			},
		},
		{
			name:    "Invalid start record",
			wantErr: true,
			fields: fields{
				ID:          "123",
				Type:        "start",
				TimeStamp:   "2016-02-29T12:00:00Z",
				CallID:      "someID",
				Source:      "1234567",
				Destination: "1234567890",
			},
		},
		{
			name:    "Invalid start record phone",
			wantErr: true,
			fields: fields{
				ID:          "123",
				Type:        "start",
				TimeStamp:   "2016-02-29T12:00:00Z",
				CallID:      "someID",
				Source:      "1234567aaa",
				Destination: "1234567890",
			},
		},
		{
			name:    "Invalid start record time",
			wantErr: true,
			fields: fields{
				ID:          "123",
				Type:        "start",
				TimeStamp:   "invalid",
				CallID:      "someID",
				Source:      "1234567890",
				Destination: "1234567890",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				ID:          tt.fields.ID,
				Type:        tt.fields.Type,
				TimeStamp:   tt.fields.TimeStamp,
				CallID:      tt.fields.CallID,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
				Month:       tt.fields.Month,
			}
			if err := r.DataChecks(); (err != nil) != tt.wantErr {
				t.Errorf("Record.DataChecks() error = %v, wantErr %v, name: %v", err, tt.wantErr, tt.name)
			}
		})
	}
}
