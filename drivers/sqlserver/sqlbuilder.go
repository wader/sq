package sqlserver

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/neilotoole/sq/libsq/sqlbuilder"

	"github.com/neilotoole/lg"

	"github.com/neilotoole/sq/libsq/ast"
	"github.com/neilotoole/sq/libsq/errz"
	"github.com/neilotoole/sq/libsq/sqlmodel"
	"github.com/neilotoole/sq/libsq/sqlz"
)

type fragBuilder struct {
	sqlbuilder.BaseFragmentBuilder
}

func newFragmentBuilder(log lg.Log) *fragBuilder {
	r := &fragBuilder{}
	r.Log = log
	r.Quote = `"`
	r.ColQuote = `"`
	r.Ops = sqlbuilder.BaseOps()
	return r
}

func (fb *fragBuilder) Range(rr *ast.RowRange) (string, error) {
	if rr == nil {
		return "", nil
	}

	/*
		SELECT * FROM tbluser
			ORDER BY (SELECT 0)
			OFFSET 1 ROWS
			FETCH NEXT 2 ROWS ONLY;
	*/

	if rr.Limit < 0 && rr.Offset < 0 {
		return "", nil
	}

	offset := 0
	if rr.Offset > 0 {
		offset = rr.Offset
	}

	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("OFFSET %d ROWS", offset))

	if rr.Limit > -1 {
		buf.WriteString(fmt.Sprintf(" FETCH NEXT %d ROWS ONLY", rr.Limit))
	}

	sql := buf.String()
	fb.Log.Debugf("returning SQL fragment: %s", sql)
	return sql, nil
}

type queryBuilder struct {
	sqlbuilder.BaseQueryBuilder
}

func (qb *queryBuilder) SQL() (string, error) {
	// SQL Server handles range (OFFSET, LIMIT) a little differently. If the query has a range,
	// then the ORDER BY clause is required. If ORDER BY is not specified, we use a trick (SELECT 0)
	// to satisfy SQL Server. For example:
	//
	//   SELECT * FROM tbluser
	//   ORDER BY (SELECT 0)
	//   OFFSET 1 ROWS
	//   FETCH NEXT 2 ROWS ONLY;
	if qb.RangeClause != "" {
		if qb.OrderByClause == "" {
			qb.OrderByClause = "ORDER BY (SELECT 0)"
		}
	}

	return qb.BaseQueryBuilder.SQL()
}

func dbTypeNameFromKind(kind sqlz.Kind) string {
	switch kind {
	default:
		panic(fmt.Sprintf("unsupported datatype %q", kind))
	case sqlz.KindUnknown:
		return "NVARCHAR(MAX)"
	case sqlz.KindText:
		return "NVARCHAR(MAX)"
	case sqlz.KindInt:
		return "BIGINT"
	case sqlz.KindFloat:
		return "FLOAT"
	case sqlz.KindDecimal:
		return "DECIMAL"
	case sqlz.KindBool:
		return "BIT"
	case sqlz.KindDatetime:
		return "DATETIME"
	case sqlz.KindTime:
		return "TIME"
	case sqlz.KindDate:
		return "DATE"
	case sqlz.KindBytes:
		return "VARBINARY(MAX)"
	}
}

// createTblKindDefaults is a map of Kind to the value
// to use for a column's DEFAULT clause in a CREATE TABLE statement.
var createTblKindDefaults = map[sqlz.Kind]string{
	sqlz.KindText:     `DEFAULT ''`,
	sqlz.KindInt:      `DEFAULT 0`,
	sqlz.KindFloat:    `DEFAULT 0`,
	sqlz.KindDecimal:  `DEFAULT 0`,
	sqlz.KindBool:     `DEFAULT 0`,
	sqlz.KindDatetime: `DEFAULT '1970-01-01T00:00:00'`,
	sqlz.KindDate:     `DEFAULT '1970-01-01'`,
	sqlz.KindTime:     `DEFAULT '00:00:00'`,
	sqlz.KindBytes:    `DEFAULT 0x`,
	sqlz.KindUnknown:  `DEFAULT ''`,
}

// buildCreateTableStmt builds a CREATE TABLE statement from tblDef.
// The implementation is minimal: it does not honor PK, FK, etc.
func buildCreateTableStmt(tblDef *sqlmodel.TableDef) string {
	sb := strings.Builder{}
	sb.WriteString(`CREATE TABLE "`)
	sb.WriteString(tblDef.Name)
	sb.WriteString("\" (")

	for i, colDef := range tblDef.Cols {
		sb.WriteString("\n\"")
		sb.WriteString(colDef.Name)
		sb.WriteString("\" ")
		sb.WriteString(dbTypeNameFromKind(colDef.Kind))

		if colDef.NotNull {
			sb.WriteRune(' ')
			sb.WriteString(createTblKindDefaults[colDef.Kind])
			sb.WriteString(" NOT NULL")
		}

		if i < len(tblDef.Cols)-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteString("\n)")

	return sb.String()
}

// buildUpdateStmt builds an UPDATE statement string.
func buildUpdateStmt(tbl string, cols []string, where string) (string, error) {
	if len(cols) == 0 {
		return "", errz.Errorf("no columns provided")
	}

	sb := strings.Builder{}
	sb.WriteString(`UPDATE "`)
	sb.WriteString(tbl)
	sb.WriteString(`" SET "`)
	sb.WriteString(strings.Join(cols, `" = ?, "`))
	sb.WriteString(`" = ?`)
	if where != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(where)
	}

	s := replacePlaceholders(sb.String())
	return s, nil
}

// replacePlaceholders replaces all instances of the question mark
// rune in input with $1, $2, $3 placeholders.
func replacePlaceholders(input string) string {
	if input == "" {
		return input
	}

	var sb strings.Builder
	pCount := 1
	var i int
	for {
		i = strings.IndexRune(input, '?')
		if i == -1 {
			sb.WriteString(input)
			break
		} else {
			// Found a ?
			sb.WriteString(input[0:i])
			sb.WriteString("@p")
			sb.WriteString(strconv.Itoa(pCount))
			pCount++
			if i == len(input)-1 {
				break
			}
			input = input[i+1:]
		}
	}

	return sb.String()
}