package migrations

const (
	createTableProvider = `CREATE TABLE IF NOT EXISTS Provider(
      id VARCHAR(40) PRIMARY KEY,
      url VARCHAR(50) NOT NULL,
	  channelRefID VARCHAR(40) NOT NULL,
      name VARCHAR(80) NOT NULL)`

	dropTableProvider = `DROP TABLE IF EXISTS Provider`

	createStatusEnum = `DO $$ BEGIN
               IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
               CREATE TYPE status AS ENUM ('YES','NO');
               END IF;
               END$$;;`

	dropStatusEnum = `DROP TYPE IF EXISTS status CASCADE`

	createTableMessage = `CREATE TABLE IF NOT EXISTS Message(
      id VARCHAR(40) PRIMARY KEY,
      message VARCHAR(100) NOT NULL, 
	  number VARCHAR(20) NOT NULL,
      transactional VARCHAR(50) NOT NULL,
	  status status NOT NULL DEFAULT 'NO',
      delivered_time TIMESTAMP)`

	dropTableMessage = `DROP TABLE IF EXISTS Message`
)
