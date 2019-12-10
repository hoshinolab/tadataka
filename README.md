**このデータはGitHub上にも公開しています**

https://github.com/hoshinolab/tadataka/tree/alpha-release

# tadataka

![ロゴマーク](./docs/tadataka-logo-small.png)

tadataka は主に日本国内の地理空間情報ビッグデータを高速に処理するための前処理ツールです。名称は江戸時代の測量家である[伊能忠敬](https://ja.wikipedia.org/wiki/%E4%BC%8A%E8%83%BD%E5%BF%A0%E6%95%AC)に由来しています。

2019年11月25日時点で巨大なCSVを分割する機能および逆ジオコーダ（座標→住所）を提供しています。近い将来、ジオコーダ（住所→座標）も実装される予定です。

### ロゴマークに関して

作者のryo-aがAdobe Illustratorで作成したものです。伊能忠敬が測量に用いた象限儀をモチーフにしています。

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

特に住居住所表示データに関しては基本測量成果に該当するため、利用時に国土地理院の規約および法令を確認した上でダウンロードする必要があります。

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
データベース上で近傍の住所が見つからなかった場合、`NA` を返します。

**必要なもの** : Redis および `stdby` コマンド実行済みの状態

```sh
$ tadataka rgc ./input/file/path.csv ./output/file/path.csv --lat 1 --lng 2
```

- `lat`: (整数値) CSVファイルにおいて緯度が記されている列の番号 ( `0` 始まり)
- `lng`: (整数値) CSVファイルにおいて経度が記されている列の番号 ( `0` 始まり)
    - `user000,30.123456,145.456789,10,true` のような CSVファイルの場合、 `lat` は `1` で `lng` は `2` になります。

## 実装予定のサブコマンド群

### status

`~/.tadataka` ディレクトリおよび Redis の状態を取得するサブコマンド

### gc (geocoder)

ジオコーダ。近日中に実装を検討していますが、住所正規化処理が存在するため、逆ジオコーダに比べると処理内容が煩雑になると考えられます。

国内住所を取り扱うジオコーダの先行研究としては東大CSISのDAMS(Distributed Address Matching System)が挙げられます。

- http://newspat.csis.u-tokyo.ac.jp/geocode/modules/dams/index.php?content_id=1


## アルゴリズム

Redisに格納した国土交通省および国土地理院のデータをもとに、与えられた座標と最も近いデータを逆ジオコーディングの結果として扱います。
住居表示住所は「号」レベルまでのデータが整備されているのですが、一部の市区町村のみでしか提供されていません。一方で位置参照情報はカバーしている市町村が広い一方、番地ごとの代表点しかデータがありません。そこで、住居表示住所を優先して探索し、該当データがない場合に位置参照情報のデータを利用する形としています。

　探索時、全国のデータに総当たりしていると計算量が膨大になる探索するにあたり、データを細かく分割する必要があります。分割にはGoogleが提唱しているOpen Location Code(OLC)を採用しました。OLCを用いると、座標を一意の文字列に変換することができます。

　また、OLCの下位コードは上位コードを包含しています。例えば、8Q7XJPXR+FXというコードは8Q7XJPXR+というエリアに含まれており、更には上位の8Q7XJPというエリアに含まれていることもわかります。
これを用いて、8Q7XJPXR+FXと変換される位置情報に対しては8Q7XJPXR+のグリッドに含まれる住所から最近傍のものを見つけ、それを住所とみなします。

　また、使用するデータはインメモリデータベースであるRedisに格納しています。Redisはデータベースエンジンの中でも高速であることが知られており、処理内容にもよるがマイクロ秒単位で保存されているデータの取得を行うことができる。RedisではOLCのグリッド（8Q7XJPXR+等）をキーとしてlist型で格納しています。


### 改善予定のアルゴリズム

　グリッドの端部に位置する座標は隣接するグリッドにより近い住所が存在する可能性もあるため、併せて隣接するグリッドも探索することでより高精度な住所探索が行えると考えています。
しかし、探索するグリッドが増加することで処理時間が数倍に増加してしまうため、的確にグリッドを絞り込む実装を検討しています。


## 評価（速度パフォーマンス）

- AWS EC2 r4.16xlarge インスタンス上で動作
- CPU:Intel(R) Xeon(R) CPU E5-2686 v4 @ 2.30GHz
- RAM: 488GB
- Storage: EBS

time コマンドを用いた計測結果です。

### 10,846件の座標データ 逆ジオコーディング

駅データ.jp(20190928)の無料版から座標のみ抽出したものです。テストデータは以下に公開しています。
- https://gist.github.com/ryo-a/8c40f1189610c9cd31ea29e0778c0433


出力結果は以下に公開しています。
- https://gist.github.com/ryo-a/a13e29f170ba8ab0af48495344977a3b

```
real    0m1.887s
user    0m1.463s
sys     0m0.246s
```

1秒あたり約5747.74件の逆ジオコーディングを処理しました。<br>
なお、NA（住所を判定できなかった）件数は1894件（約17%）です。

### 36,519,587件の座標データ 逆ジオコーディング

共同研究先からの提供データです。日本国内を中心としますが、一部海外の座標も含みます。  
なお、海外座標は国内で特定できなかった住所と同様にNAとなります。

（NDAの都合上、対象のデータ内容は公開できません）

```
real    115m11.699s
user    48m19.484s
sys     38m49.922s
```

1秒あたり約5352.37件の逆ジオコーディングを処理しました。<br>
なお、NA（住所を判定できなかった）件数は4525555件（約12%）です。

