## 概要

Go によるバックエンドのベース  
WebFW は echo を用いる

ローカルにおいてサーバは Docker 上で実行され Air によってオンタイムに修正が反映される

firebase は認証に用いている  
※client もこれに合わせる必要があるが、  
　認証を変えようとなると結局全般で作り直すことになる

### 残

Migration の対応
moduleが昔作成した`github.com/watariRyo/baby-map-server`のままなので置換

## 技術スタック

| 技術                   | 分類             | バージョン |
| ---------------------- | ---------------- | ---------- |
| Golang                 | 言語             | 1.22       |
| echo                   | WebFW            | 4.11.4     |
| firebase.google.com/go | firebase（認証） | 3.13.0     |
| redis                  | キャッシュサーバ | -          |
| MySQL                  | DB               | 8.x        |
| sqlboiler              | ORM              | 4.16.2     |
| viper                  | config           | 1.18.2     |
