package udomain

import (
	"testing"
)

func TestDomainNormal(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "a.B.c", args: args{domain: "a.B.c"}, want: "b.c"},
		{name: "adsfasd-a.B.com", args: args{domain: "adsfasd-a.B.com"}, want: "b.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainNormal(tt.args.domain); got != tt.want {
				t.Errorf("DomainNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}
