package list

import "testing"

func TestToSlice(t *testing.T) {
	list1 := List{5, &List{6, &List{1, &List{9, nil}}}}
	str1 := "5-->6-->1-->9"
	str2 := "0"
	var list2 List
	tests := []struct {
		name string
		give List
		want string
	}{
		{
			name: "non_empty_list",
			give: list1,
			want: str1,
		},
		{
			name: "empty_list",
			give: list2,
			want: str2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.ToString()
			if got != tt.want {
				const msg = `ToString : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}
