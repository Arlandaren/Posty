// src/components/CommentList.js

import React, { useEffect, useState } from 'react';
import { getComments } from '../api';

function CommentList({ postId }) {
  const [comments, setComments] = useState([]);

  useEffect(() => {
    getComments(postId)
      .then((response) => {
        console.log('Полученные комментарии:', response.data.comments);
        setComments(response.data.comments || []);
      })
      .catch((error) => {
        console.error('Ошибка при получении комментариев:', error);
      });
  }, [postId]);

  return (
    <ul className="list-group mb-4">
      {comments.length === 0 ? (
        <li className="list-group-item">Нет комментариев.</li>
      ) : (
        comments.map((comment) => {
          // Преобразование времени
          const createdAt = new Date(comment.createdAt).toLocaleString()
          return (
            <li key={comment.id} className="list-group-item">
              <p>{comment.comment}</p>
              <div className="text-muted">
                Автор ID: {comment.creatorId} | Дата: {createdAt}
              </div>
            </li>
          );
        })
      )}
    </ul>
  );
}

export default CommentList;
