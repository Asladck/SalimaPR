import React, { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Typography, Box } from '@mui/material';

function SignIn() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('/auth/sign-in', { email, password });
      alert('Sign-in successful!');
    } catch (error) {
      alert('Error during sign-in.');
    }
  };

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      textAlign="center"
      mt={4}
      p={3}
      borderRadius={2}
      boxShadow={3}
      bgcolor="#f5f5f5"
    >
      <Typography variant="h4" gutterBottom color="primary">
        Sign In
      </Typography>
      <TextField
        label="Email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        fullWidth
        margin="normal"
        variant="outlined"
      />
      <TextField
        label="Password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        fullWidth
        margin="normal"
        variant="outlined"
      />
      <Button
        type="submit"
        variant="contained"
        color="primary"
        fullWidth
        size="large"
        sx={{ mt: 2 }}
      >
        Sign In
      </Button>
    </Box>
  );
}

export default SignIn;
