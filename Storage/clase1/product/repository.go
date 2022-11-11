package product

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/Storage/clase1/domain"
)

const (
	StoreQuery  = "INSERT INTO products(name, type, count, price) VALUES(?, ?, ?, ?);"
	GetOneQuery = "SELECT id, name, type, count, price FROM products WHERE id = ?;"
	UpdateQuery = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?;"
	GetAllQuery = "SELECT id, name, type, count, price FROM products;"
	DeleteQuery = "DELETE FROM products WHERE id=?;"
	ExistsQuery = "SELECT id FROM products WHERE id=?;"
	GetFullData = "SELECT products.id, products.name, products.type, products.count, products.price, products.id_warehouse FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id WHERE warehouses.id =?;"
)

type Repository interface {
	GetOne(ctx context.Context, id int) (domain.Product, error)
	Store(ctx context.Context, p domain.Product) (int64, error)
	Update(ctx context.Context, p domain.Product, id int) error
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) bool
	GetFullData(ctx context.Context, id int) ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(ExistsQuery, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) GetOne(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GetOneQuery, id)
	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.TypeProduct, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) Store(ctx context.Context, p domain.Product) (int64, error) {
	stm, err := r.db.Prepare(StoreQuery)
	if err != nil {
		return 0, err
	}

	defer stm.Close()

	result, err := stm.Exec(p.Name, p.TypeProduct, p.Count, p.Price)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) Update(ctx context.Context, p domain.Product, id int) error {
	stm, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		return err
	}
	defer stm.Close()
	result, err := stm.Exec(p.Name, p.TypeProduct, p.Count, p.Price, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no se afectó ninguna línea")
	}
	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		return []domain.Product{}, err
	}
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.TypeProduct, &product.Count, &product.Price)
		if err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stm, err := r.db.Prepare(DeleteQuery)
	if err != nil {
		return err
	}
	defer stm.Close()
	result, err := stm.Exec(id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("error: no se afectó correctamente")
	}
	return nil
}

func (r *repository) GetFullData(ctx context.Context, id int) ([]domain.Product, error) {
	rows, err := r.db.Query(GetFullData, id)
	if err != nil {
		return []domain.Product{}, err
	}
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.TypeProduct, &product.Count, &product.Price, &product.WarehouseId); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}
