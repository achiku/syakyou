package main

const (
	cardStatusType = iota
	userStatusType
)

// statusMapper all statuses in single struct? like 'active', 'inactive', 'banned', etc.
type statusMapper struct {
	ServiceAMap map[int][]string
	ServiceBMap map[int][]string
	ServiceCMap map[int][]int
}

// mapper intermidiate state
type mapper struct {
	AStatuses   []string
	BStatuses   []string
	CStatuses   []int
	targetIndex int
}

func (m *statusMapper) Key(statusKey int) *mapper {
	mp := &mapper{
		AStatuses: m.ServiceAMap[statusKey],
		BStatuses: m.ServiceBMap[statusKey],
		CStatuses: m.ServiceCMap[statusKey],
	}
	return mp
}

func (m *mapper) ServiceA(status string) *mapper {
	for i, s := range m.AStatuses {
		if status == s {
			m.targetIndex = i
		}
	}
	return m
}

func (m *mapper) ServiceB(status string) *mapper {
	for i, s := range m.BStatuses {
		if status == s {
			m.targetIndex = i
		}
	}
	return m
}

func (m *mapper) ServiceC(status int) *mapper {
	for i, s := range m.CStatuses {
		if status == s {
			m.targetIndex = i
		}
	}
	return m
}

func (m *mapper) ToServiceA() string {
	return m.AStatuses[m.targetIndex]
}

func (m *mapper) ToServiceB() string {
	return m.BStatuses[m.targetIndex]
}

func (m *mapper) ToServiceC() int {
	return m.CStatuses[m.targetIndex]
}

var statusA = statusMapper{
	ServiceAMap: map[int][]string{
		cardStatusType: []string{"active", "inactive", "banned"},
		userStatusType: []string{"good", "bad", "not too bad"},
	},
	ServiceBMap: map[int][]string{
		cardStatusType: []string{"A", "I", "B"},
		userStatusType: []string{"ok", "bad", "shit"},
	},
	ServiceCMap: map[int][]int{
		cardStatusType: []int{1, 2, 3},
		userStatusType: []int{0, 1, 2},
	},
}
