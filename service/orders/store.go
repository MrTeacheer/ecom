package orders

import (
	"database/sql"

	"github.com/MrTeacheer/ecom/types"
)



type Store struct{
	db *sql.DB
}


func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}


func (s *Store)CreateOrder(payload types.OrdersAdd) error{
	_,err := s.db.Exec("INSERT INTO orders (user_id,total,status,address) VALUES (?,?,?,?)",
						payload.UserId,
					    payload.Total,
						payload.Status,
						payload.Address)

	if err != nil{
		return err
	}
	return nil
}


func (s *Store) GetOrders() ([]types.Orders,error){
	rows,err := s.db.Query("SELECT * FROM orders")
	if err != nil{
		return nil,err
	}
	var ods []types.Orders
	for rows.Next(){
		var o = new(types.Orders)
		o,err = scanIntoOrders(rows)
		if err != nil{
			return nil,err
		}
		ods = append(ods, *o)

	}
	return ods,nil
}


func scanIntoOrders(rows *sql.Rows) (*types.Orders,error){
	o := new(types.Orders)

	err := rows.Scan(
		&o.Id,
		&o.UserId,
		&o.Total,
		&o.Status,
		&o.Address,
		&o.CreatedAt,
	)
	if err != nil{
		return nil,err
	}
	return o,nil
}