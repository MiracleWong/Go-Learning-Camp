# 作业

作业：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

回答：应该是需要Wrap的，带有信息抛给上层，这样上层调用者，可是通过errors.Is 打印相关的错误信息和堆栈信息

以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中：
https://jinshuju.net/f/EW97fj

作业提交截止时间 4 月 25 日（周日）23:59 前。