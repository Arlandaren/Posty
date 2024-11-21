// src/components/RegisterRedirect.js

import React, { useEffect } from 'react';

function RegisterRedirect() {
  useEffect(() => {
    const redirectUri = encodeURIComponent('http://localhost:3002/auth/callback');
    window.location.href = `http://localhost:3001/register?redirect_uri=${redirectUri}`;
  }, []);

  return <div>Перенаправление на страницу регистрации...</div>;
}

export default RegisterRedirect;
