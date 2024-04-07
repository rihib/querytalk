package main_test

import (
	"context"
	"fmt"
	"testing"

	ogen "querychat/ogen"
)

// TODO: Create comprehensive tests for chat API

const (
	BASEURL  = "http://localhost:8080"
	DB_TYPE  = "SQLite3"
	PROMPT   = "What are the monthly sales for 2013?"
	EXPECTED = `{` +
		`"chart":{"type":"line","x":"SaleMonth","y":"TotalSales"},` +
		`"data":[` +
		`{"SaleMonth":"01","TotalSales":37.62},` +
		`{"SaleMonth":"02","TotalSales":27.72},` +
		`{"SaleMonth":"03","TotalSales":37.62},` +
		`{"SaleMonth":"04","TotalSales":33.66},` +
		`{"SaleMonth":"05","TotalSales":37.62},` +
		`{"SaleMonth":"06","TotalSales":37.62},` +
		`{"SaleMonth":"07","TotalSales":37.62},` +
		`{"SaleMonth":"08","TotalSales":37.62},` +
		`{"SaleMonth":"09","TotalSales":37.62},` +
		`{"SaleMonth":"10","TotalSales":37.62},` +
		`{"SaleMonth":"11","TotalSales":49.62},` +
		`{"SaleMonth":"12","TotalSales":38.62}` +
		`]` +
		`}`
	EXPECTED2 = `{` +
		`"chart":{"type":"line","x":"SaleMonth","y":"TotalSales"},` +
		`"data":[` +
		`{"SaleMonth":"2013-01","TotalSales":37.62},` +
		`{"SaleMonth":"2013-02","TotalSales":27.72},` +
		`{"SaleMonth":"2013-03","TotalSales":37.62},` +
		`{"SaleMonth":"2013-04","TotalSales":33.66},` +
		`{"SaleMonth":"2013-05","TotalSales":37.62},` +
		`{"SaleMonth":"2013-06","TotalSales":37.62},` +
		`{"SaleMonth":"2013-07","TotalSales":37.62},` +
		`{"SaleMonth":"2013-08","TotalSales":37.62},` +
		`{"SaleMonth":"2013-09","TotalSales":37.62},` +
		`{"SaleMonth":"2013-10","TotalSales":37.62},` +
		`{"SaleMonth":"2013-11","TotalSales":49.62},` +
		`{"SaleMonth":"2013-12","TotalSales":38.62}` +
		`]` +
		`}`
)

func run(ctx context.Context, dbType string, prompt string) (*ogen.VisualizableData, error) {
	c, err := ogen.NewClient(BASEURL)
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	request := ogen.OptMSG{
		Value: ogen.MSG{
			DbType: ogen.OptString{Value: dbType, Set: true},
			Prompt: ogen.OptString{Value: prompt, Set: true},
		},
		Set: true,
	}
	r, err := c.SendMSG(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("send prompt: %w", err)
	}

	return r, nil
}

func TestChatAPI(t *testing.T) {
	r, err := run(context.Background(), DB_TYPE, PROMPT)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if r.VisualizableData != EXPECTED && r.VisualizableData != EXPECTED2 {
		t.Errorf("test failed: got `%v`\n\nexpected `%v`", r.VisualizableData, EXPECTED)
	}
}
