import React from 'react';
import axios from 'axios';
import { Button, Typography, Box } from '@mui/material';

function Refresh() {
  const handleRefresh = async () => {
    try {
      const response = await axios.post('/auth/refresh');
      alert('Token refreshed!');
    } catch (error) {
      alert('Error refreshing token.');
    }
  };

  return (
    <Box
      textAlign="center"
      mt={4}
      p={3}
      borderRadius={2}
      boxShadow={3}
      bgcolor="#f5f5f5"
    >
      <Typography variant="h4" gutterBottom color="primary">
        Refresh Token
      </Typography>
      <Button
        variant="contained"
        color="primary"
        onClick={handleRefresh}
        size="large"
        sx={{ mt: 2 }}
      >
        Refresh
      </Button>
    </Box>
  );
}

export default Refresh;
