package store

//
//type recordReader struct {
//	client Transactioner
//	key    string
//}
//
//func (c *Client) Reader() io.Reader {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	keys, err := c.AllKeys("")
//	if err != nil {
//		panic("implement me") //TODO
//		return nil
//	}
//	readers := make([]io.Reader, len(keys))
//	for i, key := range keys {
//		readers[i] = &recordReader{c, key}
//	}
//	return io.MultiReader(readers...)
//}
//
//func (r *recordReader) Read(p []byte) (int, error) {
//	val, err := r.client.Fetch(r.key)
//	if err != nil {
//		return 0, err
//	}
//	p, err = proto.Marshal(&v1.StoreRequest{
//		Record: &v1.KVRecord{
//			Key:   r.key,
//			Value: val,
//		},
//	})
//	if err != nil {
//		return len(p), err
//	}
//	return len(p), io.EOF
//}
