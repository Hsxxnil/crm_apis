# CRM APIs

一套以 **Golang** + **PostgreSQL** 為後端、**Angular** 為前端框架開發的 **顧客關係管理平台**，本專案旨在透過簡單易懂的操作介面，協助業務有效管理客戶，將顧客依照業務流程歸類於「潛在線索」至「成交商機」的各個階段，並可隨時記錄與客戶之間的洽談重點與互動紀錄。系統同時支援報價與訂單建立流程，協助業務團隊全面掌握客戶狀況，提升銷售效率與客戶關係經營品質。

## 專案連結

* 相關文件：[點我查看](https://hsxxnil.notion.site/Collective-11c5b51f95f58185ba96dcb6fde626e1)
* Swagger API 文件：[點我查看](https://hsxxnil.github.io/swagger-ui/?urls.primaryName=CRM)

## 安裝
1. 下載專案

```bash
git clone https://github.com/Hsxxnil/crm_apis.git
cd crm_apis
```

2. 建立 Makefile

> 請根據您的作業系統選擇對應的範本進行複製：
* Linux / macOS
```bash
cp Makefile.example.linux Makefile
```

* Windows
```bash
copy Makefile.example.windows Makefile
```


3. 初始化

> 如為初次建立開發環境，請先根據您的作業系統安裝必要套件：
* Linux / macOS
```bash
brew install golang-migrate golangci-lint protobuf
```

* Windows（建議使用 Scoop，或手動安裝以下套件）：
```bash
scoop install golang-migrate golangci-lint protobuf
```

> 執行以下指令將自動安裝依賴套件並建立必要的目錄結構：
```bash
make setup
```

4. 設定環境參數

> 開啟並編輯以下檔案，填入資料庫連線資訊、JWT 金鑰等必要參數：
```file
config/debug_config.go
```

## 資料庫遷移

> 執行以下指令使用[golang-migrate](https://github.com/golang-migrate/migrate)做資料庫遷移及做資料表版控：
```bash
make migration
```

## 執行
> 執行以下指令在本地端啟動伺服器並自動重載：
```bash
make air
```

## License

本專案使用的 [Vodka](https://github.com/dylanlyu/vodka) 採用 [MIT License](https://opensource.org/licenses/MIT) 授權。
