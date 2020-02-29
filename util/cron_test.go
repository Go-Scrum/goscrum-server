package util

import "testing"

func TestCronExpression(t *testing.T) {
	type args struct {
		hoursMinutes string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "happy path",
			args:    args{hoursMinutes: "10:30"},
			want:    "30 10 * * 1,2,3,4,5",
			wantErr: false,
		}, {
			name:    "sad path",
			args:    args{hoursMinutes: "10"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CronExpression(tt.args.hoursMinutes)
			if (err != nil) != tt.wantErr {
				t.Errorf("CronExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CronExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}
