package products

import (
	"database/sql"

	"github.com/MrTeacheer/ecom/types"
)



type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}


func (s *Store) CreateProduct(product types.ProductAdd) error {
	_, err := s.db.Exec("INSERT INTO products (name, description,image,price,quantity) VALUES (?, ?,?, ?, ?)",
	 product.Name,
	 product.Description,
	 product.Image,
	 product.Price,
	 product.Quantity)

	if err != nil {
		return err
	}

	return nil
}



func (s *Store) GetProducts() ([]types.Product,error){
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil{
		return nil,err
	}
	var products []types.Product

	for rows.Next(){
		var p = new(types.Product)
		p,err = scanIntoProducts(rows)
		if err != nil{
			return nil,err
		}
		products = append(products, *p)
	}
	return products,nil
}


func scanIntoProducts(rows *sql.Rows) (*types.Product,error){
	product := new(types.Product)

	err := rows.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err!=nil{
		return nil,err
	}

	return product,nil
}