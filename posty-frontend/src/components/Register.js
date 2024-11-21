  // src/components/Register.js

  import React, { useState } from 'react';
  import { useNavigate } from 'react-router-dom';
  import { registerUser, setAuthToken } from '../api';

  function Register() {
    const [name, setName] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSubmit = (e) => {
      e.preventDefault();
      registerUser(name, password)
        .then((response) => {
          const token = response.data.token;
          setAuthToken(token);
          localStorage.setItem('token', token);
          navigate('/');
        })
        .catch((error) => {
          console.error('Ошибка при регистрации:', error);
        });
    };

    return (
      <div className="container mt-4">
        <h1>Регистрация</h1>
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
          <button type="submit" className="btn btn-success">
            Зарегистрироваться
          </button>
        </form>
      </div>
    );
  }

  export default Register;
