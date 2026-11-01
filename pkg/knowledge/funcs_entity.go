//ff:func feature=knowledge type=adapter control=sequence
//ff:what Entity — EntityFn을 부른다
package knowledge

import "context"

// Entity implements Knowledge.
func (f Funcs) Entity(ctx context.Context, qid string) (EntityInfo, error) {
	return f.EntityFn(ctx, qid)
}
