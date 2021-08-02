package instrumentation

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var Stats = new(Collectors)

type Collectors struct {
	RequestDurationHistogram    *prometheus.HistogramVec
	RecommendCounter            prometheus.Counter
	RecommendSelectedCounter    prometheus.Counter
	SearchCounter               prometheus.Counter
	SearchSelectedCounter       prometheus.Counter
	AutocompleteCounter         prometheus.Counter
	AutocompleteSelectedCounter prometheus.Counter
}

func (c *Collectors) Init() {
	c.RequestDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "feed_api",
		Subsystem: "api",
		Name:      "request_duration",
		Help:      "Time (in seconds) spent serving HTTP requests.",
	}, []string{"method", "route", "status_code"})

	c.RecommendCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "feed_api",
		Subsystem: "recommend",
		Name:      "recommend",
		Help:      "Counts number of recommendation served to users.",
	})

	c.RecommendSelectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "feed_api",
		Subsystem: "recommend",
		Name:      "recommend_selected",
		Help:      "Counts number of recommendation selections by users.",
	})

	c.SearchCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "feed_api",
		Subsystem: "search",
		Name:      "search",
		Help:      "Counts number of searches served to users.",
	})

	c.SearchSelectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "feed_api",
		Subsystem: "search",
		Name:      "search_selected",
		Help:      "Counts number of search results selected by users.",
	})

	c.AutocompleteCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "feed_api",
		Subsystem: "search",
		Name:      "autocomplete",
		Help:      "Counts number of autocompletes served to users.",
	})

	c.AutocompleteSelectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "feed_api",
		Subsystem: "search",
		Name:      "autocomplete_selected",
		Help:      "Counts number of autocomplete suggestions selected by users.",
	})

	prometheus.MustRegister(c.RequestDurationHistogram)
	prometheus.MustRegister(c.RecommendCounter)
	prometheus.MustRegister(c.RecommendSelectedCounter)
	prometheus.MustRegister(c.SearchCounter)
	prometheus.MustRegister(c.SearchSelectedCounter)
	prometheus.MustRegister(c.AutocompleteCounter)
	prometheus.MustRegister(c.AutocompleteSelectedCounter)
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
}

func (c *Collectors) Reset() {
	c.RequestDurationHistogram.Reset()
}
