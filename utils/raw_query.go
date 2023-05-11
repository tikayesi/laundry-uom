package utils

const (
	INSERT_UOM      = "INSERT INTO uom (id, name, is_deleted) values (:id, :name, :is_deleted)"
	SELECT_UOM_LIST = "SELECT * FROM uom"
	SELECT_UOM_ID   = "SELECT * FROM uom where id = $1"
	UPDATE_UOM      = "UPDATE uom set name=$1 where id=$2"
	DELETE_UOM      = "DELETE uom where id = $1"

	INSERT_PRODUCT      = "INSERT INTO product (name, price, uom_id) values ($1, $2, $3)"
	SELECT_PRODUCT_LIST = "SELECT * FROM product"
	SELECT_PRODUCT_ID   = "SELECT * FROM product where id = $1"
	UPDATE_PRODUCT      = "UPDATE product set name=$1, price=$2, uom_id=$3 where id=$4"
	DELETE_PRODUCT      = "DELETE product where id = $1"
)
