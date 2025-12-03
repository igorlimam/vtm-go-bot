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
	CheckDDL(db)
	return db
}

func CheckDDL(db *sql.DB) {
	ddlQueries := []string{
		`CREATE SEQUENCE IF NOT EXISTS discipline_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS blood_potency_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS clan_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS player_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS predator_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS power_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS note_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS story_note_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS mechanic_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS merit_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS attribute_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS skill_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS merit_level_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS attribute_level_id_seq;`,
		`CREATE SEQUENCE IF NOT EXISTS skill_level_id_seq;`,
		`CREATE TABLE IF NOT EXISTS BLOOD_POTENCIES (
			id INTEGER DEFAULT nextval('blood_potency_id_seq') PRIMARY KEY,
			level INTEGER,
			blood_surge INTEGER,
			mend_amount INTEGER,
			feeding_penalty INTEGER,
			rouse_reroll INTEGER,
			bane_severity INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS PREDATORS (
			id INTEGER DEFAULT nextval('predator_id_seq') PRIMARY KEY,
			description TEXT,
			resonance VARCHAR,
			bonuses TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS CLANS (
			id INTEGER DEFAULT nextval('clan_id_seq') PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			titles TEXT,
			bane TEXT,
			compulsion TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS PLAYERS (
			id INTEGER DEFAULT nextval('player_id_seq') PRIMARY KEY,
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
			id INTEGER DEFAULT nextval('discipline_id_seq') PRIMARY KEY,
			type VARCHAR,
			name VARCHAR,
			threat VARCHAR,
			resonance VARCHAR,
			description TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS CLANS_DISCIPLINES (
			clan_id INTEGER REFERENCES CLANS(id),
			discipline_id INTEGER REFERENCES DISCIPLINES(id),
			PRIMARY KEY (clan_id, discipline_id)
		)`,
		`CREATE TABLE IF NOT EXISTS POWERS(
			id INTEGER DEFAULT nextval('power_id_seq') PRIMARY KEY,
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
			id INTEGER DEFAULT nextval('note_id_seq') PRIMARY KEY,
			player_id INTEGER REFERENCES PLAYERS(id),
			date DATE DEFAULT CURRENT_DATE,
			text TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS STORY_NOTES (
			id INTEGER DEFAULT nextval('story_note_id_seq') PRIMARY KEY,
			chronicle VARCHAR,
			date DATE DEFAULT CURRENT_DATE,
			text TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS MECHANICS (
			id INTEGER DEFAULT nextval('mechanic_id_seq') PRIMARY KEY,
			subject VARCHAR,
			kind VARCHAR,
			info TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS MERITS (
			id INTEGER DEFAULT nextval('merit_id_seq') PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			kind VARCHAR
		)`,
		`CREATE TABLE IF NOT EXISTS MERITS_LEVELS (
			id INTEGER DEFAULT nextval('merit_level_id_seq') PRIMARY KEY,
			merit_id INTEGER REFERENCES MERITS(id),
			level INTEGER,
			description TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS ATTRIBUTES (
			id INTEGER DEFAULT nextval('attribute_id_seq') PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			kind VARCHAR
		)`,
		`CREATE TABLE IF NOT EXISTS ATTRIBUTES_LEVELS (
			id INTEGER DEFAULT nextval('attribute_level_id_seq') PRIMARY KEY,
			attribute_id INTEGER REFERENCES ATTRIBUTES(id),
			level INTEGER,
			description TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS SKILLS (
			id INTEGER DEFAULT nextval('skill_id_seq') PRIMARY KEY,
			name VARCHAR,
			description TEXT,
			kind VARCHAR
		)`,
		`CREATE TABLE IF NOT EXISTS SKILLS_LEVELS (
			id INTEGER DEFAULT nextval('skill_level_id_seq') PRIMARY KEY,
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
	query += ") RETURNING id;"
	var insertedId int64
	err := conn.QueryRow(query, values...).Scan(&insertedId)

	if err != nil {
		log.Fatalf("Failed to query insert statement for %s: %v", tableName, err)
	}

	log.Printf("Inserted into %s with ID: %d", tableName, insertedId)
	return insertedId
}

func SelectById(conn *sql.DB, tableName string, id int64) map[string]interface{} {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"

	rows, _ := conn.Query(query, id)
	defer rows.Close()
	log.Printf("Executing query: %s with ID: %d with rows %+v", query, id, rows)
	results, _ := convertRowsToMap(rows)
	if len(results) > 0 {
		return results[0]
	}
	return nil
}

func convertRowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	log.Printf("Converting rows to map with columns: %v", columns)
	results := []map[string]interface{}{}

	for rows.Next() {
		columnPointers := make([]interface{}, len(columns))
		columnValues := make([]interface{}, len(columns))
		for i := range columnPointers {
			columnPointers[i] = &columnValues[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnValues[i]
			b, ok := val.([]byte)
			if ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}
	return results, nil
}
