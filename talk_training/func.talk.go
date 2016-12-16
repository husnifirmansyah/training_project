package talk_training

import (
    "database/sql"
    _ "github.com/lib/pq"
    "time"
)

func GetTalks(product_id string) []Talks {
  db, err := sql.Open("postgres", "user=techacademy password=123qwe!@#QWE dbname=tokopedia-talk host=192.168.100.126 port=5432 sslmode=disable")
  checkErr(err)
  //defer db.Close()

	rows, err := db.Query("SELECT talk_id, product_id, message, create_time FROM ws_talk WHERE product_id = " + product_id)
  //talk_training.checkErr(err)
  var arr_talks []Talks
  for rows.Next() {
      var ID int64
      var ProductID int64
      var Message string
      var CreateTime time.Time
      err = rows.Scan(&ID, &ProductID, &Message, &CreateTime)
      arr_talks = append(arr_talks, Talks {
        ID:         ID,
        ProductID:  ProductID,
        Message:    Message,
        CreateTime: time.Now(),
      })
  }
  return arr_talks
}

/*func InsertTalks(product_id string, user_id string, message string) {
  db, err := sql.Open("postgres", "user=techacademy password=123qwe!@#QWE dbname=tokopedia-talk host=192.168.100.126 port=5432 sslmode=disable")
  checkErr(err)
  //defer db.Close()
  //err = db.Exec("INSERT INTO ws_talk (product_id, message, create_time, create_by) VALUES ($1, $2, $3, $4)", product_id, message, time.Now().String(), user_id)
	//talk_training.checkErr(err)
}*/
