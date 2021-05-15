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
		want string
	}{
		{
			name: "Test",
			args: args{
				sql:  "SELECT * FROM users where email = ?",
				args: []interface{}{"a@a.com"},
			},
			want: "SELECT * FROM users where email = 'a@a.com'",
		},
		{
			name: "Test multiple param",
			args: args{
				sql:  "SELECT * FROM users where email = ? and uuid = ?",
				args: []interface{}{"a@a.com", "uuid"},
			},
			want: "SELECT * FROM users where email = 'a@a.com' and uuid = 'uuid'",
		},
		{
			name: "Select all test",
			args: args{
				sql:  "SELECT * FROM users",
				args: []interface{}{"a@a.com"},
			},
			want: "SELECT * FROM users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Explain(tt.args.sql, tt.args.args...); got != tt.want {
				t.Errorf("Explain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExplainSQL(b *testing.B) {
	// run the Fib function b.N times
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Explain("SELECT * FROM users", []interface{}{})
	}
}

func BenchmarkExplainSQLWithArg(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Explain("SELECT * FROM users where email = ?", "email")
	}
}

func BenchmarkExplainSQLWith2Arg(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Explain("SELECT * FROM users where email = ? and username = ?", "email", "myusername")
	}
}

func BenchmarkExplainSQLWithManyManyArgs(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Explain("SELECT * FROM users where email = ? and username = ? and email_canonical = ? and username_canonical = ? or ID = ?", "email", "myusername", "email", "username", 1)
	}
}
