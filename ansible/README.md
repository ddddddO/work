# 概要
- 構成管理ツール

- 管理されるサーバーに必要なもの
    - Python
    - 管理するサーバーからssh接続可能なこと(管理するサーバーの公開鍵の登録)

- 管理するサーバーに必要なもの
    - Python
    - Ansible(と以下設定ファイル(必須ではない))
        - Inventory(どのサーバーを管理するかの記載)
        - ansible.cfg(Ansible全体の設定)
        - Playbook(メイン。管理されるサーバーに対してどういった構成にしたいかを記載)

# ドキュメント
- https://docs.ansible.com/
    - この中から、モジュール(playbookで指定するタスク)を探せる(「Module Index」)。どんなパラメータが必要か、とか

- https://www.ansible.com/products/ansible-training
    - [「Ansibleのこのセルフトレーニングラボは本当によくできていて、ネットワーク編でもネットワーク機器含めて環境が作られる。ブラウザ内のエディタとターミナルで説明どおりに色々試す仕組み。」](https://twitter.com/akira6592/status/1568428686336225280?s=20&t=Ph0GjmVInkrTgY5jakVl2Q) とのこと。良さそう。

## AnsibleとDocker(以下をまねて動作確認)
- https://tech-lab.sios.jp/archives/14355
