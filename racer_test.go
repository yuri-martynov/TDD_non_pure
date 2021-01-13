package racer

import ("testing")

func Test_Site_Fast_Is_Fastest(t *testing.T) {
    slowURL := "http://site.slow"
    fastURL := "http://site.fast"

    ping := func(url string) chan bool {
        ch := make(chan bool)
        go func() {
            if (url == fastURL) {
                ch <- true
            }
        }()
        return ch
    }

    want := fastURL
    got := racer(slowURL, fastURL, ping)

    if got != want{
        t.Errorf("got '%s', want '%s'", got, want)
    }
}


func Test_Site_Super_Fast_Is_Fastest(t *testing.T) {
    slowURL := "http://site.fast"
    fastURL := "http://site.superfast"

    ping := func(url string) chan bool {
        ch := make(chan bool)
        go func() {
            if (url == fastURL) {
                ch <- true
            }
        }()
        return ch
    }

    want := fastURL
    got := racer(slowURL, fastURL, ping)

    if got != want{
        t.Errorf("got '%s', want '%s'", got, want)
    }
}

func Test_Fastest_Is_First_Param(t *testing.T) {
    slowURL := "http://site.fast"
    fastURL := "http://site.superfast"

    ping := func(url string) chan bool {
        ch := make(chan bool)
        go func() {
            if (url == fastURL) {
                ch <- true
            }
        }()
        return ch
    }

    want := fastURL
    got := racer(fastURL, slowURL, ping)

    if got != want{
        t.Errorf("got '%s', want '%s'", got, want)
    }
}

func Test_Integration(t *testing.T) {
    slowURL := "https://google.com"
    fastURL := "https://yandex.ru"

    want := fastURL
    got := racer(fastURL, slowURL, ping)

    if got != want{
        t.Errorf("got '%s', want '%s'", got, want)
    }
}







