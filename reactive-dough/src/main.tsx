import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import {
  QueryClient,
  QueryClientProvider
} from '@tanstack/react-query'

import Home from './pages/home'
import AuthPage from './pages/auth'
import CreatePostPage from './pages/createpost/createpostpage'
import ViewPostPage from './pages/viewpostpage'

import {
  createBrowserRouter,
  RouterProvider
} from 'react-router-dom'

import './index.css'

import { loader as postLoader } from './pages/viewpostpage'

const queryClient = new QueryClient()

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home/>
  },
  {
    path: "/myprofile",
    element: <div>This is the profile page</div>
  },
  {
    path: "/contact",
    element: <div>This is the contact page</div>
  },
  {
    path: "/authpage",
    element: <AuthPage/>
  },
  {
    path: "/createpost",
    element: <CreatePostPage/>
  },
  {
    path: "/userdashboard",
    element: <div>This is the dashboard page</div>
  },
  {
    path: "/viewpost/:postId",
    element: <ViewPostPage />,
    loader: postLoader
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router}/>
    </QueryClientProvider>
  </React.StrictMode>,
)
