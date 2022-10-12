package migrations

const (
	createTableProvider = `CREATE TABLE IF NOT EXISTS Provider(
      ID VARCHAR(40) PRIMARY KEY,
      URL VARCHAR(50) NOT NULL,
	  ChannelRefID VARCHAR(40) NOT NULL,
      Name VARCHAR(80) NOT NULL)`

	dropTableProvider = `DROP TABLE IF EXISTS Provider`

	createStatusEnum = `DO $$ BEGIN
               IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
               CREATE TYPE status AS ENUM ('YES','NO');
               END IF;
               END$$;;`

	dropStatusEnum = `DROP TYPE IF EXISTS status CASCADE`

	createTableMessage = `CREATE TABLE IF NOT EXISTS Message(
      ID VARCHAR(40) PRIMARY KEY,
      Message VARCHAR(100) NOT NULL, 
	  Number VARCHAR(20) NOT NULL,
      Transactional VARCHAR(50) NOT NULL,
	  Status status NOT NULL DEFAULT 'NO',
      Delivered_time TIMESTAMP)`

	dropTableMessage = `DROP TABLE IF EXISTS Message`
)
