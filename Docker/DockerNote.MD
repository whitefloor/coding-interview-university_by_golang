# DockerNote

## Dockerfile範例

1. 編寫Dockerfile

        # 井號開頭是註記

        # FROM 環境
        FROM golang

        #執行shell command -p是option,可以在缺少父目錄的情況下創建
        RUN mkdir -p /app

        # 為後續的 RUN、CMD、ENTRYPOINT 指令指定工作目錄。
        WORKDIR /app

        # 複製本地端的檔案
        COPY . .

        # 如果你有設定檔要複製
        COPY ./yourfloder/config.json .

        # 執行go module download 用這種方法就不需要使用一堆go get命令
        RUN go mod download

        # 執行go build
        RUN go build -o app ./yourfloder/main.go

        ENTRYPOINT ["./app"]

2. 建立image

        docker build -t "yourtag" .

    最後的.會讀取路徑底下的Dockerfile(包括子目錄)來建立image

3. 運行

        docker run yourtag