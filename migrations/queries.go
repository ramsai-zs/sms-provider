package migrations

const (
	createTable = `CREATE TABLE IF NOT EXISTS provider(
      ID VARCHAR(40) PRIMARY KEY,
      URL VARCHAR(50), 
       channelRefID VARCHAR(40) NOT NULL,
       name VARCHAR(80))`
	dropTable = `DROP TABLE IF EXISTS provider`
)
