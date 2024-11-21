// src/App.js

import React, { useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { setAuthToken } from './api';
import Header from './components/Header';
import Footer from './components/Footer';
import PostList from './components/PostList';
import PostDetail from './components/PostDetail';
import CreatePost from './components/CreatePost';
import Login from './components/Login';
import Register from './components/Register';

function App() {
  useEffect(() => {
    setAuthToken();
  }, []);

  return (
    <Router>
      <Header />
      <div className="flex-grow-1">
        <Routes>
          <Route path="/" element={<PostList />} />
          <Route path="/posts/:id" element={<PostDetail />} />
          <Route path="/create" element={<CreatePost />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route
            path="*"
            element={
              <div className="container mt-4">
                <h1>Страница не найдена</h1>
              </div>
            }
          />
        </Routes>
      </div>
      <Footer />
    </Router>
  );
}

export default App;
