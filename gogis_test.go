package gogis_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	dbQueryString = "SELECT coordinate FROM geometry"

	dbExecString = "INSERT INTO geometry(coordinate) VALUES ($1)"
)

func matcher(expectedSQL, actualSQL string) error {
	space := regexp.MustCompile(`\s+`)

	expected := strings.Trim(space.ReplaceAllString(expectedSQL, " "), "; \t\n")
	actual := strings.Trim(space.ReplaceAllString(actualSQL, " "), "; \t\n")

	if expected != actual {
		return fmt.Errorf("\n** EXPECTED ** %s\n** ACTUAL   ** %s", expected, actual)
	}

	return nil
}

type testFixtureScan struct {
	title            string
	rawData          []byte
	expectedGeometry sql.Scanner
	scanner          sql.Scanner
}

func scanTest(t *testing.T, fixture testFixtureScan) {
	t.Helper()

	t.Run(fmt.Sprintf("scan-%s", fixture.title), func(t *testing.T) {

		dbSQL, mock, err := sqlmock.New(
			sqlmock.QueryMatcherOption(sqlmock.QueryMatcherFunc(matcher)),
		)
		require.NoError(t, err)

		mock.ExpectQuery(dbQueryString).WillReturnRows(
			sqlmock.NewRows([]string{"coordinate"}).
				AddRow(fixture.rawData))

		rows, err := dbSQL.Query(dbQueryString)
		require.NoError(t, err)
		require.NoError(t, rows.Err())

		defer func() {
			_ = rows.Close()
		}()

		if rows.Next() {
			require.NoError(t, rows.Scan(fixture.scanner))
		}

		assert.Equal(t, fixture.expectedGeometry, fixture.scanner)
	})
}

type testFixtureValue struct {
	title           string
	expectedRawData []byte
	valuer          driver.Valuer
}

func valueTest(t *testing.T, fixture testFixtureValue) {
	t.Helper()

	t.Run(fmt.Sprintf("value-%s", fixture.title), func(t *testing.T) {
		out, err := fixture.valuer.Value()
		require.NoError(t, err)

		dataByte, ok := out.([]byte)
		require.True(t, ok, "should be []byte output")

		assert.Equal(t, strings.ToUpper(string(fixture.expectedRawData)), strings.ToUpper(string(dataByte)))
	})
}
