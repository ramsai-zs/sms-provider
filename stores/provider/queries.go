package provider

const (
	create         = `INSERT INTO provider(id,url,channelRefID,name) VALUES($1,$2,$3,$4)`
	getByID        = `SELECT id,url,channelRefID,name FROM provider`
	deleteProvider = `DELETE FROM provider WHERE id = $1`
	update         = `UPDATE INTO provider SET url=$1,name=$2`
)
