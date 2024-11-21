// src/api.js

import axios from 'axios';

const API_BASE_URL = 'http://localhost:8086/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
});

const AUTH_API_URL = 'http://localhost:8080/v1/auth';

const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJyb2xlIjoidXNlciIsImV4cCI6MTczMjIxNzQ0NiwiaWF0IjoxNzMyMTMxMDQ2LCJpc3MiOiJhdXRoX3NlcnZpY2UifQ.Ib0GibbYGTQyPAsLJHBYErseUf9KmVpT3kfDQCapjLs'; // Замените на ваш токен

// Добавляем интерцептор запросов
api.interceptors.request.use(
  (config) => {
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
      console.log('Заголовок Authorization добавлен в запрос:', config.headers['Authorization']);
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export const setAuthToken = () => {
  const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJyb2xlIjoidXNlciIsImV4cCI6MTczMjIxNzQ0NiwiaWF0IjoxNzMyMTMxMDQ2LCJpc3MiOiJhdXRoX3NlcnZpY2UifQ.Ib0GibbYGTQyPAsLJHBYErseUf9KmVpT3kfDQCapjLs'; // Замените на ваш токен
  api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  console.log('Токен установлен:', api.defaults.headers.common['Authorization']);
};

// Регистрация пользователя
export const registerUser = (name, password) => {
  return axios.post(`${AUTH_API_URL}/register`, { name, password });
};

// Вход пользователя
export const loginUser = (name, password) => {
  return axios.post(`${AUTH_API_URL}/login`, { name, password });
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
  return api.post('/posts', { title, image});
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
