// src/components/CreatePost.js

import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { createPost, uploadImage } from '../api';

function CreatePost() {
  const [title, setTitle] = useState('');
  const [imageFile, setImageFile] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
    }
  }, [navigate]);

  const handleSubmit = (e) => {
    e.preventDefault();

    if (imageFile) {
      // Сначала загружаем изображение
      uploadImage(imageFile)
        .then((response) => {
          const filename = response.data.filename;
          // Затем создаем пост с полученным filename
          return createPost(title, filename);
        })
        .then((response) => {
          navigate(`/posts/${response.data.id}`);
        })
        .catch((error) => {
          console.error('Ошибка при создании поста:', error);
        });
    } else {
      // Если изображения нет, создаем пост без него
      createPost(title, null)
        .then((response) => {
          navigate(`/posts/${response.data.id}`);
        })
        .catch((error) => {
          console.error('Ошибка при создании поста:', error);
        });
    }
  };

  return (
    <div className="container mt-4">
      <h1>Создать новый пост</h1>
      <form onSubmit={handleSubmit}>
        {/* Остальная часть формы */}
      </form>
    </div>
  );
}

export default CreatePost;
