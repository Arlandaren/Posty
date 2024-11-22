import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import 'bootstrap/dist/css/bootstrap.min.css'; // Импортируем Bootstrap CSS
import '@fortawesome/fontawesome-free/css/all.min.css'; // Импортируем Font Awesome
import './index.css'; // Импортируем собственные стили
import { SnackbarProvider } from 'notistack'; // Импортируем SnackbarProvider

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <SnackbarProvider maxSnack={3}> {/*Оборачиваем приложение в SnackbarProvider*/}
    <App />
  </SnackbarProvider>
);
