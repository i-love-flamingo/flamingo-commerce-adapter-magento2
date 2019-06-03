package productSearch

import (
	"bufio"
	"encoding/csv"
	"flamingo.me/flamingo-commerce-adapter-standalone/productSearch/infrastructure/productSearch"
	domain2 "flamingo.me/flamingo-commerce/v3/price/domain"
	"flamingo.me/flamingo-commerce/v3/product/domain"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type (
	// CsvRowDto is a simple Data Transfer Object to receive CSV Rows
	CsvRowDto map[string]string
	CsvLoader struct {
		logger  flamingo.Logger
		csvFile string
	}
)

func (l *CsvLoader) Inject(logger flamingo.Logger, config *struct {
	CsvFile string `inject:"config:flamingo-commerce-adapter-magento2.product.csvPath"`
}) {
	l.logger = logger.WithField(flamingo.LogKeyModule, "magento2").WithField(flamingo.LogKeyCategory, "CsvLoader")
	if config == nil {
		return
	}
	l.csvFile = config.CsvFile
}

func (l *CsvLoader) Load(indexer productSearch.Index) error {
	l.logger.Info(fmt.Sprintf("start loading magento products frome export file: %v", l.csvFile))
	f, err := os.Open(l.csvFile)
	if err != nil {
		l.logger.Error("Error - ProductCsvRowDto %v", err)
		return err
	}

	var headerRow []string

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.LazyQuotes = true
	r.Comma = ','
	rowCount := 0
	isFirstRow := true

	for {
		rowCount++
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}

		if isFirstRow {
			isFirstRow = false
			headerRow = record
			continue
		}

		row := make(map[string]string)

		for k, colName := range headerRow {
			if len(record) > k {
				row[colName] = record[k]
			}
		}
		if row["product_type"] == "simple" {
			err = indexer.Add(l.BuildSimpleFromRow(row))
			if err != nil {
				l.logger.Warn(err)
			}
		}
	}
	l.logger.Info(fmt.Sprintf("indexed: #%v products", rowCount))

	l.logger.Info(fmt.Sprintf("finished loading magento products frome export file: %v", l.csvFile))
	return nil
}

func (l *CsvLoader) BuildSimpleFromRow(row map[string]string) domain.BasicProduct {
	price, _ := strconv.ParseFloat(row["price"], 64)
	var productCategoryTeaser []domain.CategoryTeaser
	categoryPaths := strings.Split(row["categories"], ",")
	for _, categoryPathString := range categoryPaths {
		c := l.getCategory(categoryPathString, nil)
		productCategoryTeaser = append(productCategoryTeaser, *c)
	}
	var mainCategory domain.CategoryTeaser
	if len(productCategoryTeaser) > 0 {
		mainCategory = productCategoryTeaser[0]
	}
	var medias []domain.Media
	if row["base_image"] != "" {
		medias = append(medias,
			domain.Media{
				Reference: row["base_image"],
				Type:      "magento-product-image",
				Usage: "list",
			})
		medias = append(medias,
			domain.Media{
				Reference: row["base_image"],
				Type:      "magento-product-image",
				Usage: "detail",
			})
	}
	return domain.SimpleProduct{
		Identifier: row["sku"],
		BasicProductData: domain.BasicProductData{
			Title:            row["name"],
			MarketPlaceCode:  row["sku"],
			Description:      row["description"],
			ShortDescription: row["short_description"],
			Categories:       productCategoryTeaser,
			MainCategory:     mainCategory,
			Media:medias,
		},
		Saleable: domain.Saleable{
			ActivePrice: domain.PriceInfo{Default: domain2.NewFromFloat(price, "€")},
		},
	}
}

func (l *CsvLoader) getCategory(categorypath string, parent *domain.CategoryTeaser) *domain.CategoryTeaser {
	if categorypath == "" {
		return parent
	}
	categoryNames := strings.Split(categorypath, "/")

	path := categoryNames[0]
	if parent != nil {
		path = parent.Path + "/" + categoryNames[0]
	}
	code := strings.ReplaceAll(path, "+", "-")
	code = strings.ReplaceAll(code, " ", "-")
	code = strings.ReplaceAll(code, "/", "_")
	code = url.QueryEscape(code)
	productCategoryTeaser := domain.CategoryTeaser{
		Path:   path,
		Code:   code,
		Name:   categoryNames[0],
		Parent: parent,
	}
	return l.getCategory(strings.Join(categoryNames[1:], "/"), &productCategoryTeaser)
}
