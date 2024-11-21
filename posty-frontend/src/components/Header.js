import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { setAuthToken } from '../api';
import { useSnackbar } from 'notistack'; // Импортируем useSnackbar

function Header() {
  const navigate = useNavigate();
  const { enqueueSnackbar } = useSnackbar(); // Используем useSnackbar
  const token = localStorage.getItem('token');

  const handleLogout = () => {
    // Удаляем токен из localStorage
    localStorage.removeItem('token');
    // Удаляем токен из заголовков axios
    setAuthToken(null);
    enqueueSnackbar('Вы вышли из системы', { variant: 'info' });
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark mb-4">
      {/*Остальной код навигации*/}
      <div className="navbar-nav ms-auto">
        <Link className="nav-link" to="/">
          Главная
        </Link>
        {token ? (
          <>
            <Link className="nav-link" to="/create">
              Создать пост
            </Link>
            <button className="nav-link btn btn-link" onClick={handleLogout}>
              Выйти
            </button>
          </>
        ) : (
          <>
            <Link className="nav-link" to="/login">
              Вход
            </Link>
            <Link className="nav-link" to="/register">
              Регистрация
            </Link>
          </>
        )}
      </div>
    </nav>
  );
}

export default Header;
