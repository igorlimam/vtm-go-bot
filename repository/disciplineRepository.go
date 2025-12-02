package repository

func AddDiscipline(name string, dtype string, resonance string, threat string) int64 {
	conn := ConnectDuckDB()
	columns := []string{"name", "type", "resonance", "threat"}
	values := []interface{}{name, dtype, resonance, threat}
	lastInsertedId := InsertIntoTable(conn, "DISCIPLINES", columns, values)
	return lastInsertedId
}
