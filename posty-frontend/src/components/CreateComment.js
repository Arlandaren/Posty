import React, { useState } from 'react';
import { createComment } from '../api';
import { useSnackbar } from 'notistack'; // Импортируем useSnackbar

function CreateComment({ postId }) {
  const [comment, setComment] = useState('');
  const [loading, setLoading] = useState(false);
  const { enqueueSnackbar } = useSnackbar(); // Используем useSnackbar

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);
    createComment(postId, comment)
      .then(() => {
        setComment('');
        enqueueSnackbar('Комментарий добавлен!', { variant: 'success' });
        // Обновляем комментарии без перезагрузки страницы
        if (typeof window !== 'undefined') {
          window.location.reload();
        }
      })
      .catch((error) => {
        console.error('Ошибка при добавлении комментария:', error);
        enqueueSnackbar('Ошибка при добавлении комментария', { variant: 'error' });
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4">
      <div className="mb-3">
        <label className="form-label">Добавить комментарий</label>
        <textarea
          className="form-control"
          rows="3"
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          required
          disabled={loading}
        ></textarea>
      </div>
      <button type="submit" className="btn btn-primary" disabled={loading}>
        {loading ? (
          <>
            <span className="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
            Отправка...
          </>
        ) : (
          <>
            <i className="fas fa-paper-plane me-2"></i>
            Отправить
          </>
        )}
      </button>
    </form>
  );
}

export default CreateComment;
