package jsonschema_test

//
// func TestID_Fragments(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		s         jsonschema.Anchor
// 		wantQuery []string
// 	}{
// 		{
// 			name: "Empty string",
// 			s : jsonschema.NewAnchor(""),
// 			wantQuery:[]string{},
// 		},
// 		{
// 			name: "No pointer",
// 			s : jsonschema.NewAnchor("test"),
// 			wantQuery:[]string{},
// 		},
// 		{
// 			name: "Just pointer",
// 			s : jsonschema.NewAnchor("#"),
// 			wantQuery:[]string{},
// 		},
// 		{
// 			name: "One fragment",
// 			s : jsonschema.NewAnchor("#test"),
// 			wantQuery:[]string{"test"},
// 		},
// 		{
// 			name: "Two fragment",
// 			s : jsonschema.NewAnchor("#hello/world"),
// 			wantQuery:[]string{"hello", "world"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotQuery := tt.s.Pointer(); !reflect.DeepEqual(gotQuery, tt.wantQuery) {
// 				t.Errorf("Pointer() = %v, want %v", gotQuery, tt.wantQuery)
// 			}
// 		})
// 	}
// }
