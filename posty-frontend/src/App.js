// src/App.js

import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import Footer from './components/Footer';
import PostList from './components/PostList';
import PostDetail from './components/PostDetail';
import CreatePost from './components/CreatePost';
import LoginRedirect from './components/LoginRedirect';
import RegisterRedirect from './components/RegisterRedirect';
import AuthCallback from './components/AuthCallback';

function App() {
  return (
    <Router>
      <Header />
      <div className="flex-grow-1">
        <Routes>
          <Route path="/" element={<PostList />} />
          <Route path="/posts/:id" element={<PostDetail />} />
          <Route path="/create" element={<CreatePost />} />
          <Route path="/login" element={<LoginRedirect />} />
          <Route path="/register" element={<RegisterRedirect />} />
          <Route path="/auth/callback" element={<AuthCallback />} />
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

