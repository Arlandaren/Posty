// src/components/PostDetail.js

import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getPostById, getImageUrl } from '../api';
import CommentList from './CommentList';
import CreateComment from './CreateComment';

function PostDetail() {
  const { id } = useParams();
  const [post, setPost] = useState(null);

  useEffect(() => {
    getPostById(id)
      .then((response) => {
        console.log('Данные поста:', response.data); // Логируем данные поста
        setPost(response.data);
      })
      .catch((error) => {
        console.error('Ошибка при получении поста:', error);
      });
  }, [id]);

  if (!post) {
    return (
      <div className="container mt-4">
        <p>Загрузка...</p>
      </div>
    );
  }

  return (
    <div className="container mt-4">
      <h1>{post.title}</h1>
      <p className="text-muted">Автор ID: {post.creatorId}</p>
      <p>Дата создания: {new Date(post.createdAt).toLocaleString()}</p>

      {post.image && (
        <div className="mb-4">
          <img
            src={getImageUrl(post.image)}
            alt={post.title}
            className="img-fluid"
          />
        </div>
      )}

      <h2 className="mt-5">Комментарии</h2>
      <CreateComment postId={id} />
      <CommentList postId={id} />
    </div>
  );
}

export default PostDetail;
