# coding-interview-university_by_golang

# 前言
自2019/9,在PTT上找到了人生第一份程式設計工作,也是在第一家公司開始學Go的

目前在準備computer science方面的知識,目標是希望能夠成為獨當一面的軟體工程師  
這個Repository主要使用Go進行實做,以及包含以下項目為主要目標  
並將其中的項目獨立出來研究撰寫筆記    
README.md的部分會是索引資源  
實作跟心得筆記類的會放在資料夾

* Go
* Database-SQL
* Docker
* [Coding Interview University](https://github.com/jwasham/coding-interview-university/blob/main/translations/README-tw.md)
* [system-design-primer](https://github.com/donnemartin/system-design-primer)

# Go 學習資源

我在學習Go的過程中,這些資源幾乎涵蓋了我所有問題的解答  
還有我看過的實體書開放在網路上讓人免費學習,可以省下一些花費

- [Go的官方網站](https://golang.org/doc/) 推薦先看下面這些大標題的內容,剛好可以練習英文

    - A Tour of Go
    - Effective Go
    - Package Documentation
    - Go Modules Reference

- [Go語言聖經](https://github.com/gopl-zh/gopl-zh.github.com)
- [Go语言高级编程 (Advanced Go Programming)](https://github.com/chai2010/advanced-go-programming-book)  
- Stack & Heap 詳見Go資料夾內整理
- [Slice tricks](https://ueokande.github.io/go-slice-tricks/)
- [Go official wiki](https://github.com/golang/go/wiki)
- [Go 面試問題整體](https://zhuanlan.zhihu.com/p/471490292)

# 演算法-Algorithm

* [漸進分析(Asymptotic analysis)](https://www.itread01.com/content/1550231649.html)
* [Dynamic Programming](http://web.ntnu.edu.tw/~algo/DynamicProgramming.html) 這個網站還有包含其他computer science方面的知識
* [Know Thy Complexities](https://www.bigocheatsheet.com/) 常見演算法時間複雜度整理

# Database-SQL
## MySQL
- [MySQL](https://www.itread01.com/study/mysql-tutorial.html) 這個網站還包含了前端/程式語言/Linux/Docker/Redis...等教學
- [Use MySQL 8.0 Explain](https://dev.mysql.com/doc/refman/8.0/en/execution-plan-information.html) 使用 MySQL Explain 查詢 Query 效能並做優化的工具
- [MySQL 8.0 Optimization](https://dev.mysql.com/doc/refman/8.0/en/optimization.html) MySQL 可用的優化方式文件
## TiDB
- [TiDB](https://docs.pingcap.com/zh/tidb/stable/performance-tuning-overview) 主要是借用文件內的優化概念，可以挪用到別的DB上運用類似的概念，文件解釋的很清楚，還有些實際案例可以參考

# 設計模式

- 深入淺出設計模式(Head First Design Patterns) ISBN:9867794524
- 無瑕的程式碼：整潔的軟體設計與架構篇 ISBN:9789864342945
- [Clean Architecture 概念篇 + Dependency injection](https://ithelp.ithome.com.tw/articles/10240228?sc=iThomeR)
- [(IoC) Inversion of Control + (DI) Dependency Inversion Principle](https://iter01.com/562085.html)

# 網路協定

- [TCP/IP協定]http://kevin.hwai.edu.tw/~kevin/material/EAssistant/TCP.htm)
- [TCP/IP 協定與 Internet 網路](http://www.tsnien.idv.tw/Internet_WebBook/Internet.html) 這是博士寫的電子書,講解的很細緻
- [TCP/UDP的差別](https://www.tsnien.idv.tw/Network_WebBook/chap13/13-6%20TCP%20%E9%80%9A%E8%A8%8A%E5%8D%94%E5%AE%9A.html)
- [SSL/TSL](https://www.websecurity.digicert.com/zh/hk/security-topics/what-is-ssl-tls-https)
- [HTTP/2](https://developers.google.com/web/fundamentals/performance/http2?hl=zh-cn) 跟grpc有關
- [WebSocket](https://docs.microsoft.com/zh-tw/archive/msdn-magazine/2012/may/cutting-edge-understanding-the-power-of-websockets)

# Docker

- [Docker —— 從入門到實踐](https://github.com/philipz/docker_practice)
- [docker docs](https://docs.docker.com/language/golang/build-images/)
- [使用 Docker 封裝與運行 Go 程式](https://ithelp.ithome.com.tw/articles/10240352) 系列文,同個作者有很多可以參考的文章
- [使用 Docker Compose 摻在一起做懶人包](https://ithelp.ithome.com.tw/articles/10243618) 系列文,同個作者有很多可以參考的文章
- [Kitematic — Docker 圖形化管理工具](https://medium.com/@bee811101/kitematic-docker-%E5%9C%96%E5%BD%A2%E5%8C%96%E7%AE%A1%E7%90%86%E5%B7%A5%E5%85%B7-60ffe5e3605a)

# System Design

- [秒殺系統](https://juejin.cn/post/7203136448333332535) 裡面有大量系統設計的案例跟講解，非常不錯

# Git的坑

- [刪除 root commit](https://stackoverflow.com/questions/10911317/how-to-remove-the-first-commit-in-git/32765827#32765827)
- [GitHub SSH](https://gist.github.com/xirixiz/b6b0c6f4917ce17a90e00f9b60566278) 不然會因為認證問題沒辦法 push，clone 時是也要用 ssh clone 後面才會正常
