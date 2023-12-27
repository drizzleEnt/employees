package employe

import "context"

func (s *Service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
