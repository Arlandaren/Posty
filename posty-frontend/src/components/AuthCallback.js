// src/components/AuthCallback.js

import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { setAuthToken } from '../api';

function AuthCallback() {
  const navigate = useNavigate();

  useEffect(() => {
    // Получаем токен из фрагмента URL
    const hash = window.location.hash.substring(1);
    const params = new URLSearchParams(hash);
    const token = params.get('token');

    if (token) {
      // Сохраняем токен в localStorage
      localStorage.setItem('token', token);
      // Устанавливаем токен в заголовках axios
      setAuthToken(token);
      // Перенаправляем пользователя на главную страницу
      navigate('/');
    } else {
      // Если токена нет, перенаправляем на страницу входа
      navigate('/login');
    }
  }, [navigate]);

  return <div>Обработка аутентификации...</div>;
}

export default AuthCallback;

