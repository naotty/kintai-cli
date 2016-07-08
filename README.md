# kintai-cli

kintai-cli is puncing time tool for KINTAI.


## Usage

```
$ kintai-cli -version
version v0.1.0


$ kintai-cli -show-config
url: https://example.com
user_id: 1
token: hogehoge


# puncing time
$ kintai-cli
{
    "meta": {
        "url": "hoge",
        "method": "POST"
    },
    "response": {
        "result": true,
        "message": fugafuga
    }
}

```


## Config
you need config.yml.
see config.yml.sample file.
