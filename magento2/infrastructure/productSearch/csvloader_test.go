package productSearch

import (
	"github.com/stretchr/testify/assert"

	"testing"

)

func TestCsvLoader_BuildSimpleFromRow(t *testing.T) {

		t.Run("simple", func(t *testing.T) {
			l := &CsvLoader{}
			row := make(map[string]string)
			row["categories"] = "Default Category/Gear,Default Category/Gear/Bags"
			row["price"] = "34.000"
			row["name"] = "Name"

			simple := l.BuildSimpleFromRow(row)

			assert.Equal(t,"Name",simple.BaseData().Title)
			assert.Equal(t,"Gear",simple.BaseData().MainCategory.Name)
			assert.Equal(t,"Default Category",simple.BaseData().MainCategory.Parent.Name)
		})

}
