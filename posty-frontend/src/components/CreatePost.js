// src/components/CreatePost.js

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createPost, uploadImage } from '../api';

function CreatePost() {
  const [title, setTitle] = useState('');
  const [imageFile, setImageFile] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();

    if (imageFile) {
      // Сначала загружаем изображение
      uploadImage(imageFile)
        .then((response) => {
          const imageId = response.data.id;
          const filename = response.data.filename;
          // Затем создаем пост с полученными imageId и filename
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
      createPost(title, null, null)
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
        <div className="mb-3">
          <label className="form-label">Заголовок</label>
          <input
            type="text"
            className="form-control"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label">Изображение</label>
          <input
            type="file"
            className="form-control"
            accept="image/*"
            onChange={(e) => setImageFile(e.target.files[0])}
          />
        </div>
        <button type="submit" className="btn btn-success">
          Создать
        </button>
      </form>
    </div>
  );
}

export default CreatePost;
