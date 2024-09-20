package randomanimals

import "testing"

func TestRandomAnimal(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "1 test",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := RandomAnimal(); len(got) < tt.want {
				t.Errorf("RandomAnimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
