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
				TimeStamp: "something",
				CallID:    "someID",
			},
		},
		{
			name:    "Invalid end record",
			wantErr: true,
			fields: fields{
				Type:      "end",
				TimeStamp: "something",
				CallID:    "someID",
			},
		},
		{
			name:    "Valid start record",
			wantErr: false,
			fields: fields{
				ID:          "123",
				Type:        "start",
				TimeStamp:   "something",
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
				TimeStamp:   "something",
				CallID:      "someID",
				Source:      "1234567",
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
			}
			if err := r.DataChecks(); (err != nil) != tt.wantErr {
				t.Errorf("Record.DataChecks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
