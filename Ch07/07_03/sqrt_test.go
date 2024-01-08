package sqrt

import (
	"testing"
	"encoding/csv"
	"os"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	   // Open the CSV file
	   file, err := os.Open("./sqrt_cases.csv")
	   if err != nil {
		   panic(err)
	   }
	   defer file.Close()
	val, err := Sqrt(2)
	require.NoError(t, err)
	require.InDelta(t, 1.414214, val, 0.001)
}
