package repository

func AddDiscipline(name string, dtype string, resonance string, threat string, description string) map[string]interface{} {
	conn := ConnectDuckDB()
	columns := []string{"name", "type", "resonance", "threat", "description"}
	values := []interface{}{name, dtype, resonance, threat, description}
	lastInsertedId := InsertIntoTable(conn, "DISCIPLINES", columns, values)
	insertedDiscipline := SelectById(conn, "DISCIPLINES", lastInsertedId)
	conn.Close()
	return insertedDiscipline
}
