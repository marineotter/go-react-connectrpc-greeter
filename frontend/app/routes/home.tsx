import React from 'react';
import { Container, CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import { GreeterForm } from '../components/GreeterForm';

const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#1976d2',
    },
  },
});

export function meta() {
  return [
    { title: "Greeter Service" },
    { name: "description", content: "Connect-RPC Greeter Service" },
  ];
}

export default function Home() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Container>
        <GreeterForm />
      </Container>
    </ThemeProvider>
  );
}
