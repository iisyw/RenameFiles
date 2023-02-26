**功能描述**

将当前工作路径中第 n 层的`所有文件及文件夹`重命名为`所在文件夹名-8 位随机字符串.后缀`，迁出`所在文件夹`并删除`所在文件夹`

**举例说明**

    当前工作路径：
    此程序.exe
    文件夹A(第一层文件夹)
      20220901(第二层文件夹)
        asdohasisaf.jpg(第三层文件)
        qweqwrasdfas.jpg(第三层文件)
        sandaofdna.png(第三层文件)
      20220902(第二层文件夹)
        adasknfadnf.exe(第三层文件)
        sadmpanfda.gif(第三层文件)
        asdjbocasda.txt(第三层文件)
    文件夹B(第一层文件夹)
      20220903(第二层文件夹)
        nfdosfdaif(第三层文件夹)
        dapsojfpaf.jpeg(第三层文件)
        asdsakfadfa.md(第三层文件)
      20220904(第二层文件夹)
        panfpandva(第三层文件夹)
        onasfdabsad.go(第三层文件)
        csnoafncoda.c(第三层文件)

**======运行程序，输入层数 3 进行确认======**  
_说明：将当前工作路径中第 3 层的`所有文件及文件夹`重命名为`所在文件夹名-8 位随机字符串.后缀`，迁出`所在文件夹`并删除`所在文件夹`_

    当前工作路径：
    此程序.exe
    文件夹A(第一层文件夹)
      20220901-1oascder.jpg(第二层文件)
      20220901-idjg4ltd.jpg(第二层文件)
      20220901-91jS0scd.png(第二层文件)
      20220902-8sJXNsax.exe(第二层文件)
      20220902-fosJaq2s.gif(第二层文件)
      20220902-asmCmse2.txt(第二层文件)
    文件夹B(第一层文件夹)
      20220903-fandpkfn(第二层文件夹)
      20220904-sadn3fma(第二层文件夹)
      20220903-qasfpoid.jpeg(第二层文件)
      20220903-AS9jsfhg.md(第二层文件)
      20220904-sdaclfgf.go(第二层文件)
      20220904-qwedsavx.c(第二层文件)

**更新**  
2023.2.26  
没想到竟然还有人收藏，修复一个bug，忘记增加`rand.Seed(time.Now().Unix())`初始化种子了，增加一下