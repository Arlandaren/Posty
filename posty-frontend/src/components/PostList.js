// src/components/PostList.js

import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { getPosts, getImageUrl } from '../api';

function PostList() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    getPosts()
      .then((response) => {
        setPosts(response.data.posts || []);
      })
      .catch((error) => {
        console.error('Ошибка при получении постов:', error);
      });
  }, []);

  return (
    <div className="container mt-4">
      <h1 className="mb-4">Список постов</h1>
      <Link to="/create" className="btn btn-primary mb-3">
        <i className="fas fa-plus me-2"></i>
        Создать новый пост
      </Link>
      {posts.length === 0 ? (
        <p>Нет доступных постов.</p>
      ) : (
        <div className="row">
          {posts.map((post) => (
            <div key={post.id} className="col-md-4 mb-3">
              <div className="card h-100">
                {post.image && (
                  <img
                    src={getImageUrl(post.image)}
                    className="card-img-top"
                    alt={post.title}
                  />
                )}
                <div className="card-body d-flex flex-column">
                  <h5 className="card-title">{post.title}</h5>
                  <p className="card-text text-muted mb-4">Автор ID: {post.creatorId}</p>
                  <Link to={`/posts/${post.id}`} className="mt-auto btn btn-outline-primary">
                    Читать больше
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default PostList;
