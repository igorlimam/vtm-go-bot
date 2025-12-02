package repository

import (
	"database/sql"
	"log"

	_ "github.com/duckdb/duckdb-go/v2"
)

func ConnectDuckDB() *sql.DB {
	db, err := sql.Open("duckdb", "vtm.duckdb")
	if err != nil {
		log.Fatal("Failed to connect to DuckDB:", err)
	}

	log.Println("Connected to DuckDB successfully.")
	return db
}

func CheckDDL(db *sql.DB) {
	ddlQueries := []string{
		`CREATE TABLE IF NOT EXISTS BLOOD_POTENCIES (
			id INTEGER PRIMARY KEY,
			level INTEGER,
			blood_surge INTEGER,
			mend_amount INTEGER,
			feeding_penalty INTEGER,
			rouse_reroll INTEGER,
			bane_severity INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS PREDATORS (
			id INTEGER PRIMARY KEY,
			description TEXT,
			resonance VARCHAR,
			bonuses TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS CLANS (
			id INTEGER PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			titles TEXT,
			bane TEXT,
			compulsion TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS PLAYERS (
			id INTEGER PRIMARY KEY,
			name VARCHAR,
			chronicle VARCHAR,
			generation INTEGER,
			concept VARCHAR,
			ambition VARCHAR,
			desire VARCHAR,
			title VARCHAR,
			xp INTEGER DEFAULT 0,
			blood_potency_id INTEGER REFERENCES BLOOD_POTENCIES(id),
			clan_id INTEGER REFERENCES CLANS(id),
			predator_id INTEGER REFERENCES PREDATORS(id)
		)`,
		`CREATE TABLE IF NOT EXISTS DISCIPLINES (
			id INTEGER PRIMARY KEY,
			type VARCHAR,
			name VARCHAR,
			description TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS CLANS_DISCIPLINES (
			clan_id INTEGER REFERENCES CLANS(id),
			discipline_id INTEGER REFERENCES DISCIPLINES(id),
			PRIMARY KEY (clan_id, discipline_id)
		)`,
		`CREATE TABLE IF NOT EXISTS POWERS(
			id INTEGER PRIMARY KEY,
			discipline_id INTEGER REFERENCES DISCIPLINES(id),
			name VARCHAR,
			description TEXT,
			dice_pool TEXT,
			cost TEXT,
			duration TEXT,
			system TEXT,
			kind VARCHAR,
			level INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS RELATIONSHIPS (
			player_id INTEGER REFERENCES PLAYERS(id),
			target_id INTEGER REFERENCES PLAYERS(id),
			label VARCHAR,
			PRIMARY KEY (player_id, target_id)
		)`,
		`CREATE TABLE IF NOT EXISTS NOTES (
			id INTEGER PRIMARY KEY,
			player_id INTEGER REFERENCES PLAYERS(id),
			date DATE DEFAULT CURRENT_DATE,
			text TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS STORY_NOTES (
			id INTEGER PRIMARY KEY,
			chronicle VARCHAR,
			date DATE DEFAULT CURRENT_DATE,
			text TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS MECHANICS (
			id INTEGER PRIMARY KEY,
			subject VARCHAR,
			kind VARCHAR,
			info TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS MERITS (
			id INTEGER PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			kind VARCHAR
		)`,
		`CREATE TABLE IF NOT EXISTS MERITS_LEVELS (
			id INTEGER PRIMARY KEY,
			merit_id INTEGER REFERENCES MERITS(id),
			level INTEGER,
			description TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS ATTRIBUTES (
			id INTEGER PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			kind VARCHAR
		)`,
		`CREATE TABLE IF NOT EXISTS ATTRIBUTES_LEVELS (
			id INTEGER PRIMARY KEY,
			attribute_id INTEGER REFERENCES ATTRIBUTES(id),
			level INTEGER,
			description TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS SKILLS (
			id INTEGER PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			kind VARCHAR
		)`,
		`CREATE TABLE IF NOT EXISTS SKILLS_LEVELS (
			id INTEGER PRIMARY KEY,
			skill_id INTEGER REFERENCES SKILLS(id),
			level INTEGER,
			description TEXT
		)`,
	}

	for _, query := range ddlQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Failed to execute DDL query: %v", err)
		}
	}

}

func InsertIntoTable(conn *sql.DB, tableName string, columns []string, values []interface{}) int64 {
	query := "INSERT INTO " + tableName + " ("
	for i, col := range columns {
		query += col
		if i < len(columns)-1 {
			query += ", "
		}
	}
	query += ") VALUES ("
	for i := range values {
		query += "?"
		if i < len(values)-1 {
			query += ", "
		}
	}
	query += ")"

	stmt, _ := conn.Prepare(query)
	sqlResult, err := stmt.Exec(values...)

	if err != nil {
		log.Fatalf("Failed to insert into %s: %v", tableName, err)
	}

	returnedID, err := sqlResult.LastInsertId()
	if err != nil {
		log.Fatalf("Failed to retrieve last insert ID: %v", err)
	}

	log.Printf("Inserted into %s with ID: %d", tableName, returnedID)
	return returnedID
}
