package db

// AddArticle 添加文章
func AddArticle(username string, title string, email string, content string, cateId int64, summary string) (id int64, err error) {
	sql := `insert into article (user_name,email,cate_id,title,summary,content) values (?,?,?,?,?,?)`
	exec, err := db.Exec(sql, username, email, cateId, title, summary, content)
	if err != nil {
		return 0, err
	}
	id, err = exec.LastInsertId()
	if err != nil {
		return
	}
	return
}