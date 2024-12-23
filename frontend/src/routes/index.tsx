/* eslint-disable react-refresh/only-export-components */
import { Suspense, lazy } from 'react';
import { createBrowserRouter } from 'react-router-dom';
// import ErrorPage from '@/pages/ErrorPage';
import SuspenseWrapper from './SuspenseWrapper';
// // import { LinearProgress } from '@mui/joy';

import ErrorPage from '@/pages/ErrorPage';


const App = lazy(() => import('@/App'));
const Home = lazy(() => import('@/pages/Home'));


const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <Suspense fallback={<SuspenseWrapper />}>
        <App />
      </Suspense>
    ),
    children: [
      {path: '', element: <Home />},
      { path: '*', element: <ErrorPage /> } // 404 page
    ]
  }
]);

export default router;
// const router = createBrowserRouter([
//   {
//     path: '/',
//     element: (
//       <Suspense fallback={<SuspenseWrapper />}>
//         <App />
//       </Suspense>
//     ),
//     children: [
//       { path: '', element: <Home /> },
//       {
//         path: '/auth/',
//         element: <SuspenseWrapper />,
//         children: [
//           { path: 'login', element: <Login /> },
//           { path: 'register', element: <Register /> }
//         ]
//       },
//       {
//         path: '/',
//         element: <Root />,
//         children: [
//           {
//             path: 'blog',
//             element: <Blog />
//           },
//           {
//             path: 'about',
//             element: <About />
//           },
//           {
//             path: 'certificate',
//             element: <Certificate />
//           },
//           {
//             path: 'setting',
//             element: (
//               <AuthProvider>
//                 <Setting />
//               </AuthProvider>
//             )
//           },
//           {
//             path: 'assistant',
//             element: (
//               <AuthProvider>
//                 <Assistant />
//               </AuthProvider>
//             )
//           }
//         ]
//       },
//       { path: '*', element: <ErrorPage /> } // 404 page
//     ]
//   }
// ]);

// export default router;
