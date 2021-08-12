# Stack and Heap

## 一般的Stack & Heap介紹

- https://nwpie.blogspot.com/2017/05/5-stack-heap.html
- https://medium.com/@yauhsienhuang/stack-%E8%88%87-heap-%E6%9C%89%E4%BD%95%E5%B7%AE%E5%88%A5-acdcc11263a0

## Golang的Stack & Heap

- https://golang.org/doc/faq#stack_or_heap
- http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#stack_heap_vars
- https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
- https://stackoverflow.com/questions/10866195/stack-vs-heap-allocation-of-structs-in-go-and-how-they-relate-to-garbage-collec

# 總結

- 別用中文翻譯去記,很難翻譯而且很難記

## Stack

- 靜態記憶體配置

        A call stack is a LIFO stack data structure that stores arguments, local variables, and other data tracked as a thread executes functions

將可預期的變數或是function等...的處理方法

## Heap

- 動態記憶體配置

將不可預期的變數或資料分配管理,且會被GC回收的處理方法

## Golang的Stack & Heap

    Go在使用上並沒有提到Stack or Heap  
    只會在GC(garbage-collected)內遵循一套嚴格的規則進行處理

    If you need to know where your variables are allocated pass the "-m" gc flag to "go build" or "go run" (e.g., go run -gcflags -m app.go).

