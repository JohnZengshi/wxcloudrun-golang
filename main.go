/*
 * @LastEditors: John
 * @Date: 2022-11-12 06:38:37
 * @LastEditTime: 2022-11-12 09:39:42
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

	shareDir := "./client/bazi-app/dist" // 分享文件的路径 这里使用的是当前目录
	mux := http.NewServeMux()
	mux.Handle("/", shareServer(http.FileServer(http.Dir(shareDir))))

	// http.HandleFunc("/", service.IndexHandler)
	// mux.HandleFunc("/api/count", service.CounterHandler)

	mux.HandleFunc("/api/bazi", service.GetBazi)

	log.Fatal(http.ListenAndServe(":80", mux))
	// log.Fatal(http.ListenAndServe(":8081", mux))
}

func shareServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 打印来源ip及访问的文件夹/文件
		log.Printf("remote form ip:%s, uri: %s\n", r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
