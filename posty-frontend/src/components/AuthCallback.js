import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { setAuthToken } from '../api';
import { useSnackbar } from 'notistack'; // Импортируем useSnackbar

function AuthCallback() {
  const navigate = useNavigate();
  const { enqueueSnackbar } = useSnackbar(); // Используем useSnackbar

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
      enqueueSnackbar('Вы успешно вошли в систему!', { variant: 'success' });
      // Перенаправляем пользователя на главную страницу
      navigate('/');
    } else {
      // Если токена нет, отображаем уведомление об ошибке
      enqueueSnackbar('Ошибка аутентификации', { variant: 'error' });
      // Перенаправляем на страницу входа
      navigate('/login');
    }
  }, [navigate, enqueueSnackbar]);

  return <div>Обработка аутентификации...</div>;
}

export default AuthCallback;
