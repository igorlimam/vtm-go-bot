package repository

func AddDiscipline(name string, dtype string, resonance string, threat string) map[string]interface{} {
	conn := ConnectDuckDB()
	columns := []string{"name", "type", "resonance", "threat"}
	values := []interface{}{name, dtype, resonance, threat}
	lastInsertedId := InsertIntoTable(conn, "DISCIPLINES", columns, values)
	insertedDiscipline := SelectById(conn, "DISCIPLINES", lastInsertedId)
	conn.Close()
	return insertedDiscipline
}
