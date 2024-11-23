// src/components/Footer.js

import React from 'react';

function Footer() {
  return (
    <footer className="bg-dark text-white mt-4 py-3">
      <div className="container text-center">
        © {new Date().getFullYear()} PostyApp. Powered by VTB.
      </div>
    </footer>
  );
}

export default Footer;
