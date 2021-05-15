package sqlexplain

import "testing"

func TestExplain(t *testing.T) {
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test",
			args: args{
				sql:  "SELECT * FROM users where email = ?",
				args: []interface{}{"a@a.com"},
			},
			want: []byte("SELECT * FROM users where email = 'a@a.com'"),
		},
		{
			name: "Test",
			args: args{
				sql:  "SELECT * FROM users where email = ? and uuid = ?",
				args: []interface{}{"a@a.com", "uuid"},
			},
			want: []byte("SELECT * FROM users where email = 'a@a.com' and uuid = 'uuid'"),
		},
		{
			name: "Select all test",
			args: args{
				sql:  "SELECT * FROM users",
				args: []interface{}{"a@a.com"},
			},
			want: []byte("SELECT * FROM users"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Explain(tt.args.sql, tt.args.args...); string(got) != string(tt.want) {
				t.Errorf("Explain() = %v, want %v", got, tt.want)
			}
		})
	}
}
