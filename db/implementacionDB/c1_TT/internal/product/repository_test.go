package product

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/domain"
)

var product_test = domain.Product{
	ID:           1,
	Name:         "Product Test",
	Type:         "Testing",
	Count:        5,
	Price:        10.0,
	Warehouse_id: 1,
}

func TestGetOneWithContext(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Equal(t, err, nil)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price"} // Estas son las columnas que estan definidas en la query que vamos a testear en este caso GET_PRODUCT
	rows := sqlmock.NewRows(columns)
	rows.AddRow(product_test.ID, product_test.Name, product_test.Type,
		product_test.Count, product_test.Price)
	mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT)).WithArgs(product_test.ID).WillReturnRows(rows)

	myRepo := NewRepo(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productResult, err := myRepo.GetOneWithContext(ctx, product_test.ID)
	assert.Equal(t, err, nil)
	assert.Equal(t, product_test.Name, productResult.Name)
	assert.Equal(t, product_test.ID, productResult.ID)
	assert.Equal(t, nil, mock.ExpectationsWereMet())

}
