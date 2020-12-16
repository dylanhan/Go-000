package server

type Server struct {
	service *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{service: s}
}

func (srv *Server) Run() error {
	g, ctx := errgroup.WithContext(context.Background())
	s := http.Server{Addr: "127.0.0.0:8000", Handler: nil}
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			s.Shutdown(context.Background())
		}()
		return s.ListenAndServe()
	})

	exit := make(chan os.Signal)
	g.Go(func() error {
		signal.Notify(exit)
		select {
		case <-exit:
			return errors.New("receive signal, exit")
		case <-ctx.Done():
			fmt.Println("signal all cancel")
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}