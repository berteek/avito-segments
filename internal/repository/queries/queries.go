package queries

import (
    "fmt"
    "strings"
)

func InsertIntoUsers() string {
    return "INSERT INTO users DEFAULT VALUES;"
}

func InsertIntoSegments(name string) string {
    return fmt.Sprintf("INSERT INTO segments (name) VALUES ('%s');", name)
}

func DeleteFromSegments(name string) string {
    return fmt.Sprintf("DELETE FROM segments WHERE name='%s';", name)
}

func AddAndDeleteSegmentsFromUser(userID int, segAdd, segDel []string) string {
    query := "BEGIN; "
    query += fmt.Sprintf("INSERT INTO segments_users (segment_id, user_id) SELECT id, %d", userID)
    query += " FROM segments WHERE name IN ("
    quotedSegAdd := make([]string, len(segAdd))
    for i, seg := range segAdd {
        quotedSegAdd[i] = fmt.Sprintf("'%s'", seg)
    }
    query += strings.Join(quotedSegAdd, ", ")
    query += "); "
    query += "DELETE FROM segments_users WHERE segment_id IN (SELECT id FROM segments WHERE NAME IN ("
    quotedSegDel := make([]string, len(segDel))
    for i, seg := range segDel {
        quotedSegDel[i] = fmt.Sprintf("'%s'", seg)
    }
    query += strings.Join(quotedSegDel, ", ")
    query += ")); COMMIT;"
    return query
}

func ActiveSegmentsForUser(userID int) string {
    return fmt.Sprintf("SELECT segments.name FROM segments_users INNER JOIN segments ON segments_users.segment_id = segments.id WHERE segments_users.user_id = %d;", userID)
}

func InsertIntoHistory(userID int, segName, operation string) string {
    return fmt.Sprintf("INSERT INTO segments_users_history (user_id, segment_name, operation) VALUES (%d, '%s', '%s');", userID, segName, operation)
}
