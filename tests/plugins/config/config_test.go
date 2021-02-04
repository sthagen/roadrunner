package config

import (
	"os"
	"os/signal"
	"testing"
	"time"

	endure "github.com/spiral/endure/pkg/container"
	"github.com/spiral/roadrunner/v2/plugins/config"
	"github.com/stretchr/testify/assert"
)

func TestViperProvider_Init(t *testing.T) {
	container, err := endure.NewContainer(nil, endure.RetryOnFail(true), endure.SetLogLevel(endure.ErrorLevel))
	if err != nil {
		t.Fatal(err)
	}
	vp := &config.Viper{}
	vp.Path = ".rr.yaml"
	vp.Prefix = "rr"
	err = container.Register(vp)
	if err != nil {
		t.Fatal(err)
	}

	err = container.Register(&Foo{})
	if err != nil {
		t.Fatal(err)
	}

	err = container.Init()
	if err != nil {
		t.Fatal(err)
	}

	errCh, err := container.Serve()
	if err != nil {
		t.Fatal(err)
	}

	// stop by CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	tt := time.NewTicker(time.Second * 2)
	defer tt.Stop()

	for {
		select {
		case e := <-errCh:
			assert.NoError(t, e.Error)
			assert.NoError(t, container.Stop())
			return
		case <-c:
			er := container.Stop()
			assert.NoError(t, er)
			return
		case <-tt.C:
			assert.NoError(t, container.Stop())
			return
		}
	}
}
