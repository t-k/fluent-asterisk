fluent-logger-asterisk
====

## Small script which behaves as an API for fluentd.

## How to install

```
git clone https://github.com/t-k/fluent-logger-asterisk
go get github.com/t-k/fluent-logger-golang/fluent
go build fluent.go
```

##Example

```
; /etc/asterisk/extensions.conf
; System(path_to_fluent_script tag_prefix message})

exten => _1000,1,Answer()
 same => n,System(/etc/asterisk/fluent echo CALLERID:${CALLERID(num)})
 same => n,Playback(beep)
 same => n,Echo()
 same => n,Hangup()

```

## Setting config values

```
# bash_profile

export FLUENT_TAG='foobar'
export FLUENT_HOST='foo.com'
export FLUENT_PORT='8888'
```