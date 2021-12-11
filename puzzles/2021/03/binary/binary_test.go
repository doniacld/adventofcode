package binary

import (
	"testing"
)

func TestComputeDiags(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    int
		wantErr bool
	}{
		{"nominal case", "./input-test.txt", 198, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags, err := ExtractDiag(tt.file)
			got, err := ComputeDiags(diags)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeDiags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComputeDiags() got = %v, want %v", got, tt.want)
			}
		})
	}
}
