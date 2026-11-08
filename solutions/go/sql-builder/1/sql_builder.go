package sqlbuilder

// import "fmt"
import "strings"

type ColumnValue struct {
	Column string
	Value  string
}

func BuildSelectQuery(table string, columns []string) string {
	var sb strings.Builder
    sb.WriteString("SELECT ")
    if len(columns) == 0 {
        sb.WriteString("*")
    }
    sb.WriteString(strings.Join(columns, ", "))
    sb.WriteString(" FROM ")
    sb.WriteString(table)
    sb.WriteString(";")
    return sb.String()
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	if len(rows) == 0 {
        return make([]string, 0)
    }
    result := make([]string, 0, len(rows))
    var sb strings.Builder
    for _, row := range rows {
        if len(row) == 0 {
            continue
        }
        sb.WriteString("INSERT INTO ")
        sb.WriteString(table)
        sb.WriteString(" (")
        for i, col := range row {
            if i != 0 {
                sb.WriteString(", ")
            }
            sb.WriteString(col.Column)
        }
        sb.WriteString(") VALUES (")
        for i, col := range row {
            if i != 0 {
                sb.WriteString(", ")
            }
            sb.WriteString("'")
            sb.WriteString(col.Value)
            sb.WriteString("'")
        }
        sb.WriteString(");")
        result = append(result, sb.String())
        sb.Reset()
    }
	return result
}
