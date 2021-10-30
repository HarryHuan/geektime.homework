/* 
以下伪代码，具体场景具体分析。
此处只进行作业问题分析。

个人认为，无结果，必须作为错误上抛。
因为上面逻辑层可能会有后续数据操作，可能会造成错误，
即使没有错误，也会造成额外处理开销。
*/

db, _ := sql.Open("driverName", "dataSourceName")
row, _ := db.QueryRow("select name from customer where id = ?", 1)
err := row.Scan(...)
if err != nil {
	if err == sql.ErrNoRows {
		err = errors.Wrap(err, "querySet is nil")
	}
	return err
}