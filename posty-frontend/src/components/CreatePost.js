import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { createPost, uploadImage } from '../api';
import { useSnackbar } from 'notistack'; // Импортируем хук useSnackbar
import { CircularProgress } from '@mui/material'; // Импортируем CircularProgress

function CreatePost() {
  const [title, setTitle] = useState('');
  const [imageFile, setImageFile] = useState(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const { enqueueSnackbar } = useSnackbar(); // Используем useSnackbar

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
    }
  }, [navigate]);

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);

    if (imageFile) {
      // Сначала загружаем изображение
      uploadImage(imageFile)
        .then((response) => {
          const filename = response.data.filename;
          // Затем создаем пост с полученным filename
          return createPost(title, filename);
        })
        .then((response) => {
          enqueueSnackbar('Пост успешно создан!', { variant: 'success' }); // Успешное уведомление
          navigate(`/posts/${response.data.id}`);
        })
        .catch((error) => {
          console.error('Ошибка при создании поста:', error);
          enqueueSnackbar('Ошибка при создании поста', { variant: 'error' }); // Уведомление об ошибке
        })
        .finally(() => {
          setLoading(false);
        });
    } else {
      // Если изображения нет, создаем пост без него
      createPost(title, null)
        .then((response) => {
          enqueueSnackbar('Пост успешно создан!', { variant: 'success' });
          navigate(`/posts/${response.data.id}`);
        })
        .catch((error) => {
          console.error('Ошибка при создании поста:', error);
          enqueueSnackbar('Ошибка при создании поста', { variant: 'error' });
        })
        .finally(() => {
          setLoading(false);
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
        <button type="submit" className="btn btn-success" disabled={loading}>
          {loading ? (
            <>
              <CircularProgress size={20} color="inherit" /> Создание...
            </>
          ) : (
            'Создать'
          )}
        </button>
      </form>
    </div>
  );
}

export default CreatePost;
