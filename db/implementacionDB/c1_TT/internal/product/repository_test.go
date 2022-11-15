package product

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/domain"
	"github.com/stretchr/testify/assert"
)

var product_test = domain.Product{
	ID:           1,
	Name:         "Product Test",
	Type:         "Testing",
	Count:        5,
	Price:        10.0,
	Warehouse_id: 1,
}

var (
	ERRORFORZADO = errors.New("Error forzado")
)

func TestGetOneWithContext(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	assert.Equal(t, product_test.Name, productResult.Name)
	assert.Equal(t, product_test.ID, productResult.ID)
	assert.Equal(t, nil, mock.ExpectationsWereMet())

}

func TestExistOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(product_test.ID) // es el argumento que pide en la Query del Exists
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_PRODCUT)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepo(db)
	exists := repo.Exists(context.TODO(), 1)

	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestExistFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(product_test.ID) // es el argumento que pide en la Query del Exists
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_PRODCUT)).WithArgs(2).WillReturnRows(rows)

	repo := NewRepo(db)
	exists := repo.Exists(context.TODO(), 1)

	assert.False(t, exists)

}

func TestSaveOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Save Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "name", "type", "count", "price", "id_warehouse"} // Estas son las columnas que estan definidas en la query que vamos a testear en este caso GET_PRODUCT
		rows := sqlmock.NewRows(columns)

		rows.AddRow(product_test.ID, product_test.Name, product_test.Type,
			product_test.Count, product_test.Price, product_test.Warehouse_id)

		mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT)).WithArgs(1).WillReturnRows(rows)

		repo := NewRepo(db)

		id, err := repo.Store(context.TODO(), product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.Warehouse_id)
		assert.NoError(t, err)

		product, err := repo.GetOne(context.TODO(), int(id))
		assert.NoError(t, err)

		assert.NotNil(t, product)
		assert.Equal(t, product_test.ID, product.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Save Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "name", "type", "count", "price", "id_warehouse"} // Estas son las columnas que estan definidas en la query que vamos a testear en este caso GET_PRODUCT
		rows := sqlmock.NewRows(columns)

		rows.AddRow(product_test.ID, product_test.Name, product_test.Type,
			product_test.Count, product_test.Price, product_test.Warehouse_id)

		mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT)).WithArgs(1).WillReturnRows(rows)

		repo := NewRepo(db)

		id, err := repo.Store(context.TODO(), product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.Warehouse_id)
		assert.NoError(t, err)

		product, err := repo.GetOne(context.TODO(), int(id))
		assert.NoError(t, err)

		assert.NotNil(t, product)
		assert.Equal(t, product_test.ID, product.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Save Fail", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnError(ERRORFORZADO)

		repo := NewRepo(db)

		id, err := repo.Store(context.TODO(), product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.Warehouse_id)

		assert.EqualError(t, err, ERRORFORZADO.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAllOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"} // Estas son las columnas que estan definidas en la query que vamos a testear en este caso GET_PRODUCT
	rows := sqlmock.NewRows(columns)
	products := []domain.Product{{ID: 1, Name: "Product Test", Type: "Testing", Count: 5, Price: 10.0}, {ID: 2, Name: "Product Test 2", Type: "Testing 2", Count: 10, Price: 20.0}}

	for _, p := range products {
		rows.AddRow(p.ID, p.Name, p.Type, p.Count, p.Price, p.Warehouse_id)
	}
	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_PRODUCTS)).WillReturnRows(rows)

	repo := NewRepo(db)

	productsResult, err := repo.GetAll(context.TODO())
	assert.NoError(t, err)

	assert.Equal(t, products, productsResult)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestGetAllFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"} // Estas son las columnas que estan definidas en la query que vamos a testear en este caso GET_PRODUCT
	rows := sqlmock.NewRows(columns)
	products := []domain.Product{{ID: 1, Name: "Product Test", Type: "Testing", Count: 5, Price: 10.0}, {ID: 2, Name: "Product Test 2", Type: "Testing 2", Count: 10, Price: 20.0}}

	for _, p := range products {
		rows.AddRow(p.ID, p.Name, p.Type, p.Count, p.Price, p.Warehouse_id)
	}
	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_PRODUCTS)).WillReturnError(ERRORFORZADO)

	repo := NewRepo(db)

	productsResult, err := repo.GetAll(context.TODO())

	assert.EqualError(t, err, ERRORFORZADO.Error())
	assert.Empty(t, productsResult)
	assert.NoError(t, mock.ExpectationsWereMet())

}
