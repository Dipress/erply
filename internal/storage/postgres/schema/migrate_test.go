package schema

import (
	"testing"
)

func Test_Migrate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	t.Log("with given database connection.")
	{
		m, err := newMigration(db)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould up schema.")
		{
			if err := Migrate(db); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}

		t.Log("\ttest:1\tshould down schema.")
		{
			if err := m.Down(); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}
