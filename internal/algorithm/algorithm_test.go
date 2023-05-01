package algorithm

import (
	"advocate-back/internal/bot"
	"advocate-back/internal/states"
	"testing"
)

func TestCheckAlgorithm(t *testing.T) {
	type args struct {
		s states.BotStates
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "t1",
			args: args{
				s: bot.MockBotStates,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckAlgorithm(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CheckAlgorithm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
