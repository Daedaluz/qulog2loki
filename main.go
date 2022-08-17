package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/backoff"
	"github.com/grafana/loki-client-go/pkg/labelutil"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	"github.com/influxdata/go-syslog/v3"
	"github.com/influxdata/go-syslog/v3/nontransparent"
	"github.com/influxdata/go-syslog/v3/octetcounting"
	"github.com/influxdata/go-syslog/v3/rfc5424"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	log    = logrus.New()
	client *loki.Client
)

func parseStream(r io.Reader, callback func(res *syslog.Result)) error {
	buf := bufio.NewReader(r)

	firstByte, err := buf.Peek(1)
	if err != nil {
		return err
	}

	b := firstByte[0]
	if b == '<' {
		nontransparent.NewParser(syslog.WithListener(callback)).Parse(buf)
	} else if b >= '0' && b <= '9' {
		octetcounting.NewParser(syslog.WithListener(callback)).Parse(buf)
	} else {
		return fmt.Errorf("invalid or unsupported framing. first byte: '%s'", firstByte)
	}

	return nil
}

var (
	listen         = flag.String("listen", ":1514", "Listen address")
	lokiBase       = flag.String("loki", "http://localhost:3100", "loki base url")
	useMachineTime = flag.Bool("machine-time", false, "use machine time instead of message time")
)

func main() {
	flag.Parse()
	server, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatal(err)
	}
	lokiUrl, _ := url.Parse(fmt.Sprintf("%s/loki/api/v1/push", *lokiBase))
	client, err = loki.New(loki.Config{
		URL: urlutil.URLValue{
			URL: lokiUrl,
		},
		BatchWait: time.Second * 1,
		BatchSize: 100000,
		BackoffConfig: backoff.BackoffConfig{
			MinBackoff: time.Second * 1,
			MaxBackoff: time.Second * 5,
			MaxRetries: 20,
		},
		ExternalLabels: labelutil.LabelSet{},
		Timeout:        time.Second * 10,
	})
	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
	client.Stop()
}

func handle(conn net.Conn) {
	defer conn.Close()
	err := parseStream(conn, func(res *syslog.Result) {
		if err := res.Error; err != nil {
			handleError(err)
			return
		}
		msg := res.Message.(*rfc5424.SyslogMessage)
		msg.SeverityLevel()
		handleMessage(msg)
	})
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	log.Println(err)
}

func handleMessage(msg *rfc5424.SyslogMessage) {
	t := time.Now()
	ls := model.LabelSet{}
	hostname := "unknown"

	if host, err := os.Hostname(); err == nil {
		hostname = host
	}
	if msg.Hostname != nil {
		hostname = *msg.Hostname
	}
	var message = "<Empty message>"
	if msg.Message != nil {
		message = *msg.Message
	}

	ls["level"] = model.LabelValue(severityLevels[*msg.Severity])
	ls["host"] = model.LabelValue(hostname)
	for id, fields := range *msg.StructuredData {
		parts := strings.Split(id, "@")
		if parts[0] == "QuLog" {
			ls["type"] = model.LabelValue(parts[1])
			for k, v := range fields {
				if v == "---" || strings.HasSuffix(k, "_id") || v == "" {
					continue
				}
				switch k {
				case "service":
					id, _ := strconv.Atoi(v)
					ls["service"] = model.LabelValue(serviceMap[id])
				case "action":
					id, _ := strconv.Atoi(v)
					ls["action"] = model.LabelValue(actionMap[id])
					switch id {
					case actionLoginSucc:
						message = fmt.Sprintf("user %s login success", fields["user"])
					case actionLoginFail:
						message = fmt.Sprintf("user %s login fail", fields["user"])
					case actionLogout:
						message = fmt.Sprintf("user %s logout", fields["user"])
					}
				default:
					ls[model.LabelName(k)] = model.LabelValue(v)
				}
			}
		}
	}
	if msg.Timestamp != nil {
		t = *msg.Timestamp
	}
	if *useMachineTime {
		t = time.Now()
	}
	err := client.Handle(ls, t, message)
	if err != nil {
		handleError(err)
	}
}
