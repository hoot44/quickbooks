package quickbooks

type QuickbooksError struct {
	Fault struct {
		Error []struct {
			Message string
			Detail  string
			Code    string
		}
		Type string
	}
	Time string
	// Avoid the function name
	Err       string `json:"error,omitempty"`
	IntuitTid string
}

func (q *QuickbooksError) Error() string {
	if q == nil {
		return "??"
	}
	if len(q.Fault.Error) == 0 {
		if q.Err != "" {
			return q.Err
		}
		return q.Fault.Type
	}
	s := ""
	for _, l := range q.Fault.Error {
		s += l.Detail + "\n"
	}
	return s
}
