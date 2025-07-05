import React, { useState } from 'react';
import { 
  Box, 
  TextField, 
  Button, 
  Typography, 
  Paper,
  Alert 
} from '@mui/material';
import { greeterClient } from '../services/greeter';
import { create } from '@bufbuild/protobuf';
import { GreetRequestSchema } from '../gen/greeter/v1/greeter_pb';

export const GreeterForm: React.FC = () => {
  const [name, setName] = useState('');
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      const request = create(GreetRequestSchema, { name });
      const response = await greeterClient.greet(request);
      setMessage(response.message);
    } catch (err) {
      setError('Failed to get greeting');
      console.error('Greeting error:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box sx={{ maxWidth: 500, mx: 'auto', mt: 4 }}>
      <Paper elevation={3} sx={{ p: 4 }}>
        <Typography variant="h4" gutterBottom align="center">
          Greeter Service
        </Typography>
        
        <form onSubmit={handleSubmit}>
          <TextField
            fullWidth
            label="Your Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            margin="normal"
            placeholder="Enter your name"
          />
          
          <Button
            type="submit"
            variant="contained"
            fullWidth
            disabled={loading}
            sx={{ mt: 2 }}
          >
            {loading ? 'Greeting...' : 'Get Greeting'}
          </Button>
        </form>

        {error && (
          <Alert severity="error" sx={{ mt: 2 }}>
            {error}
          </Alert>
        )}

        {message && (
          <Box sx={{ mt: 3 }}>
            <Typography variant="h6" align="center" color="primary">
              {message}
            </Typography>
          </Box>
        )}
      </Paper>
    </Box>
  );
};
