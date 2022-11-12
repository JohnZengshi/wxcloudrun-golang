/*
 * @LastEditors: John
 * @Date: 2022-11-12 06:38:37
 * @LastEditTime: 2022-11-12 07:04:18
 * @Author: John
 */
package main

import (
	"log"
	"net/http"

	"wxcloudrun-golang/service"
)

func main() {
	// if err := db.Init(); err != nil {
	// 	panic(fmt.Sprintf("mysql init failed with %+v", err))
	// }

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

	http.HandleFunc("/api/bazi", service.GetBazi)

	// log.Fatal(http.ListenAndServe(":80", nil))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
