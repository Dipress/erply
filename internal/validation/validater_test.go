package validation

import (
	"context"
	"reflect"
	"testing"

	"github.com/romanyx/erply/internal/create"
)

func TestCreateValidate(t *testing.T) {
	tt := []struct {
		name    string
		f       create.Form
		wantErr bool
		expect  Errors
	}{
		{
			name: "valid",
			f: create.Form{
				Title: "title",
				Body:  "body",
			},
		},
		{
			name: "missing title",
			f: create.Form{
				Body: "body",
			},
			wantErr: true,
			expect: Errors{
				"title": "cannot be blank",
			},
		},
		{
			name: "blank title",
			f: create.Form{
				Title: "",
				Body:  "body",
			},
			wantErr: true,
			expect: Errors{
				"title": "cannot be blank",
			},
		},
		{
			name: "long title",
			f: create.Form{
				Title: "This is long title, this title is way larger is allowed one, an ti's used for testing.",
				Body:  "body",
			},
			wantErr: true,
			expect: Errors{
				"title": "the length must be between 1 and 50",
			},
		},
		{
			name: "missing body",
			f: create.Form{
				Title: "title",
			},
			wantErr: true,
			expect: Errors{
				"body": "cannot be blank",
			},
		},
		{
			name: "blank body",
			f: create.Form{
				Title: "title",
				Body:  "",
			},
			wantErr: true,
			expect: Errors{
				"body": "cannot be blank",
			},
		},
	}

	v := Create{}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err := v.Validate(ctx, &tc.f)
			if tc.wantErr {
				got, ok := err.(Errors)
				if !ok {
					t.Errorf("unknown error: %v", err)
					return
				}

				if !reflect.DeepEqual(tc.expect, got) {
					t.Errorf("expected: %+#v got: %+#v", tc.expect, got)
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
