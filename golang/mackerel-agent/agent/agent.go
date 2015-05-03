package agent

import (
	"time"

	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/mackerel"
	"github.com/mackerelio/mackerel-agent/metrics"
)

type Agent struct {
	MetricsGenerators []metrics.Generator
	PluginGenerators  []metrics.PluginGenerator
}

type MetricsResult struct {
	Created time.Time
	Values  metrics.Values
}

func (agent *Agent) CollectMetrics(collectedTime time.Time) *MetricsResult {
	generators := agent.MetricsGenerators
	for _, g := range agent.PluginGenerators {
		generators = append(generators, g)
	}
	result := generateValues(generators)
	values := <-result
	return &MetricsResult{Created: collectedTime, Values: values}
}

func (agent *Agent) Watch() chan *MetricsResult {
	metricsResult := make(chan *MetricsResult)
	ticker := make(chan time.Time)

	go func() {
		c := time.Tick(1 * time.Second)
		last := time.Now()
		ticker <- last

		for t := range c {
			if t.Second()%int(config.PostMetricsInterval.Seconds()) == 0 || t.After(last.Add(config.PostMetricsInterval)) {
				last = t
				ticker <- t
			}
		}
	}()

	const collectMetricsWorkerMax = 3
	go func() {
		sem := make(chan uint, collectMetricsWorkerMax)
		for tickedTime := range ticker {
			ti := tickedTime
			sem <- 1
			go func() {
				metricsResult <- agent.CollectMetrics(ti)
				<-sem
			}()
		}
	}()

	return metricsResult
}

func (agent *Agent) CollectGraphDefsOfPlugins() []mackerel.CreateGraphDefsPayload {
	payloads := []mackerel.CreateGraphDefsPayload{}

	for _, g := range agent.PluginGenerators {
		p, err := g.PrepareGraphDefs()
		if err != nil {
			logger.Debugf("failed to fetch meta information from plugin %s (non critical); seems that this plugin does not have meta information: %s", g, err)
		}
		if p != nil {
			payloads = append(payloads, p...)
		}
	}

	return payloads
}

func (agent *Agent) InitPluginGenerators(api *mackerel.API) {
	payloads := agent.CollectGraphDefsOfPlugins()

	if len(payloads) > 0 {
		err := api.CreateGraphDefs(payloads)
		if err != nil {
			logger.Errorf("failed to create graphdefs: %s", err)
		}
	}
}
