package sqlite

const (
	selectAvgRateQuery = `SELECT AVG(rate) as rate, COUNT(*) as voters FROM rate_records WHERE manga_id = ?`
)
