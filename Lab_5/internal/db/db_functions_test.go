package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

var testGN = []struct {
	names       []string
	errExpected error
}{
	{
		names:       []string{"Alice", "Bob"},
		errExpected: nil,
	},
	{
		names:       []string{},
		errExpected: errors.New("no rows in result set"),
	},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	dbService := Service{DB: mockDB}

	for i, row := range testGN {
		rows := sqlmock.NewRows([]string{"name"})

		for _, name := range row.names {
			rows.AddRow(name)
		}

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows).WillReturnError(row.errExpected)

		names, err := dbService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %v, actual error: %v", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
		} else {
			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.names, names, "row: %d, expected names: %v, actual names: %v", i, row.names, names)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

var testSUV = []struct {
	name       string
	columnName string
	tableName  string
	want       []string
	wantErr    bool
	err        error
	rows       *sqlmock.Rows
}{
	{
		name:       "Success case",
		columnName: "column",
		tableName:  "table",
		want:       []string{"Value1", "Value2"},
		wantErr:    false,
		err:        nil,
		rows: sqlmock.NewRows([]string{"value"}).
			AddRow("Value1").
			AddRow("Value2"),
	},
	{
		name:       "Error case",
		columnName: "column",
		tableName:  "table",
		want:       nil,
		wantErr:    true,
		err:        sql.ErrNoRows,
		rows:       sqlmock.NewRows([]string{"value"}),
	},
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := New(mockDB)

	for _, tt := range testSUV {
		t.Run(tt.name, func(t *testing.T) {
			query := "SELECT DISTINCT " + tt.columnName + " FROM " + tt.tableName
			if tt.wantErr {
				mock.ExpectQuery(query).WillReturnError(tt.err)
			} else {
				mock.ExpectQuery(query).WillReturnRows(tt.rows)
			}

			got, err := dbService.SelectUniqueValues(tt.columnName, tt.tableName)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
