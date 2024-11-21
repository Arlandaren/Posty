// src/components/LoginRedirect.js

import React, { useEffect } from 'react';

function LoginRedirect() {
  useEffect(() => {
    const redirectUri = encodeURIComponent('http://localhost:3002/auth/callback');
    window.location.href = `http://localhost:3001/login?redirect_uri=${redirectUri}`;
  }, []);

  return <div>Перенаправление на страницу входа...</div>;
}

export default LoginRedirect;
