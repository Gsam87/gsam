package userservice

import "testing"

func TestCheckID(t *testing.T) {
	service := NewUserService()

	tests := []struct {
		name      string
		inputID   string
		expectErr bool
	}{
		{"valid ID", "12345", false},
		{"empty ID", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CheckID(tt.inputID)
			if (err != nil) != tt.expectErr {
				t.Errorf("CheckID(%q) error = %v, expectErr = %v", tt.inputID, err, tt.expectErr)
			}
		})
	}
}
