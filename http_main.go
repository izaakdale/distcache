package main

// func main() {
// 	err := store.Init(
// 		store.WithConfig(
// 			store.Config{
// 				RedisAddr: "localhost:6379",
// 				Password:  "",
// 				DB:        0,
// 				RecordTTL: 10,
// 			},
// 		),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	mux := httprouter.New()
// 	mux.HandlerFunc(http.MethodPost, "/", Store)
// 	mux.HandlerFunc(http.MethodGet, "/:key", Retreive)
// 	log.Fatal(http.ListenAndServe("localhost:8080", mux))
// }

// type record struct {
// 	K string `json:"key,omitempty"`
// 	V string `json:"value,omitempty"`
// }

// func Store(w http.ResponseWriter, r *http.Request) {
// 	var rec record
// 	err := json.NewDecoder(r.Body).Decode(&rec)
// 	if err != nil {
// 		response.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("read body error"))
// 		return
// 	}

// 	err = store.Insert(rec.K, rec.V)
// 	if err != nil {
// 		response.WriteJsonError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	response.WriteJson(w, http.StatusOK, nil)
// 	return
// }

// func Retreive(w http.ResponseWriter, r *http.Request) {
// 	params := httprouter.ParamsFromContext(r.Context())
// 	key := params.ByName("key")
// 	if key == "" {
// 		response.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("please provide a key in url"))
// 		return
// 	}

// 	val, err := store.Fetch(key)
// 	if err != nil {
// 		if err == redis.Nil {
// 			response.WriteJsonError(w, http.StatusNotFound, fmt.Errorf("no records exist for that key"))
// 			return
// 		}
// 		response.WriteJsonError(w, http.StatusInternalServerError, fmt.Errorf("error fetching from store"))
// 		return
// 	}

// 	response.WriteJson(w, http.StatusOK, record{key, val})
// }
