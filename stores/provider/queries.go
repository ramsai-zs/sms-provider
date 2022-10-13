package provider

const (
	INSERT  = `INSERT INTO Provider(ID, URL, ChannelRefID) VALUES($1, $2, $3)`
	GETBYID = `SELECT ID,URL FROM Provider WHERE ID=$1`
	DELETE  = `DELETE FROM Provider WHERE ID=$1`
	UPDATE  = `UPDATE Provider SET ID=$1, URL=$2 WHERE ChannelRefID=$3`
)
