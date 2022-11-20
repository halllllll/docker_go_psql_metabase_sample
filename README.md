# Whats This?

`docker-comopse`でやる、metagbase,postgresql,golangアプリの連携のサンプル。

# my env

同一ネットワーク下のm1 macとwindows11で確認。mac上のDockerで動かし、windowsからmetabaseへIPアドレス:ポート番号で接続してみた。

# memo
- `metabase`が起動するのけっこう時間かかるし、なぜかたまに失敗する。Docker Desktopのターミナルでしばらく監視してたが、原因はよくわからない
- Dockerのサンプルとしては、postgres.confを用意してDockerfileでCOPYするとモアベターかもしれない
