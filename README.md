# YaDisk API

This is an unofficial Yandex.Disk API in Go.

I only implemented what I needed.

## Usage

Add it to your project:

```go
import (
  yadisk "github.com/MOZGIII/yadisk-api"
)
```

And then just use it:

```go
api := yadisk.NewAPI(oAuthToken)
err := api.Upload(reader, "app:/test.txt", true)
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
