# tadataka

![ロゴマーク](./docs/tadataka-logo-small.png)

tadataka は主に日本の地理空間情報ビッグデータを高速に処理するための前処理ツールです。名称は江戸時代の測量家である[伊能忠敬](https://ja.wikipedia.org/wiki/%E4%BC%8A%E8%83%BD%E5%BF%A0%E6%95%AC)に由来しています。

2019年11月25日時点で巨大なCSVを分割する機能および逆ジオコーダ（座標→住所）を提供しています。近い将来、ジオコーダ（住所→座標）も実装される予定です。

## インストール

現時点ではUNIXライクOSのみサポートしています。（GNU/Linuxが推奨されます。開発者はUbuntuにて動作確認をしています。）
Windowsは将来的にサポートされます。

### Redis

tadatakaは住所-座標データの格納にインメモリデータベースの Redis を使用するため、tadatakaのインストール前にRedisをインストールしておく必要があります。
インストール方法は指定されていませんので、使用する環境において最も適したものをご利用ください。

- [Redis](https://redis.io/) (version 4.0.0 以降が必要です。)

インストール後、`$ redis-server` でRedisを起動してください。また、必要に応じて自動起動スクリプトを有効にしてください。

### tadataka

**現時点ではソースからのインストールのみサポートしています**

このrepositoryをcloneし、`make build` を実行してください。（Go 1.13以降が必要です）<br>
ビルド後、`~/.tadataka/bin` を PATH に追加してください。

```
$ git clone https://github.com/hoshinolab/tadataka.git
$ make build
$ echo export PATH='$HOME/.tadataka/bin:$PATH' >> ~/.bash_profile
$ source ~/.bash_profile
```

バイナリの配布は近日中に行う予定です。

### tadataka の基本的なセットアップ

初回起動時、日本国内の住所データをダウンロードする必要があります。以下のコマンドを実行してください。<br>
CLI内で国土交通省・国土地理院のサイトにおける利用規約などを表示します。同意が出来たらダウンロードしてください。<br>
また両サイトに負荷を掛けないため、一定のインターバルを設けてダウンロードします。

```
$ tadataka download
```

ダウンロード完了後、`stdby` サブコマンドで住所データをRedisに読み込みます。

```
$ tadataka stdby
```

ディスクI/O速度にもよりますが、数十分程度の時間を要する場合があります。2回目以降はRedisのダンプデータを用いることで読み込みの手間と所要時間を省けます。

Redisではデータベース番号として 10 および 11 を使用します。

## サブコマンド


- `download`:  `~/.tadataka` ディレクトリに住所データをダウンロードします。
- `stdby`: `~/.tadataka` 内のデータをRedisに読み込みます。先に `download` コマンドを実行してください。 `rgc` コマンドを実行する場合はこの操作が必要です。
- `subdiv`: 巨大なCSVファイルを[Open Location Code (plus codes)](https://en.wikipedia.org/wiki/Open_Location_Code)に基づいたグリッド単位に分割します。
- `rgc`: リバースジオコーディングを実行します。
- `version`: tadatakaのバージョンを表示します。

### `download`

国土交通省の位置参照情報および国土地理院の住居住所表示データを `~/.tadataka` にダウンロードします。

**必要なもの** : 十分なストレージ（空き容量6GB程度）

```
$ tadataka download
```



### `stdby`

Redisにデータを読み込みます。

**必要なもの** : `download` コマンドでダウンロードした住所データ

```
$ tadataka stdby
```


### `subdiv`

巨大なCSVファイルをOpen Location Codeによるグリッド単位に分割します。

```sh
$ tadataka olc ./input/file/path.csv ./output/directory/path --lat 1 --lng 2 --header false
```

- `lat`: (整数値) CSVファイルにおいて緯度が記されている列の番号 ( `0` 始まり)
- `lng`: (整数値) CSVファイルにおいて経度が記されている列の番号 ( `0` 始まり)
    - `user000,30.123456,145.456789,10,true` のような CSVファイルの場合、 `lat` は `1` で `lng` は `2` になります。
- `header`: (boolean) CSVファイルにヘッダがあるかどうか (default: `true`)



### `rgc`

インメモリデータベースを利用した高速逆ジオコーダ (Reverse Geocoder) を実行します。

**必要なもの** : Redis および `stdby` コマンド実行済みの状態

```sh
$ tadataka rgc ./input/file/path.csv ./output/file/path.csv --lat 1 --lng 2
```

- `lat`: (整数値) CSVファイルにおいて緯度が記されている列の番号 ( `0` 始まり)
- `lng`: (整数値) CSVファイルにおいて経度が記されている列の番号 ( `0` 始まり)
    - `user000,30.123456,145.456789,10,true` のような CSVファイルの場合、 `lat` は `1` で `lng` は `2` になります。