package functional_options

// Interface OMIT
type KV interface {
	Put(ctx context.Context, key, val string, opts ...OpOption) (*PutResponse, error)
	Get(ctx context.Context, key string, opts ...OpOption) (*GetResponse, error)
	Delete(ctx context.Context, key string, opts ...OpOption) (*DeleteResponse, error)
	Compact(ctx context.Context, rev int64, opts ...CompactOption) (*CompactResponse, error)
	Do(ctx context.Context, op Op) (OpResponse, error)
	Txn(ctx context.Context) Txn
}
// Interface# OMIT