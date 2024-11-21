// src/components/Header.js

import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { setAuthToken } from '../api';

function Header() {
  const navigate = useNavigate();
  const token = localStorage.getItem('token');

  const handleLogout = () => {
    localStorage.removeItem('token');
    setAuthToken(null);
    navigate('/login');
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark mb-4">
      <div className="container">
        <Link className="navbar-brand" to="/">
          PostyApp
        </Link>
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
          aria-controls="navbarNav"
          aria-expanded="false"
          aria-label="Переключить навигацию"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarNav">
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
        </div>
      </div>
    </nav>
  );
}

export default Header;
