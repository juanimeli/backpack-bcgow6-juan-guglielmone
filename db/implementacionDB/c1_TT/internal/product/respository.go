package product

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/domain"
)

const (
	SAVE_PRODUCT = "INSERT INTO products (name, type, count, price, id_warehouse) VALUES (?,?,?,?,?)"

	GET_PRODUCT = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE id=?;"

	GET_ALL_PRODUCTS = "SELECT id, name, type, count, price, id_warehouse FROM products;"

	UPDATE_PRODUCT = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?;"

	DELETE_PRODUCT = "DELETE FROM products WHERE id = ?;"

	EXIST_PRODCUT = "SELECT p.id FROM products p WHERE m.id=?"

	GET_PRODUCT_FULL_DATA = "SELECT p.id, p.name, p.type, p.count, p.price, w.name, w.adress " +
		"FROM products p INNER JOIN warehouses w ON p.id_warehouse = w.id " +
		"WHERE p.id = ?;"

	GET_PRODUCT_TIMEOUT = "SELECT SLEEP(20) FROM DUAL WHERE id=?;"
)

type Respository interface {
	Store(ctx context.Context, name, productType string, count int, price float64, id_warehouse int) (int, error)
	GetOne(ctx context.Context, id int) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id int) error
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) bool
	GetFullData(ctx context.Context, id int) (domain.Product, error)
	GetOneWithContext(ctx context.Context, id int) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Respository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_PRODCUT, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) Store(ctx context.Context, name, productType string, count int, price float64, id_warehouse int) (int, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT)
	if err != nil {
		return 0, err
	}
	defer stm.Close()

	result, err := stm.Exec(name, productType, count, price, id_warehouse)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
func (r *repository) GetOne(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT, id)
	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse_id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
func (r *repository) Update(ctx context.Context, product domain.Product, id int) error {
	stm, err := r.db.Prepare(UPDATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close()

	result, err := stm.Exec(product.Name, product.Type, product.Count, product.Price, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("error: no affected rows")
	}

	return nil

}
func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GET_ALL_PRODUCTS)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse_id)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DELETE_PRODUCT)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetFullData(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_FULL_DATA, id)
	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse, &product.WarehouseAdress)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r repository) GetOneWithContext(ctx context.Context, id int) (domain.Product, error) {

	var product domain.Product

	row := r.db.QueryRowContext(ctx, GET_PRODUCT, id) //ORIGINAL DE LA FUNCION
	//row := r.db.QueryRowContext(ctx, GET_PRODUCT_TIMEOUT, id) // PARA PROBAR QUE SE PASE DE TIEMPO
	err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		log.Fatal(err)
		return product, err
	}

	return product, nil
}
