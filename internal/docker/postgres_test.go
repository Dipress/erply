package docker

import (
	"testing"

	"github.com/ory/dockertest"
)

const exec = `
CREATE TABLE example (
	uuid	CHAR (36) PRIMARY KEY
);
`

func Test_NewPostgres(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	t.Log("with dockertest pool")
	{
		pool, err := dockertest.NewPool("")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest: 0\t should start postgres image and return valid connection and resource to purge")
		{
			pd, err := NewPostgres(pool)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if _, err := pd.DB.Exec(exec); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if err := pool.Purge(pd.Resource); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}
