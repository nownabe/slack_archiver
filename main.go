package slackarchiver

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"

	"github.com/nownabe/slack_archiver/slack"
)

type slackArchiver struct {
	*slack.Client
}

func init() {
	sa := &slackArchiver{
		Client: slack.New(os.Getenv("SLACK_API_TOKEN")),
	}

	http.HandleFunc("/", sa.notFound)
	http.HandleFunc("/slack_archiver/archive", sa.archiveHandler)
}

func (sa *slackArchiver) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (sa *slackArchiver) archiveHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "text/plain")

	params := r.URL.Query()

	if params.Get("async") == "true" {
		t := &taskqueue.Task{
			Method: "GET",
			Path:   "/slack_archiver/archive",
		}
		_, err := taskqueue.Add(ctx, t, "")
		if err != nil {
			log.Errorf(ctx, "Failed to queue archive task. %v", err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		log.Infof(ctx, "Queued archive task.")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
		return
	}

	channels, err := sa.ChannelsList(ctx, true, true)
	if err != nil {
		log.Errorf(ctx, "Failed to get channels list. %v", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	total := len(channels)
	archived := 0
	now := time.Now()

	for _, c := range channels {
		for count := 0; count < 5; count++ {
			history, err := sa.ChannelsHistory(ctx, c.ID, 1)
			if err == nil {
				msg := history.Messages[0]
				unixNano, err := strconv.ParseInt(strings.Replace(msg.TS, ".", "", -1), 10, 64)
				if err != nil {
					log.Errorf(ctx, "Failed to convert message timestamp into int64 %s. %v", msg.TS, err)
					break
				}
				latest := time.Unix(unixNano/1000000, 0)
				log.Debugf(ctx, "Latest post in #%s was posted at %s.", c.Name, latest.Format(time.RFC3339))

				if now.Sub(latest) > time.Duration(24*30*2)*time.Hour {
					// if err := sa.ChannelsArchive(ctx, c.ID); err != nil {
					//   log.Errorf(ctx, "Failed to archive #%s. %v", c.Name, err)
					//   break
					// }
					log.Infof(ctx, "Archived #%s %s", c.Name, c.ID)
					archived++
				}

				break
			}

			if rlerr, ok := err.(slack.RateLimitError); ok {
				log.Warningf(ctx, "Rate Limit Error. Retrying. %v", err)
				time.Sleep(time.Duration(rlerr.RetryAfter) * time.Second)
				continue
			}

			log.Errorf(ctx, "Failed to get history of channel %s (%s). %v", c.Name, c.ID, err)
			break
		}
	}

	log.Infof(ctx, "Archived %d channels (total: %d)", archived, total)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
