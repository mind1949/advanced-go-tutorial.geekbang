# 问题
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

# 解法
结合问题背景，再根据以下几点考虑：
* 业务层不能和数据库层耦合；
* 为避免业务层与具体的存储方式耦合，也不能业务层使用sql.ErrNoRows做判断；
* 为方便定位错误发生的流程，要方便查看栈信息；
* 为方便定位错误发生的上下文，要方便查看上下文信息；

可得出以下方案：

在一个第三方包（例如叫`bizerr`）定义一个的`ErrNotFound`错误，用`github.com/pkg/errors`的`Wrap`包装`ErrNotFound`错误，`sql.ErrNoRows`作为上下文信息；
具体代码：
```golang
if err == sql.ErrNoRows {
	return nil, errors.Wrapf(bizerr.ErrNotFound, "err: %s, query: %q, args: {id: %d}", err, query, id)
}
if err != nil {
	return nil, errors.Wrapf(err, "query: %q, args: {id: %d}", query, id)
}
```


