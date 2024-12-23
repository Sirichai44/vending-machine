import { RouterProvider } from 'react-router';
import { createRoot } from 'react-dom/client';
import { CssVarsProvider } from '@mui/joy';

import '@/styles/tailwind.css';
import '@/styles/global.css';
import router from '@/routes/index.tsx';
import theme from '@/styles/theme.ts';

createRoot(document.getElementById('root')!).render(
  <CssVarsProvider theme={theme}>
    <RouterProvider router={router} />
  </CssVarsProvider>
);
