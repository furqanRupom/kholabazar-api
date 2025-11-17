package repo

import (
	"database/sql"
	"kholabazar/domain"
	"kholabazar/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.Service
}

type productRepo struct {
	db *sqlx.DB
}

/* Constructor */
func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{
		db: db,
	}

}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	INSERT INTO products(name,image,price,description,category)
	VALUES(
	$1, $2, $3, $4, $5
	)
	RETURNING id
	`
	row := r.db.QueryRow(query, p.Name, p.Image, p.Price, p.Description, p.Category)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	var productList []*domain.Product
	offset := ((page - 1) * limit) + 1
	query := `
	     SELECT 
	         id,
	         name,
	         description,
	         price,
	         image,
	         category
	      FROM 
	          products
	       LIMIT $1
	       OFFSET $2
	`
	err := r.db.Select(&productList, query, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return productList, nil

}
func (r *productRepo) Count() (int64,error){
	query := `
	SELECT COUNT(*) FROM products
	`
	var count int64
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count,nil
}
func (r *productRepo) Get(ID int) (*domain.Product, error) {
	var product domain.Product
	query := `
	SELECT id,name,price,description,image,category FROM products WHERE id=$1
	`
	err := r.db.Get(&product, query, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `UPDATE products 
	SET 
	name=$1, 
	image=$2, 
	price=$3, 
	description=$4,
	category=$5
	WHERE 
	id=$6
	`
	row := r.db.QueryRow(query, p.Name, p.Image, p.Price, p.Description, p.Category, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *productRepo) Delete(ID int) error {
	query := `
	DELETE FROM products WHERE id=$1
	`
	_, err := r.db.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil

}
