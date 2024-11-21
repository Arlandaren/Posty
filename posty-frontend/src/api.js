// src/api.js

import axios from 'axios';

const API_BASE_URL = 'http://localhost:8086/v1'; // Адрес API вашего бэкенда Posty

const api = axios.create({
  baseURL: API_BASE_URL,
});

// Добавляем интерцептор запросов
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token'); // Получаем токен из localStorage
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
      console.log('Заголовок Authorization добавлен в запрос:', config.headers['Authorization']);
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Функция для установки токена при входе или регистрации
export const setAuthToken = (token) => {
  if (token) {
    api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    delete api.defaults.headers.common['Authorization'];
  }
};


// Функции API

// Получить список постов
export const getPosts = (cursor = 0, limit = 10) => {
  return api.get('/posts', {
    params: { cursor, limit },
  });
};

// Получить пост по ID
export const getPostById = (id) => {
  return api.get(`/posts/${id}`);
};

// Создать новый пост
export const createPost = (title, image) => {
  return api.post('/posts', { title, image });
};

// Получить список комментариев к посту
export const getComments = (postId, cursor = 0, limit = 10) => {
  return api.get(`/posts/${postId}/comments`, {
    params: { cursor, limit },
  });
};

// Добавить комментарий к посту
export const createComment = (postId, comment) => {
  return api.post(`/posts/${postId}/comments`, { postId, comment });
};

export const uploadImage = (file) => {
  const formData = new FormData();
  formData.append('file', file);

  return api.post('/images', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// Получить URL изображения
export const getImageUrl = (image) => {
  return `${API_BASE_URL}/images/${image}`;
};

export default api;
