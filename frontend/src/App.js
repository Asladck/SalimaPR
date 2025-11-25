import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Container, CssBaseline } from '@mui/material';
import SignUp from './pages/SignUp';
import SignIn from './pages/SignIn';
import Refresh from './pages/Refresh';

function App() {
  return (
    <Router>
      <CssBaseline />
      <Container
        maxWidth="sm"
        sx={{
          bgcolor: "#e0f7fa",
          minHeight: "100vh",
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          py: 4,
        }}
      >
        <Routes>
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="/sign-in" element={<SignIn />} />
          <Route path="/refresh" element={<Refresh />} />
        </Routes>
      </Container>
    </Router>
  );
}

export default App;
