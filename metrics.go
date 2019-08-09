package gomnik

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	PVVoltage   *prometheus.GaugeVec
	PVCurrent   *prometheus.GaugeVec
	ACCurrent   *prometheus.GaugeVec
	ACVoltage   *prometheus.GaugeVec
	ACPower     *prometheus.GaugeVec
	ACFrequency *prometheus.GaugeVec
	EnergyToday prometheus.Gauge
	EnergyTotal prometheus.Gauge
	HoursTotal  prometheus.Gauge
}

func valueOrDefault(v interface{}, s uint) (o float64) {
	switch u := v.(type) {
	case uint16:
		if u == 0xffff {
			return -1.0
		}
		o = float64(u) / float64(s)
	case uint32:
		if u == 0xffffffff {
			return -1.0
		}
		o = float64(u) / float64(s)
	}
	return
}

func (m *Metrics) SetFromResponse(r Response) {
	m.PVVoltage.WithLabelValues("1").Set(valueOrDefault(r.PVVoltage1, 10))
	m.PVVoltage.WithLabelValues("2").Set(valueOrDefault(r.PVVoltage2, 10))
	m.PVVoltage.WithLabelValues("3").Set(valueOrDefault(r.PVVoltage3, 10))
	m.PVCurrent.WithLabelValues("1").Set(valueOrDefault(r.PVCurrent1, 10))
	m.PVCurrent.WithLabelValues("2").Set(valueOrDefault(r.PVCurrent2, 10))
	m.PVCurrent.WithLabelValues("3").Set(valueOrDefault(r.PVCurrent3, 10))
	m.ACCurrent.WithLabelValues("1").Set(valueOrDefault(r.ACCurrent1, 10))
	m.ACCurrent.WithLabelValues("2").Set(valueOrDefault(r.ACCurrent2, 10))
	m.ACCurrent.WithLabelValues("3").Set(valueOrDefault(r.ACCurrent3, 10))
	m.ACVoltage.WithLabelValues("1").Set(valueOrDefault(r.ACVoltage1, 10))
	m.ACVoltage.WithLabelValues("2").Set(valueOrDefault(r.ACVoltage2, 10))
	m.ACVoltage.WithLabelValues("3").Set(valueOrDefault(r.ACVoltage3, 10))
	m.ACPower.WithLabelValues("1").Set(valueOrDefault(r.ACPower1, 1))
	m.ACPower.WithLabelValues("2").Set(valueOrDefault(r.ACPower2, 1))
	m.ACPower.WithLabelValues("3").Set(valueOrDefault(r.ACPower3, 1))
	m.ACFrequency.WithLabelValues("1").Set(valueOrDefault(r.ACFrequency1, 100))
	m.ACFrequency.WithLabelValues("2").Set(valueOrDefault(r.ACFrequency2, 100))
	m.ACFrequency.WithLabelValues("3").Set(valueOrDefault(r.ACFrequency3, 100))
	m.EnergyToday.Set(valueOrDefault(r.EToday, 100))
	m.EnergyTotal.Set(valueOrDefault(r.ETotal, 10))
	m.HoursTotal.Set(valueOrDefault(r.HTotal, 1))
	return
}

func NewMetrics(r *prometheus.Registry) (m *Metrics) {
	m = &Metrics{}
	m.PVVoltage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "omnik_pv_voltage",
		Help: "",
	},
		[]string{
			"string",
		},
	)
	m.PVCurrent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "omnik_pv_current",
		Help: "",
	},
		[]string{
			"string",
		},
	)
	m.ACCurrent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "omnik_ac_current",
		Help: "",
	},
		[]string{
			"phase",
		},
	)
	m.ACVoltage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "omnik_ac_voltage",
		Help: "Voltage of AC outputs",
	},
		[]string{
			"phase",
		},
	)
	m.ACPower = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "omnik_ac_power",
		Help: "Power of AC outputs",
	},
		[]string{
			"phase",
		},
	)
	m.ACFrequency = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "omnik_ac_frequency",
		Help: "",
	},
		[]string{
			"phase",
		},
	)
	m.EnergyToday = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "omnik_energy_today",
		Help: "",
	})
	m.EnergyTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "omnik_energy_total",
		Help: "",
	})
	m.HoursTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "omnik_hours_total",
		Help: "",
	})
	r.MustRegister(
		m.PVVoltage,
		m.PVCurrent,
		m.ACCurrent,
		m.ACVoltage,
		m.ACFrequency,
		m.ACPower,
		m.EnergyToday,
		m.EnergyTotal,
		m.HoursTotal,
	)
	return
}
