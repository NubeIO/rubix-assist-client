package rest

/*
REST Setters
*/

func (s *Service) SetPath(path string) *Service {
	method := s.Method
	if s.NubeProxy.UseRubixProxy { //set rubix proxy
		r := s.GetToken()
		s.Options.Headers = map[string]interface{}{"Authorization": r.Token}
		s.Path = path
		p, port := s.FixPath()
		s.Path = p
		s.Port = port
		s.Method = method
	} else {
		s.Path = path
		s.Method = method
	}
	return s
}

func (s *Service) SetBody(body interface{}) *Service {
	if s.Options != nil {
		s.Options.Body = body
	}
	return s
}

func (s *Service) SetMethod(method string) *Service {
	s.Method = method
	return s
}
