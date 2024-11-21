// src/components/CreateComment.js

import React, { useState } from 'react';
import { createComment } from '../api';

function CreateComment({ postId }) {
  const [comment, setComment] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    createComment(postId, comment)
      .then(() => {
        setComment('');
        window.location.reload(); // Обновляем страницу после добавления комментария
      })
      .catch((error) => {
        console.error('Ошибка при добавлении комментария:', error);
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
        ></textarea>
      </div>
      <button type="submit" className="btn btn-primary">
        <i className="fas fa-paper-plane me-2"></i>
        Отправить
      </button>
    </form>
  );
}

export default CreateComment;
