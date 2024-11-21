  // src/components/Login.js

  import React, { useState } from 'react';
  import { useNavigate } from 'react-router-dom';
  import { loginUser, setAuthToken } from '../api';

  function Login() {
    const [name, setName] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSubmit = (e) => {
      e.preventDefault();
      loginUser(name, password)
        .then((response) => {
          const token = response.data.token;
          setAuthToken(token);
          localStorage.setItem('token', token);
          navigate('/');
        })
        .catch((error) => {
          console.error('Ошибка при входе:', error);
        });
    };

    return (
      <div className="container mt-4">
        <h1>Вход</h1>
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label className="form-label">Имя</label>
            <input
              type="text"
              className="form-control"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <div className="mb-3">
            <label className="form-label">Пароль</label>
            <input
              type="password"
              className="form-control"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="btn btn-primary">
            Войти
          </button>
        </form>
      </div>
    );
  }

  export default Login;
